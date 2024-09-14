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

// Expert holds the schema definition for the Expert entity.
type Expert struct {
	ent.Schema
}

// Fields of the Expert.
func (Expert) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("username").Unique().MaxLen(shared.MAX_LEN_20),

		field.String("first_name").MaxLen(shared.MAX_LEN_50),
		field.String("last_name").MaxLen(shared.MAX_LEN_50),

		field.Time("start_time").Default(time.Now),
		field.Time("end_time").Default(time.Now),
		field.Bool("is_online"),

		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp(6)", // Override Postgres.
			}).Default(time.Now).Immutable(),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

func (Expert) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the Expert.
func (Expert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("departments", Department.Type).Ref("experts"),
		edge.To("tickets", Ticket.Type),
	}
}
