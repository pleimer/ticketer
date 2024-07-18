package repositories

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/ent"
	"github.com/pleimer/ticketer/server/ent/ticket"
)

// TODO: delete this

type TicketsRepository struct {
	db *db.DB
}

func NewTicketsRepository(db *db.DB) *TicketsRepository {
	return &TicketsRepository{
		db,
	}
}

func (r *TicketsRepository) CreateNewTicket(ctx context.Context) (ticket *ent.Ticket, err error) {
	// return r.db.Client.Ticket.
	// 	Create().
	// 	SetPriority(0).
	// 	SetThreadID("").
	// 	SetTitle("").
	// 	SetStatus(0).
	// 	Save(ctx)
	return nil, nil
}

func (r *TicketsRepository) GetTicketByID(ctx context.Context, ID int) (*ent.Ticket, error) {
	res, err := r.db.Client.Ticket.
		Query().
		Where(ticket.ID(ID)).
		Only(ctx)

	if err != nil {
		return nil, errors.Wrapf(err, "querying for ticket id %d", ID)
	}
	return res, nil
}
