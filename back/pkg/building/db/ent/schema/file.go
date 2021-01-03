package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").Unique(),
		field.String("Path").Unique().NotEmpty(),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
		index.Fields("Path").Unique(),
	}
}
