package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PathNode holds the schema definition for the PathNode entity.
type PathNode struct {
	ent.Schema
}

// Fields of the PathNode.
func (PathNode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Id").Unique(),
		field.Int("X_Coordinate").Default(0),
		field.Int("Y_Coordinate").Default(0),
		field.Int("Floor").Default(0),
	}
}

// Edges of the PathNode.
func (PathNode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("door", Door.Type).
			Ref("pathNodes").
			Unique(),

		edge.From("room", Room.Type).
			Ref("pathNodes").
			Unique(),

		edge.To("linkedTo", PathNode.Type).
			From("linkedFrom"),
	}
}
