package services

import "github.com/labstack/echo/v4"

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-models.cfg.yaml ../../internal/api/tickets.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-service.cfg.yaml ../../internal/api/tickets.yaml

type Tickets struct {
}

func NewTickets() *Tickets {
	return &Tickets{}
}

// Create a new ticket
// (POST /tickets)
func (t *Tickets) PostTickets(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Add a comment to a ticket
// (POST /tickets/{ticketId}/comments)
func (t *Tickets) PostTicketsTicketIdComments(ctx echo.Context, ticketId int) error {
	panic("not implemented") // TODO: Implement
}

// Update the status of a ticket
// (PUT /tickets/{ticketId}/status)
func (t *Tickets) PutTicketsTicketIdStatus(ctx echo.Context, ticketId int) error {

	panic("not implemented") // TODO: Implement

}

// Get ticket
// (GET /tickets/{ticketId})
func (t *Tickets) GetTicket(ctx echo.Context, ticketId string) error {
	panic("not implemented") // TODO: Implement
}
