// Package ticketsservice provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package ticketsservice

// Defines values for TicketPriority.
const (
	TicketPriorityHigh   TicketPriority = "high"
	TicketPriorityLow    TicketPriority = "low"
	TicketPriorityMedium TicketPriority = "medium"
)

// Defines values for TicketStatus.
const (
	TicketStatusDone       TicketStatus = "done"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusNotStarted TicketStatus = "not_started"
)

// Defines values for ListTicketParamsStatus.
const (
	ListTicketParamsStatusDone       ListTicketParamsStatus = "done"
	ListTicketParamsStatusInProgress ListTicketParamsStatus = "in_progress"
	ListTicketParamsStatusNotStarted ListTicketParamsStatus = "not_started"
)

// Defines values for CreateTicketJSONBodyPriority.
const (
	CreateTicketJSONBodyPriorityHigh   CreateTicketJSONBodyPriority = "high"
	CreateTicketJSONBodyPriorityLow    CreateTicketJSONBodyPriority = "low"
	CreateTicketJSONBodyPriorityMedium CreateTicketJSONBodyPriority = "medium"
)

// Defines values for CreateTicketJSONBodyStatus.
const (
	CreateTicketJSONBodyStatusDone       CreateTicketJSONBodyStatus = "done"
	CreateTicketJSONBodyStatusInProgress CreateTicketJSONBodyStatus = "in_progress"
	CreateTicketJSONBodyStatusNotStarted CreateTicketJSONBodyStatus = "not_started"
)

// Defines values for UpdateTicketJSONBodyPriority.
const (
	High   UpdateTicketJSONBodyPriority = "high"
	Low    UpdateTicketJSONBodyPriority = "low"
	Medium UpdateTicketJSONBodyPriority = "medium"
)

// Defines values for UpdateTicketJSONBodyStatus.
const (
	UpdateTicketJSONBodyStatusDone       UpdateTicketJSONBodyStatus = "done"
	UpdateTicketJSONBodyStatusInProgress UpdateTicketJSONBodyStatus = "in_progress"
	UpdateTicketJSONBodyStatusNotStarted UpdateTicketJSONBodyStatus = "not_started"
)

// Ticket defines model for Ticket.
type Ticket struct {
	Assignee *string        `json:"assignee,omitempty"`
	Id       int            `json:"id"`
	OpenedBy string         `json:"opened_by"`
	Priority TicketPriority `json:"priority"`
	Status   TicketStatus   `json:"status"`
	ThreadID string         `json:"threadID"`
	Title    string         `json:"title"`
}

// TicketPriority defines model for Ticket.Priority.
type TicketPriority string

// TicketStatus defines model for Ticket.Status.
type TicketStatus string

// Resp400 defines model for resp400.
type Resp400 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// Resp404 defines model for resp404.
type Resp404 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// Resp409 defines model for resp409.
type Resp409 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// Resp500 defines model for resp500.
type Resp500 struct {
	Code   int          `json:"code"`
	Errors *interface{} `json:"errors,omitempty"`
	Status string       `json:"status"`
}

// ListTicketParams defines parameters for ListTicket.
type ListTicketParams struct {
	// Status ticket status filter
	Status *ListTicketParamsStatus `form:"status,omitempty" json:"status,omitempty"`

	// Page what page to render
	Page *int `form:"page,omitempty" json:"page,omitempty"`

	// ItemsPerPage item count to render per page
	ItemsPerPage *int `form:"itemsPerPage,omitempty" json:"itemsPerPage,omitempty"`
}

// ListTicketParamsStatus defines parameters for ListTicket.
type ListTicketParamsStatus string

// CreateTicketJSONBody defines parameters for CreateTicket.
type CreateTicketJSONBody struct {
	Assignee *string                      `json:"assignee,omitempty"`
	OpenedBy string                       `json:"opened_by"`
	Priority CreateTicketJSONBodyPriority `json:"priority"`
	Status   CreateTicketJSONBodyStatus   `json:"status"`
	ThreadID string                       `json:"threadID"`
	Title    string                       `json:"title"`
}

// CreateTicketJSONBodyPriority defines parameters for CreateTicket.
type CreateTicketJSONBodyPriority string

// CreateTicketJSONBodyStatus defines parameters for CreateTicket.
type CreateTicketJSONBodyStatus string

// UpdateTicketJSONBody defines parameters for UpdateTicket.
type UpdateTicketJSONBody struct {
	Assignee *string                       `json:"assignee,omitempty"`
	OpenedBy *string                       `json:"opened_by,omitempty"`
	Priority *UpdateTicketJSONBodyPriority `json:"priority,omitempty"`
	Status   *UpdateTicketJSONBodyStatus   `json:"status,omitempty"`
	ThreadID *string                       `json:"threadID,omitempty"`
	Title    *string                       `json:"title,omitempty"`
}

// UpdateTicketJSONBodyPriority defines parameters for UpdateTicket.
type UpdateTicketJSONBodyPriority string

// UpdateTicketJSONBodyStatus defines parameters for UpdateTicket.
type UpdateTicketJSONBodyStatus string

// CreateTicketJSONRequestBody defines body for CreateTicket for application/json ContentType.
type CreateTicketJSONRequestBody CreateTicketJSONBody

// UpdateTicketJSONRequestBody defines body for UpdateTicket for application/json ContentType.
type UpdateTicketJSONRequestBody UpdateTicketJSONBody