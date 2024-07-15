package app

import (
	"github.com/pleimer/ticketer/server/lib/once"
	"github.com/pleimer/ticketer/server/repositories"
)

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
