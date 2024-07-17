package cmd

import (
	"github.com/pleimer/ticketer/server/app"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

type RunWorker struct {
	db.DBConnectionConfig
	nylas.NylasClientConfig
}

func (r *RunWorker) Execute(args []string) error {

	c, err := client.Dial(client.Options{})
	if err != nil {
		app.App().Logger().Sugar().Fatalf("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, "email-ingestor-taskqueue", worker.Options{})
	w.RegisterWorkflow(app.App().LongRunningOperationsService().EmailIngestorWorkflow)
	w.RegisterActivity(app.App().LongRunningOperationsService().QueryNewMessagesActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		app.App().Logger().Fatal("starting worker", zap.Error(err))
	}

	return nil
}
