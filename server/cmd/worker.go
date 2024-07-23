package cmd

import (
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/env"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"github.com/pleimer/ticketer/server/services/ticketsservice"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
)

type RunWorker struct {
	ticketsservice.TemporalConfig
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
	app.TemporalConfig = r.TemporalConfig

	c, err := client.Dial(client.Options{
		HostPort: r.TemporalConfig.HostPort,
	})
	if err != nil {
		app.Logger().Sugar().Fatalf("Unable to create Temporal client.", err)
	}
	defer c.Close()

	// TODO: cleanup registration, can be registered with obj
	w := worker.New(c, "email-ingestor-taskqueue", worker.Options{})
	w.RegisterWorkflow(app.LongRunningOperationsService().EmailIngestorWorkflow)
	w.RegisterWorkflow(app.LongRunningOperationsService().TicketCreationAcknowledgementChildWorkflow)
	w.RegisterActivity(app.LongRunningOperationsService().QueryNewMessagesActivity)
	w.RegisterActivity(app.LongRunningOperationsService().UpdateMessageReadStatusActivity)
	w.RegisterActivity(app.LongRunningOperationsService().ProcessNewMessagesActivity)
	w.RegisterActivity(app.LongRunningOperationsService().SendTicketCreationAcknowledgementActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		app.Logger().Fatal("starting worker", zap.Error(err))
	}

	return nil
}
