package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title").MaxLen(50),
		field.String("content").MaxLen(500),
		field.Int("status"),
		field.String("ticket_file_id").Nillable(),
		field.Bool("is_viewed"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer_id", Customer.Type).Ref("tickets"),
		edge.From("expert_id", Expert.Type).Ref("tickets"),
		edge.From("department_id", Department.Type).Ref("tickets"),
	}
}
