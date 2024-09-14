package schema

import (
	"TinyCRM/shared"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title").MaxLen(shared.MAX_LEN_50),

		field.Time("content").
			SchemaType(map[string]string{
				dialect.Postgres: "text", // Override Postgres.
			}),

		field.Enum("status").
			Values(shared.TicketStatus_Unknown, shared.TicketStatus_Pending, shared.TicketStatus_Referenced, shared.TicketStatus_Answered, shared.TicketStatus_Closed),

		field.String("ticket_file_id").Nillable(),
		field.Bool("is_viewed").Default(false),

		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now).Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Ticket) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer_id", Customer.Type).Ref("tickets").Unique().Required(),
		edge.From("expert_id", Expert.Type).Ref("tickets").Unique().Required(),
		edge.From("department_id", Department.Type).Ref("tickets").Unique().Required(),
	}
}
