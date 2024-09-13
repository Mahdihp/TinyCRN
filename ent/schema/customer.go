package schema

import "entgo.io/ent"

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
	return nil
}
