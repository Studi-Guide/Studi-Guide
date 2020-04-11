package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Section holds the schema definition for the Section entity.
type Section struct {
	ent.Schema
}

// Fields of the Section.
func (Section) Fields() []ent.Field {
	return []ent.Field{
		field.Int("X_Start").Default(0),
		field.Int("Y_Start").Default(0),
		field.Int("X_End").Default(0),
		field.Int("Y_End").Default(0),
		field.Int("Z_Start").Default(0),
		field.Int("Z_End").Default(0),
	}
}

// Edges of the Section.
func (Section) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("door", Door.Type).
			Ref("section").
			Unique(),

		edge.From("room", Room.Type).
			Ref("sections").
			Unique(),
	}
}
