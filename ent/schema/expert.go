package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.String("username").Unique().MaxLen(20),
		field.String("first_name").MaxLen(50),
		field.String("last_name").MaxLen(50),
		field.Time("start_time").Default(time.Now),
		field.Time("end_time").Default(time.Now),
		field.Bool("is_online"),
	}
}

// Edges of the Expert.
func (Expert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("department", Department.Type),
	}
}
