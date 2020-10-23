package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("mapitem", MapItem.Type).Unique().Required(),
		edge.To("location", Location.Type).Unique().Required(),
	}
}

func (Room) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
	}
}
