package env

import (
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/services/messagesservice"
	"github.com/pleimer/ticketer/server/services/ticketsservice"
)

type servicesConfig struct {
	TemporalConfig ticketsservice.TemporalConfig

	threadsService *messagesservice.MessagesService
	ThreadsService func() *messagesservice.MessagesService

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

	s.ThreadsService = func() *messagesservice.MessagesService {
		once.Once(func() {
			s.threadsService = messagesservice.NewMessagesService(loggerConfig.Logger(), integrationsConfig.NylasClient())
		})
		return s.threadsService
	}

	s.LongRunningOperationsService = func() *ticketsservice.LongRunningOperationsService {
		once.Once(func() {
			s.longRunningOperationsService = ticketsservice.NewLongRunningOperationsService(s.TemporalConfig, loggerConfig.Logger(), integrationsConfig.NylasClient(), dbConfig.DB(), repositoriesConfig.TicketsRepository())
		})
		return s.longRunningOperationsService
	}
}
