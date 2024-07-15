package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.Int("status"),
		field.String("title"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return nil
}
