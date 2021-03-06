package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// PathNodes holds the schema definition for the PathNodes entity.
type PathNode struct {
	ent.Schema
}

// Fields of the PathNodes.
func (PathNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Int("X_Coordinate").Default(0),
		field.Int("Y_Coordinate").Default(0),
		field.Int("Z_Coordinate").Default(0),
	}
}

// Edges of the PathNodes.
func (PathNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("door", Door.Type).
			Ref("pathNode").
			Unique(),

		edge.From("mapitem", MapItem.Type).
			Ref("pathNodes").Unique(),

		edge.From("location", Location.Type).
			Ref("pathnode").Unique(),

		edge.To("linkedTo", PathNode.Type).
			From("linkedFrom"),

		edge.From("pathGroups", PathNodeGroup.Type).Ref("pathNodes"),
	}
}
