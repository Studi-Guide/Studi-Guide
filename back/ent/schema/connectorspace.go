package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// ConnectorSpace holds the schema definition for the ConnectorSpace entity.
type ConnectorSpace struct {
	ent.Schema
}

// Fields of the Room.
func (ConnectorSpace) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Description").Default(""),
		field.Int("Floor").Default(0),
	}
}

// Edges of the Room.
func (ConnectorSpace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connectorDoors", Door.Type),
		edge.To("connectorSections", Section.Type),
		edge.To("connectorPathNodes", PathNode.Type).Required(),
		edge.To("connectorColor", Color.Type).Unique(),
	}
}

func (ConnectorSpace) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
