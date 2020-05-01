package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PathNodeGroup holds the schema definition for the PathNodeGroup entity.
type PathNodeGroup struct {
	ent.Schema
}

// Fields of the PathNodeGroup.
func (PathNodeGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
	}
}

// Edges of the PathNodeGroup.
func (PathNodeGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pathNodes", PathNode.Type),
	}
}
