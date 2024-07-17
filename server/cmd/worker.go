package cmd

import (
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/env"
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

	app := env.NewEnv()
	defer app.Cleanup()

	// Usually, config should be applied based on the type of environment being instantiated (stg, prod, test, dev).
	// For this project, since there is only one environment, will just apply configurations here
	app.NylasClientConfig = r.NylasClientConfig
	app.DBConnectionConfig = r.DBConnectionConfig

	c, err := client.Dial(client.Options{})
	if err != nil {
		app.Logger().Sugar().Fatalf("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, "email-ingestor-taskqueue", worker.Options{})
	w.RegisterWorkflow(app.LongRunningOperationsService().EmailIngestorWorkflow)
	w.RegisterActivity(app.LongRunningOperationsService().QueryNewMessagesActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		app.Logger().Fatal("starting worker", zap.Error(err))
	}

	return nil
}
