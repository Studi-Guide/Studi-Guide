package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// Door holds the schema definition for the Door entity.
type Door struct {
	ent.Schema
}

// Fields of the Door.
func (Door) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Door.
func (Door) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("section", Section.Type).Unique().Required(),
		edge.To("pathNode", PathNode.Type).Unique(),
		edge.From("owner", MapItem.Type).Ref("doors"),
	}
}
