package cmd

import (
	"context"

	"github.com/pleimer/ticketer/server/app"
	"github.com/pleimer/ticketer/server/db"
	"go.uber.org/zap"
)

// Migrate command for versioned migrations in production. For this
// excercise, just use auto-migration
type Migrate struct {
	db.DBConnectionConfig
}

func (m *Migrate) Execute(args []string) error {

	d := db.NewDB(app.App().Logger())
	d.Open(m.DBConnectionConfig)
	defer d.Close()

	if err := d.Client.Schema.Create(context.Background()); err != nil {
		app.App().Logger().Fatal("creating schema", zap.Error(err))
	}

	return nil
}
