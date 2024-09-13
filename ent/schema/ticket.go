package schema

import "entgo.io/ent"

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return nil
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return nil
}
