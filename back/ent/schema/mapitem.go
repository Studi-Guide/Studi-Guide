package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Room holds the schema definition for the Room entity.
type MapItem struct {
	ent.Schema
}

// Fields of the Room.
func (MapItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Floor").Default(0),
	}
}

// Edges of the Room.
func (MapItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doors", Door.Type),
		edge.To("sections", Section.Type),
		edge.To("pathNodes", PathNode.Type),
		edge.To("color", Color.Type).Unique(),
		edge.From("room", Room.Type).Ref("mapitem"),
		edge.To("building", Building.Type).Unique().Required(),
	}
}

func (MapItem) Indexes() []ent.Index {
	return []ent.Index{}
}
