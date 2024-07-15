package app

import (
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/services"
	"go.uber.org/zap"
)

type routerConfig struct {
	router *echo.Echo
	Router func() *echo.Echo
}

func (r *routerConfig) init(loggerConfig *loggerConfig, servicesConfig *servicesConfig) {
	r.Router = func() *echo.Echo {
		once.Once(func() {

			swagger, err := services.GetSwagger()
			if err != nil {
				panic(err)
			}

			swagger.Servers = nil

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

			services.RegisterHandlers(r.router, servicesConfig.TicketsService())

		})
		return r.router
	}
}
