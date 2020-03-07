package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default("unknown"),
		field.Int("Floor").Default(0),
		field.Int("Id").Unique(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doors", Door.Type),
		edge.To("sequences", Sequence.Type),
		edge.To("pathNodes", PathNode.Type),
	}
}
