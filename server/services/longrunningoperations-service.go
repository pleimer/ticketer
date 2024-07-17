package services

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/pleimer/ticketer/server/integration/nylas"
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
		MaximumInterval:    100 * time.Second,
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

	// No need to wait on update message status. If this fails, new messages that have already
	// been processed will be filtered our in the process messages activity
	workflow.ExecuteActivity(ctx, lro.UpdateMessageReadStatusActivity)

	var processNewMessagesRes []SendTicketCreationAcknowledgementRequest
	err = workflow.ExecuteActivity(ctx, lro.ProcessNewMessagesActivity, newMessagesResp).Get(ctx, &processNewMessagesRes)
	if err != nil {
		return err
	}

	// TODO: this needs to retry on failure, but only for messages messages that were not successfully replied to already
	return workflow.ExecuteActivity(ctx, lro.SendTicketCreationAcknowledgementActivity, processNewMessagesRes).Get(ctx, nil)
}

func (lro *LongRunningOperationsService) QueryNewMessagesActivity(ctx context.Context) (*nylas.MessagesResponse, error) {

	// TODO: paging
	msgs, err := lro.nylasClient.GetUnreadMessages(5)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving messages from nylas")
	}

	// mark messages as read

	return msgs, nil
}

// UpdateMessageReadStatusActivity updates status of messages to 'read'
func (lro *LongRunningOperationsService) UpdateMessageReadStatusActivity(ctx context.Context, msgs *nylas.MessagesResponse) error {

	if msgs == nil {
		return nil
	}

	for _, d := range msgs.Data {
		_, err := lro.nylasClient.UpdateMessageReadStatus(d.ID, false)
		if err != nil {
			return err
		}
	}

	return nil
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

func (lro *LongRunningOperationsService) SendTicketCreationAcknowledgementActivity(ctx context.Context, reqs []SendTicketCreationAcknowledgementRequest) error {

	for _, r := range reqs {
		_, err := lro.nylasClient.SendMessage(&nylas.SendMessageRequest{
			Subject: "[Ticketer] New Ticket Created",
			Body:    "New Ticket Created",
			ReplyTo: []nylas.Participant{
				r.Initiator,
			},
			ReplyToMessageID: r.MessageID,
		})
		if err != nil {
			// TODO: handle specific error codes
			return err
		}
	}

	return nil
}
