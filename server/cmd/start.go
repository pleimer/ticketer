package cmd

import (
	"net"

	"github.com/labstack/echo/v4"

	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"

	"go.uber.org/zap"

	"github.com/pleimer/ticketer/server/app"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/services"
)

type Start struct {
	db.DBConnectionConfig
}

func (s *Start) Execute(args []string) error {

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
			app.App().Logger().Info("request",
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

	d := db.NewDB(app.App().Logger())

	d.Open(s.DBConnectionConfig)
	defer d.Close()

	return nil
}
