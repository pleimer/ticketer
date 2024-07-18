package services

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

var ingestorWorkflowID = "email-ingestor-workflow"

type LongRunningOperationsService struct {
	client      client.Client
	nylasClient *nylas.NylasClient
	logger      *zap.Logger
}

func NewLongRunningOperationsService(logger *zap.Logger, nylas *nylas.NylasClient) *LongRunningOperationsService {
	temporalClient, err := client.Dial(client.Options{})
	if err != nil {
		logger.Fatal("dialing temporal cluster", zap.Error(err))
	}

	lro := &LongRunningOperationsService{
		client:      temporalClient,
		nylasClient: nylas,
		logger:      logger,
	}

	lro.Start()

	return lro
}

func (lro *LongRunningOperationsService) Start() {
	opts := client.StartWorkflowOptions{
		ID:           ingestorWorkflowID,
		TaskQueue:    "email-ingestor-taskqueue",
		CronSchedule: "*/1 * * * *", // Run every 5 minutes
	}

	// Even in the case of multiple instances of this server running,
	// only one workflow will be executed at a time
	we, err := lro.client.ExecuteWorkflow(context.Background(), opts, lro.EmailIngestorWorkflow)
	if err != nil {
		lro.logger.Sugar().Fatalf("Unabe to execute workflow: %v", err)
	}

	lro.logger.Sugar().Infof("Started workflow: WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
}

func (lro *LongRunningOperationsService) Close() {
	// for dev, simpler to go ahead and cancel the running workflows when exiting
	// In a production environment, this would not be ideal as there may be other instances
	// of the server still running or data should still be processed and stored
	// even if the servers go down. For production, a more detailed plan must be
	// created for managing workflows
	err := lro.client.CancelWorkflow(context.Background(), ingestorWorkflowID, "")
	if err != nil {
		lro.logger.Error("cancelling ingestor ingestor workflow", zap.Error(err))
	}
	lro.client.Close()
}

func (lro *LongRunningOperationsService) EmailIngestorWorkflow(ctx workflow.Context) error {

	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumAttempts:    5,
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var newMessagesResp *nylas.MessagesResponse

	err := workflow.ExecuteActivity(ctx, lro.QueryNewMessagesActivity).Get(ctx, &newMessagesResp)
	if err != nil {
		return err
	}

	// run update message statuses async
	updateStatusFutures := []workflow.Future{}
	for _, msg := range newMessagesResp.Data {
		updateStatusFutures = append(updateStatusFutures, workflow.ExecuteActivity(ctx, lro.UpdateMessageReadStatusActivity, msg.ID))
	}

	var processNewMessagesRes []SendTicketCreationAcknowledgementRequest
	err = workflow.ExecuteActivity(ctx, lro.ProcessNewMessagesActivity, newMessagesResp).Get(ctx, &processNewMessagesRes)
	if err != nil {
		return err
	}

	// spawn child workflow "fire and forget" style for notifying users of new tickets created in response to their inquiries
	childWFOpts := workflow.ChildWorkflowOptions{
		ParentClosePolicy: enums.PARENT_CLOSE_POLICY_ABANDON, // allow child workflow to continue after parent completes
	}
	ctx = workflow.WithChildOptions(ctx, childWFOpts)
	childFuture := workflow.ExecuteChildWorkflow(ctx, lro.TicketCreationAcknowledgementChildWorkflow, processNewMessagesRes)
	var childWE workflow.Execution
	if err := childFuture.GetChildWorkflowExecution().Get(ctx, &childWE); err != nil {
		return err
	}

	// wait for message status update activities to complete
	for _, f := range updateStatusFutures {
		err = f.Get(ctx, nil)
		if err != nil {
			// no need to treat these as workflow failures
			// if message status update failed, process messages activity will filter
			// ones out that already resulted in a new ticket.
			// just log the failure
			lro.logger.Error("update message status failure", zap.Error(err))
		}
	}

	return nil
}

func (lro *LongRunningOperationsService) TicketCreationAcknowledgementChildWorkflow(ctx workflow.Context, processNewMessagesRes []SendTicketCreationAcknowledgementRequest) (err error) {

	// since this workflow runs as "fire and forget", retries can happen for a much longer period of time for message send failures
	sendAckActivityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    100 * time.Second,
			MaximumAttempts:    500,
		},
	}

	ctx = workflow.WithActivityOptions(ctx, sendAckActivityOptions)

	futures := []workflow.Future{}

	for _, r := range processNewMessagesRes {
		futures = append(futures, workflow.ExecuteActivity(ctx, lro.SendTicketCreationAcknowledgementActivity, r))
	}

	for _, f := range futures {
		err = f.Get(ctx, nil)
		if err != nil {
			// log failures, but not much else can be done after retries exhausted
			lro.logger.Error("sending message failure", zap.Error(err))
		}
	}

	return nil
}

func (lro *LongRunningOperationsService) QueryNewMessagesActivity(ctx context.Context) (*nylas.MessagesResponse, error) {

	// TODO: paging
	msgs, err := lro.nylasClient.GetUnreadMessages(5)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving messages from nylas")
	}

	return msgs, nil
}

// UpdateMessageReadStatusActivity updates status of messages to 'read'
func (lro *LongRunningOperationsService) UpdateMessageReadStatusActivity(ctx context.Context, messageID string) (err error) {
	_, err = lro.nylasClient.UpdateMessageReadStatus(messageID, false)
	return err
}

// ProcessNewMessagesActivity creates or updates tickets for incoming messages
func (lro *LongRunningOperationsService) ProcessNewMessagesActivity(ctx context.Context, msgs *nylas.MessagesResponse) (res []SendTicketCreationAcknowledgementRequest, err error) {
	if msgs == nil {
		return nil, nil
	}

	res = []SendTicketCreationAcknowledgementRequest{}

	for _, m := range msgs.Data {
		res = append(res, SendTicketCreationAcknowledgementRequest{
			Initiator: m.From[0], // Initial messages resulting in ticket will only have one "from" entry
			TicketID:  10,
			MessageID: m.ID,
		})
	}

	return res, nil
}

type SendTicketCreationAcknowledgementRequest struct {
	Initiator nylas.Participant
	TicketID  int
	MessageID string
}

func (lro *LongRunningOperationsService) SendTicketCreationAcknowledgementActivity(ctx context.Context, req SendTicketCreationAcknowledgementRequest) (err error) {
	_, err = lro.nylasClient.SendMessage(&nylas.SendMessageRequest{
		Subject: "[Ticketer] New Ticket Created",
		Body:    "New Ticket Created",
		To: []nylas.Participant{
			req.Initiator,
		},
		ReplyToMessageID: req.MessageID,
	})
	return
}
