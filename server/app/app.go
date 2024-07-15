package app

import (
	"fmt"

	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/repositories"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var App func() *app

type app struct {
	*loggerConfig
	*dbConfig
	*repositoriesConfig
}

func init() {
	var a *app

	loggerConfig := loggerConfig{}
	dbConfig := dbConfig{}
	repositoriesConfig := repositoriesConfig{}

	// setup the singleton dependancy tree
	loggerConfig.init()
	dbConfig.init(&loggerConfig)
	repositoriesConfig.init(&dbConfig)

	App = func() *app {
		once.Once(func() {
			a = &app{&loggerConfig, &dbConfig, &repositoriesConfig}
		})
		return a
	}
}

type loggerConfig struct {
	logger *zap.Logger
	Logger func() *zap.Logger
}

func (l *loggerConfig) init() {
	l.Logger = func() *zap.Logger {

		once.Once(func() {
			fmt.Println("initialized logger")
			config := zap.Config{
				Encoding:         "json",
				Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
				OutputPaths:      []string{"stdout"},
				ErrorOutputPaths: []string{"stderr"},
				EncoderConfig: zapcore.EncoderConfig{
					MessageKey:   "message",
					LevelKey:     "level",
					TimeKey:      "time",
					CallerKey:    "caller",
					EncodeCaller: zapcore.ShortCallerEncoder,
					EncodeLevel:  zapcore.CapitalLevelEncoder,
					EncodeTime:   zapcore.ISO8601TimeEncoder,
				},
			}

			var err error

			l.logger, err = config.Build()
			if err != nil {
				panic(err)
			}
		})

		return l.logger
	}
}

type dbConfig struct {
	db *db.DB
	DB func() *db.DB
}

func (d *dbConfig) init(loggerConfig *loggerConfig) {
	d.DB = func() *db.DB {
		once.Once(func() {
			fmt.Println("initialized db")
			d.db = db.NewDB(
				loggerConfig.Logger(),
			)
		})
		return d.db
	}
}

type repositoriesConfig struct {
	ticketsRepository *repositories.TicketsRepository
	TicketsRepository func() *repositories.TicketsRepository
}

func (r *repositoriesConfig) init(dbConfig *dbConfig) {
	r.TicketsRepository = func() *repositories.TicketsRepository {
		once.Once(func() {
			r.ticketsRepository = repositories.NewTicketsRepository(
				dbConfig.DB(),
			)
		})
		return r.ticketsRepository
	}
}
