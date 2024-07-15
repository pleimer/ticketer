package app

import (
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"github.com/pleimer/ticketer/server/services"

	"github.com/labstack/echo/v4"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/repositories"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// App contains all application singletons and lazy loads the dependancy tree
var App func() *app

type app struct {
	*loggerConfig
	*dbConfig
	*repositoriesConfig
	*routerConfig
}

func init() {
	var a *app

	loggerConfig := loggerConfig{}
	dbConfig := dbConfig{}
	repositoriesConfig := repositoriesConfig{}
	routerConfig := routerConfig{}

	// setup the singleton dependancy tree
	loggerConfig.init()
	dbConfig.init(&loggerConfig)
	repositoriesConfig.init(&dbConfig)
	routerConfig.init(&loggerConfig)

	App = func() *app {
		once.Once(func() {
			a = &app{&loggerConfig, &dbConfig, &repositoriesConfig, &routerConfig}
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

type routerConfig struct {
	router *echo.Echo
	Router func() *echo.Echo
}

func (r *routerConfig) init(loggerConfig *loggerConfig) {
	r.Router = func() *echo.Echo {
		once.Once(func() {

			swagger, err := services.GetSwagger()
			if err != nil {
				panic(err)
			}

			swagger.Servers = nil

			ticketsService := services.NewTickets()

			r.router = echo.New()

			r.router.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
				LogURI:    true,
				LogStatus: true,
				LogValuesFunc: func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
					loggerConfig.Logger().Info("request",
						zap.String("URI", v.URI),
						zap.Int("status", v.Status),
					)

					return nil
				},
			}))

			// might as well validate the requests agains the schema
			r.router.Use(middleware.OapiRequestValidator(swagger))

			services.RegisterHandlers(r.router, ticketsService)

		})
		return r.router
	}
}
