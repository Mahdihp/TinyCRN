package schema

import (
	"TinyCRM/shared"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

//Financial = 1,
/// <summary>
/// فنی
/// </summary>
//Technical = 2,
/// <summary>
/// روابط عمومی
/// </summary>
//PublicRelations = 3,

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").MaxLen(shared.MAX_LEN_50),
	}
}

func (Department) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("experts", Expert.Type),
		edge.To("tickets", Ticket.Type),
	}
}
