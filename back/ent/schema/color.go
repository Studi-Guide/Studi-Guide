package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type Color struct {
	ent.Schema
}

// Fields of the Room.
func (Color) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Color").Default(""),
	}
}

func (Color) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Color) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
