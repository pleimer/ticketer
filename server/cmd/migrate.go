package cmd

import (
	"context"

	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/env"
	"go.uber.org/zap"
)

// Migrate command for versioned migrations in production. For this
// excercise, just use auto-migration
type Migrate struct {
	db.DBConnectionConfig
}

func (m *Migrate) Execute(args []string) error {

	app := env.NewEnv()
	defer app.Cleanup()

	app.DBConnectionConfig = m.DBConnectionConfig

	if err := app.DB().Client.Schema.Create(context.Background()); err != nil {
		app.Logger().Fatal("creating schema", zap.Error(err))
	}

	return nil
}
