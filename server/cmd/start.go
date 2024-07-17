package cmd

import (
	"net"

	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/env"
	"github.com/pleimer/ticketer/server/integration/nylas"
)

type Start struct {
	db.DBConnectionConfig
	nylas.NylasClientConfig
}

func (s *Start) Execute(args []string) error {

	app := env.NewEnv()
	defer app.Cleanup()

	// Usually, config should be applied based on the type of environment being instantiated (stg, prod, test, dev).
	// For this project, since there is only one environment, will just apply configurations here
	app.NylasClientConfig = s.NylasClientConfig
	app.DBConnectionConfig = s.DBConnectionConfig

	// TODO: move this to a cleanup function in the env package

	app.Logger().Sugar().Fatal(
		app.Router().Start(net.JoinHostPort("0.0.0.0", "8080")),
	)

	return nil
}
