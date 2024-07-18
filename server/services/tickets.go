package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pleimer/ticketer/server/db"
	"github.com/pleimer/ticketer/server/ent/ticket"
	"github.com/pleimer/ticketer/server/models"
	"go.uber.org/zap"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-models.cfg.yaml ../../internal/api/tickets.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=tickets-service.cfg.yaml ../../internal/api/tickets.json

type Tickets struct {
	logger *zap.Logger
	db     *db.DB
}

func NewTickets(db *db.DB, logger *zap.Logger) *Tickets {
	return &Tickets{logger, db}
}

// List Tickets
// (GET /tickets)
func (t *Tickets) ListTicket(ctx echo.Context, params models.ListTicketParams) error {
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
	status := ctx.QueryParam("status")
	if status != "" {
		query = query.Where(ticket.StatusEQ(ticket.Status(status)))
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
	var response []models.Ticket
	for _, t := range tickets {
		response = append(response, models.Ticket{
			Assignee: t.Assignee,
			Id:       t.ID,
			Priority: models.TicketPriority(t.Priority),
			Status:   models.TicketStatus(t.Status),
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
	panic("not implemented") // TODO: Implement
}

// Updates a Ticket
// (PATCH /tickets/{id})
func (t *Tickets) UpdateTicket(ctx echo.Context, id int) error {
	panic("not implemented") // TODO: Implement
}
