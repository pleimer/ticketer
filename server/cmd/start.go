package cmd

import (
	"context"
	"fmt"
	"net"

	"github.com/pleimer/ticketer/server/app"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/integration/nylas"
)

type Start struct {
	db.DBConnectionConfig
	nylas.NylasClientConfig
}

func (s *Start) Execute(args []string) error {

	app.App().NylasClientConfig = s.NylasClientConfig

	app.App().DB().Open(s.DBConnectionConfig)
	defer app.App().DB().Close()

	r, err := app.App().NylasClient().ListThreadMessages(context.Background(), "AQQkADAwATNiZmYAZS04MTEAMi0zNGVjLTAwAi0wMAoAEADwPf5pU6GwRKQJc6h3MguA")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n\n", r)

	app.App().Logger().Sugar().Fatal(
		app.App().Router().Start(net.JoinHostPort("0.0.0.0", "8080")),
	)

	return nil
}
