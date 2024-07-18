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
		field.Int("status"),
		field.Int("priority"),
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
