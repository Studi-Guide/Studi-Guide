package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Sequence holds the schema definition for the Sequence entity.
type Sequence struct {
	ent.Schema
}

// Fields of the Sequence.
func (Sequence) Fields() []ent.Field {
	return []ent.Field{
		field.Int("X_Start").Default(0),
		field.Int("Y_Start").Default(0),
		field.Int("X_End").Default(0),
		field.Int("Y_End").Default(0),
		field.Int("Id").Unique(),
	}
}

// Edges of the Sequence.
func (Sequence) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("door", Door.Type).
			Ref("sequence").
			Unique(),

		edge.From("room", Room.Type).
			Ref("sequences").
			Unique(),
	}
}
