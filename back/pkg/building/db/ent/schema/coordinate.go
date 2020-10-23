package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// PathNodes holds the schema definition for the PathNodes entity.
type Coordinate struct {
	ent.Schema
}

// Fields of the PathNodes.
func (Coordinate) Fields() []ent.Field {
	return []ent.Field{
		field.Float("Latitude").Default(0),
		field.Float("Longitude").Default(0),
	}
}

// Edges of the PathNodes.
func (Coordinate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("building", Building.Type).
			Ref("body").Unique(),
	}
}
