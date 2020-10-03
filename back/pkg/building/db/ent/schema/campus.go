package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Campus holds the schema definition for the Campus entity.
type Campus struct {
	ent.Schema
}

// Fields of the Campus.
func (Campus) Fields() []ent.Field {
	return []ent.Field{
		field.String("ShortName"),
		field.String("Name").Unique(),
		field.Float("Longitude"),
		field.Float("Latitude"),
	}
}

// Edges of the Campus.
func (Campus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).Unique().Required(),
	}
}

// Indexes of the Campus
func (Campus) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
