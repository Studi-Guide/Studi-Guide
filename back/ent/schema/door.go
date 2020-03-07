package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Door holds the schema definition for the Door entity.
type Door struct {
	ent.Schema
}

// Fields of the Door.
func (Door) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Id").Unique(),
	}
}

// Edges of the Door.
func (Door) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sequence", Sequence.Type).Unique(),
		edge.To("pathNodes", PathNode.Type).Unique(),
	}
}
