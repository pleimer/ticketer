package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("assignee").
			Optional().Nillable(),
		field.String("opened_by"),
		field.Enum("status").
			Values("not_started", "in_progress", "done").
			Default("not_started"),
		field.Enum("priority").
			Values("low", "medium", "high").
			Default("low"),
		field.String("threadID"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return nil
}

// Edges of the Ticket.
func (Ticket) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("threadID"),
	}
}
