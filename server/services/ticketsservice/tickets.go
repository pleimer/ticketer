package ticketsservice

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/ent"
	"github.com/pleimer/ticketer/server/ent/ticket"
	"go.uber.org/zap"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-models.cfg.yaml ../../../internal/api/tickets.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-service.cfg.yaml ../../../internal/api/tickets.json

type Tickets struct {
	logger *zap.Logger
	db     *db.DB
}

func NewTickets(db *db.DB, logger *zap.Logger) *Tickets {
	return &Tickets{logger, db}
}

// List Tickets
// (GET /tickets)
func (t *Tickets) ListTicket(ctx echo.Context, params ListTicketParams) error {
	// Get query parameters
	page := params.Page
	if page == nil || *page < 1 {
		defaultPage := 1
		page = &defaultPage
	}

	itemsPerPage := params.ItemsPerPage
	if itemsPerPage == nil || *itemsPerPage < 1 || *itemsPerPage > 255 {
		defaultItemsPerPage := 20
		itemsPerPage = &defaultItemsPerPage
	}

	// Calculate offset
	offset := (*page - 1) * *itemsPerPage

	// Start building the query
	query := t.db.Client.Ticket.Query()

	// Add status filter if provided
	if params.Status != nil {
		query = query.Where(ticket.StatusEQ(ticket.Status(*params.Status)))
	}

	// Execute the query with pagination
	tickets, err := query.
		Limit(*itemsPerPage).
		Offset(offset).
		All(ctx.Request().Context())

	if err != nil {
		t.logger.Error("failed quering for tickets", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch tickets")
	}

	// Convert ent.Ticket to models.Ticket
	var response []Ticket

	for _, t := range tickets {
		response = append(response, Ticket{
			Assignee: t.Assignee,
			Id:       t.ID,
			Priority: TicketPriority(t.Priority),
			Status:   TicketStatus(t.Status),
			ThreadID: t.ThreadID,
			Title:    t.Title,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

// Create a new Ticket
// (POST /tickets)
func (t *Tickets) CreateTicket(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Deletes a Ticket by ID
// (DELETE /tickets/{id})
func (t *Tickets) DeleteTicket(ctx echo.Context, id int) error {
	panic("not implemented") // TODO: Implement
}

// Find a Ticket by ID
// (GET /tickets/{id})
func (t *Tickets) ReadTicket(ctx echo.Context, id int) error {

	ticket, err := t.db.Client.Ticket.Get(ctx.Request().Context(), id)
	if err != nil {
		t.logger.Error("querying ticket", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed retrieving ticket")
	}
	return ctx.JSON(http.StatusOK, ticket)
}

// Updates a Ticket
// (PATCH /tickets/{id})
func (t *Tickets) UpdateTicket(ctx echo.Context, id int) error {

	var updateData UpdateTicketJSONBody

	if err := ctx.Bind(&updateData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	// Start a transaction
	tx, err := t.db.Client.Tx(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to start transaction")
	}
	defer tx.Rollback()

	// Get the ticket
	ticketToUpdate, err := tx.Ticket.Query().Where(ticket.ID(id)).Only(ctx.Request().Context())
	if err != nil {
		if ent.IsNotFound(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Ticket not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch ticket")
	}

	// Update the ticket
	update := ticketToUpdate.Update()

	if updateData.Title != nil {
		update = update.SetTitle(*updateData.Title)
	}
	if updateData.Assignee != nil {
		update = update.SetAssignee(*updateData.Assignee)
	}
	if updateData.Status != nil {
		update = update.SetStatus(ticket.Status(*updateData.Status))
	}
	if updateData.Priority != nil {
		update = update.SetPriority(ticket.Priority(*updateData.Priority))
	}
	if updateData.ThreadID != nil {
		update = update.SetThreadID(*updateData.ThreadID)
	}
	if updateData.OpenedBy != nil {
		update = update.SetCreatedBy(*updateData.OpenedBy)
	}

	// Save the changes
	updatedTicket, err := update.Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update ticket")
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to commit transaction")
	}

	// Return the updated ticket
	return ctx.JSON(http.StatusOK, updatedTicket)

}
