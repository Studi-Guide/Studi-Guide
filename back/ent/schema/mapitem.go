package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type MapItem struct {
	ent.Schema
}

// Fields of the Room.
func (MapItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Default(""),
		field.String("Description").Default(""),
		field.Int("Floor").Default(0),
	}
}

// Edges of the Room.
func (MapItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doors", Door.Type).Required(),
		edge.To("sections", Section.Type),
		edge.To("pathNodes", PathNode.Type).Required(),
		edge.To("color", Color.Type).Unique(),
		edge.From("tags", Tag.Type).Ref("mapitems"),
		edge.From("room", Room.Type).Ref("mapitem"),
	}
}

func (MapItem) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
