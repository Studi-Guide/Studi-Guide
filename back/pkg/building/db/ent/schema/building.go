package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Building holds the schema definition for the Building entity.
type Building struct {
	ent.Schema
}

// Fields of the Building.
func (Building) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Campus"),
	}
}

// Edges of the Building.
func (Building) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("mapitems", MapItem.Type).Ref("building"),
		edge.From("location", MapItem.Type).Ref("building"),
	}
}

func (Building) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
