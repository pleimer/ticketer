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

	env.App().NylasClientConfig = s.NylasClientConfig

	env.App().DB().Open(s.DBConnectionConfig)
	defer env.App().DB().Close()

	// r, err := app.App().NylasClient().ListThreadMessages(context.Background(), "AQQkADAwATNiZmYAZS04MTEAMi0zNGVjLTAwAi0wMAoAEADwPf5pU6GwRKQJc6h3MguA")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n\n", r)
	env.App().LongRunningOperationsService().Start()
	defer env.App().LongRunningOperationsService().Close()

	env.App().Logger().Sugar().Fatal(
		env.App().Router().Start(net.JoinHostPort("0.0.0.0", "8080")),
	)

	return nil
}
