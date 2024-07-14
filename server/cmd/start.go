package cmd

import (
	"net"

	"github.com/labstack/echo/v4"

	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/services"
)

// TODO: move
func initLogger() (*zap.Logger, error) {
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

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

type Start struct {
	db.DBConnectionConfig
}

func (s *Start) Execute(args []string) error {

	// logger

	logger, err := initLogger()
	if err != nil {
		panic(err)
	}

	// routes

	swagger, err := services.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil

	ticketsService := services.NewTickets()

	e := echo.New()

	e.Use(echomiddleware.RequestLoggerWithConfig(echomiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v echomiddleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	// might as well validate the requests agains the schema
	e.Use(middleware.OapiRequestValidator(swagger))

	services.RegisterHandlers(e, ticketsService)

	e.Logger.Fatal(e.Start(net.JoinHostPort("0.0.0.0", "8080")))

	d := db.NewDB(logger)
	d.Open(s.DBConnectionConfig)
	defer d.Close()

	return nil
}
