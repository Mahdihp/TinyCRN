package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return nil
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type),
	}
}
