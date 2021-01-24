package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type Icon struct {
	ent.Schema
}

// Fields of the Room.
func (Icon) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
	}
}

func (Icon) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Icon) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
