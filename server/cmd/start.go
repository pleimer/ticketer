package cmd

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/env"
	"github.com/pleimer/ticketer/server/integration/nylas"
	"github.com/pleimer/ticketer/server/services/ticketsservice"
	"go.uber.org/zap"
)

type Start struct {
	ticketsservice.TemporalConfig
	db.DBConnectionConfig
	nylas.NylasClientConfig
}

func (s *Start) Execute(args []string) error {

	app := env.NewEnv()

	// Usually, config should be applied based on the type of environment being instantiated (stg, prod, test, dev).
	// For this project, since there is only one environment, will just apply configurations here
	app.NylasClientConfig = s.NylasClientConfig
	app.DBConnectionConfig = s.DBConnectionConfig
	app.TemporalConfig = s.TemporalConfig

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		app.Cleanup()
	}()

	// run the db-migration for developer convenience. In production env, migrations would be
	// versioned and run as a pre-deployment step
	m := Migrate{
		true,
		s.DBConnectionConfig,
	}
	err := m.Execute(nil)
	if err != nil {
		app.Logger().Fatal("running auto schema migration", zap.Error(err))
	}

	// Starts a temporal workflow that polls the nylas API. Typically, we would
	app.LongRunningOperationsService()

	app.Logger().Sugar().Error(
		app.Router().Start(net.JoinHostPort("0.0.0.0", "8080")),
	)

	return nil
}
