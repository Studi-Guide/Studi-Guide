package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Room holds the schema definition for the Room entity.
type Location struct {
	ent.Schema
}

// Fields of the Room.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Description").Default(""),
		field.String("Floor").Default("0"),
	}
}

// Edges of the Room.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pathnode", PathNode.Type).Unique().Required(),
		edge.From("tags", Tag.Type).Ref("locations"),
		edge.To("building", Building.Type).Unique(),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}