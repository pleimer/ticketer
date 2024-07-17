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

	env.App().DB().Open(m.DBConnectionConfig)
	defer env.App().DB().Close()

	if err := env.App().DB().Client.Schema.Create(context.Background()); err != nil {
		env.App().Logger().Fatal("creating schema", zap.Error(err))
	}

	return nil
}
