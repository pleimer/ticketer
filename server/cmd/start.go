package cmd

import (
	"net"

	"github.com/pleimer/ticketer/server/app"
	"github.com/pleimer/ticketer/server/db"
)

type Start struct {
	db.DBConnectionConfig
}

func (s *Start) Execute(args []string) error {

	// routes

	app.App().DB().Open(s.DBConnectionConfig)
	defer app.App().DB().Close()

	app.App().Logger().Sugar().Fatal(
		app.App().Router().Start(net.JoinHostPort("0.0.0.0", "8080")),
	)

	return nil
}
