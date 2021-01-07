package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"regexp"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("Path").Unique().NotEmpty().Match(regexp.MustCompile("^(\\/[[:alnum:]]+[[:graph:]]*)$")),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("locations", Location.Type).Ref("images"),
	}
}

func (File) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Path").Unique(),
	}
}
