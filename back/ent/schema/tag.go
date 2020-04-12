package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type Tag struct {
	ent.Schema
}

// Fields of the Room.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
	}
}

// Edges of the Room.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("locations", Location.Type).Required(),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
