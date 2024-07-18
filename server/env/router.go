package env

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/services/threadsservice"
	"github.com/pleimer/ticketer/server/services/ticketsservice"
	"go.uber.org/zap"
)

type routerConfig struct {
	router *echo.Echo
	Router func() *echo.Echo
}

func (r *routerConfig) init(loggerConfig *loggerConfig, servicesConfig *servicesConfig) {
	r.Router = func() *echo.Echo {
		once.Once(func() {

			ticketsSwagger, err := ticketsservice.GetSwagger()
			if err != nil {
				panic(err)
			}

			messagesSwagger, err := threadsservice.GetSwagger()
			if err != nil {
				panic(err)
			}

			ticketsSwagger.Servers = nil
			messagesSwagger.Servers = nil

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

			api := r.router.Group("/api/v1/")

			// Oapi swagger validator does not take group prefixes into account,
			// adapt the paths here
			fixSwaggerPrefix("/api/v1/tickets", ticketsSwagger)
			fixSwaggerPrefix("/api/v1/messages", messagesSwagger)

			tickets := api.Group("tickets", middleware.OapiRequestValidator(ticketsSwagger))
			messages := api.Group("messages", middleware.OapiRequestValidator(messagesSwagger))

			ticketsservice.RegisterHandlers(tickets, servicesConfig.TicketsService())
			threadsservice.RegisterHandlers(messages, servicesConfig.ThreadsService())
		})
		return r.router
	}
}

func fixSwaggerPrefix(prefix string, swagger *openapi3.T) {
	var updatedPaths openapi3.Paths = openapi3.Paths{}

	for key, value := range swagger.Paths.Map() {
		updatedPaths.Set(prefix+key, value)
	}

	swagger.Paths = &updatedPaths
}
