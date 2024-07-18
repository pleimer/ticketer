package env

import (
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/services"
)

type servicesConfig struct {
	ticketsService *services.Tickets
	TicketsService func() *services.Tickets

	longRunningOperationsService *services.LongRunningOperationsService
	LongRunningOperationsService func() *services.LongRunningOperationsService
}

func (s *servicesConfig) init(loggerConfig *loggerConfig, repositoriesConfig *repositoriesConfig, integrationsConfig *integrationsConfig, dbConfig *dbConfig) {
	s.TicketsService = func() *services.Tickets {
		once.Once(func() {
			s.ticketsService = services.NewTickets()
		})
		return s.ticketsService
	}

	s.LongRunningOperationsService = func() *services.LongRunningOperationsService {
		once.Once(func() {
			s.longRunningOperationsService = services.NewLongRunningOperationsService(loggerConfig.Logger(), integrationsConfig.NylasClient(), dbConfig.DB(), repositoriesConfig.TicketsRepository())
		})
		return s.longRunningOperationsService
	}
}
