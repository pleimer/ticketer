package services

import (
	"context"
	"time"

	"github.com/pleimer/ticketer/server/integration/nylas"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

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

	return &LongRunningOperationsService{
		client:      temporalClient,
		nylasClient: nylas,
		logger:      logger,
	}
}

func (lro *LongRunningOperationsService) Start() {
	opts := client.StartWorkflowOptions{
		ID:           "email-ingestor-workflow",
		TaskQueue:    "email-ingestor-taskqueue",
		CronSchedule: "*/1 * * * *", // Run every 5 minutes
	}

	we, err := lro.client.ExecuteWorkflow(context.Background(), opts, lro.EmailIngestorWorkflow)
	if err != nil {
		lro.logger.Sugar().Fatalf("Unabe to execute workflow: %v", err)
	}

	lro.logger.Sugar().Infof("Started workflow: WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

}

func (lro *LongRunningOperationsService) Close() {
	lro.client.Close()
}

func (lro *LongRunningOperationsService) EmailIngestorWorkflow(ctx workflow.Context) (string, error) {

	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"InvalidAccountError", "InsufficientFundsError"},
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var res string

	err := workflow.ExecuteActivity(ctx, lro.QueryNewMessagesActivity).Get(ctx, &res)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (lro *LongRunningOperationsService) QueryNewMessagesActivity(ctx context.Context) (string, error) {

	lro.nylasClient.ListThreadMessages("")

	return "", nil
}
