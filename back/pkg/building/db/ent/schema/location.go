package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Description").Default(""),
		field.String("Floor").Default("0"),
	}
}

// Edges of the Location.
func (Location) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pathnode", PathNode.Type).Unique().Required(),
		edge.From("tags", Tag.Type).Ref("locations"),
		edge.To("building", Building.Type).Unique(),
		edge.To("images", File.Type),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
