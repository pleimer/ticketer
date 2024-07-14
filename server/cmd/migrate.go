package cmd

import (
	"context"

	"github.com/pleimer/ticketer/server/db"
	"go.uber.org/zap"
)

type Migrate struct {
	db.DBConnectionConfig
}

func (m *Migrate) Execute(args []string) error {

	logger, err := initLogger()
	if err != nil {
		panic(err)
	}

	d := db.NewDB(logger)
	d.Open(m.DBConnectionConfig)
	defer d.Close()

	if err := d.Cli.Schema.Create(context.Background()); err != nil {
		logger.Fatal("creating schema", zap.Error(err))
	}

	return nil
}
