package entitymapper

import (
	"errors"
	"regexp"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/color"
	"studi-guide/pkg/building/db/ent/tag"
)

func (r *EntityMapper) mapTag(t string, entLocation *ent.Location) (*ent.Tag, error) {
	entTag, err := r.client.Tag.Query().Where(tag.NameEqualFold(t)).First(r.context)
	if err != nil && entTag == nil {
		entTag, err = r.client.Tag.Create().SetName(t).AddLocations(entLocation).Save(r.context)
		if err != nil {
			return nil, err
		}
	} else {
		entTag, err = entTag.Update().AddLocations(entLocation).Save(r.context)
	}
	return entTag, err
}

func (r *EntityMapper) mapTagArray(ts []string, entLocation *ent.Location) ([]*ent.Tag, error) {
	var entTags []*ent.Tag
	for _, t := range ts {
		entTag, err := r.mapTag(t, entLocation)
		if err != nil {
			return nil, err
		}
		entTags = append(entTags, entTag)
	}
	return entTags, nil
}

func (r *EntityMapper) tagMapper(entTag *ent.Tag) string {
	return entTag.Name
}

func (r *EntityMapper) tagsArrayMapper(entTags []*ent.Tag) []string {
	var tags []string
	for _, t := range entTags {
		tags = append(tags, r.tagMapper(t))
	}
	return tags
}

func (r *EntityMapper) mapColor(c string) (*ent.Color, error) {

	format := "#[0-9a-fA-F]{3}$|#[0-9a-fA-F]{6}$"
	reg := regexp.MustCompile(format)

	if !reg.MatchString(c) {
		return nil, errors.New("color " + c + " does not match the required format: " + format)
	}

	col, err := r.client.Color.Query().Where(color.Color(c)).First(r.context)

	if err != nil && col == nil {
		col, err = r.client.Color.Create().SetName(c).SetColor(c).Save(r.context)
		if err != nil {
			return nil, err
		}
	}

	return col, nil
}
