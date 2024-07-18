package env

import (
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/services/threadsservice"
	"github.com/pleimer/ticketer/server/services/ticketsservice"
)

type servicesConfig struct {
	threadsService *threadsservice.ThreadsService
	ThreadsService func() *threadsservice.ThreadsService

	ticketsService *ticketsservice.Tickets
	TicketsService func() *ticketsservice.Tickets

	longRunningOperationsService *ticketsservice.LongRunningOperationsService
	LongRunningOperationsService func() *ticketsservice.LongRunningOperationsService
}

func (s *servicesConfig) init(loggerConfig *loggerConfig, repositoriesConfig *repositoriesConfig, integrationsConfig *integrationsConfig, dbConfig *dbConfig) {
	s.TicketsService = func() *ticketsservice.Tickets {
		once.Once(func() {
			s.ticketsService = ticketsservice.NewTickets(dbConfig.DB(), loggerConfig.Logger())
		})
		return s.ticketsService
	}

	s.ThreadsService = func() *threadsservice.ThreadsService {
		once.Once(func() {
			s.threadsService = threadsservice.NewThreadsService(loggerConfig.Logger(), integrationsConfig.NylasClient())
		})
		return s.threadsService
	}

	s.LongRunningOperationsService = func() *ticketsservice.LongRunningOperationsService {
		once.Once(func() {
			s.longRunningOperationsService = ticketsservice.NewLongRunningOperationsService(loggerConfig.Logger(), integrationsConfig.NylasClient(), dbConfig.DB(), repositoriesConfig.TicketsRepository())
		})
		return s.longRunningOperationsService
	}
}
