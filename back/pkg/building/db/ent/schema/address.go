package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("Street").NotEmpty(),
		field.String("Number").NotEmpty(),
		field.Int("PLZ").NonNegative(),
		field.String("City").NotEmpty(),
		field.String("Country").NotEmpty(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("campus", Campus.Type).Ref("address"),
	}
}

func (Address) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Street", "Number", "PLZ").
			Unique(),
	}
}
