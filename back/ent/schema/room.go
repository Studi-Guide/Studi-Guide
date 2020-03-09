package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Description").Default(""),
		field.Int("Floor").Default(0),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doors", Door.Type).Required(),
		edge.To("sequences", Sequence.Type),
		edge.To("pathNode", PathNode.Type),
		edge.To("color", Color.Type).Unique(),
	}
}

func (Room) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
