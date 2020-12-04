package schema

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Campus holds the schema definition for the Campus entity.
type RssFeed struct {
	ent.Schema
}

// Fields of the RssFeed.
func (RssFeed) Fields() []ent.Field {
	return []ent.Field{
		field.String("Url").
			NotEmpty().Validate(func(s string) error {
			if !govalidator.IsURL(s) {
				return errors.New("url is invalid")
			}

			return nil
		}),

		field.String("Name").Unique(),
	}
}

// Indexes of the RssFeed
func (RssFeed) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("Name").Unique(),
	}
}
