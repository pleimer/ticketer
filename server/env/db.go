package env

import (
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/lib/once"
)

type dbConfig struct {
	db *db.DB
	DB func() *db.DB
}

func (d *dbConfig) init(loggerConfig *loggerConfig) {
	d.DB = func() *db.DB {
		once.Once(func() {
			d.db = db.NewDB(
				loggerConfig.Logger(),
			)
		})
		return d.db
	}
}
