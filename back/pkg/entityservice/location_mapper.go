package entityservice

import (
	"studi-guide/ent"
	"studi-guide/pkg/navigation"
)

func (r *EntityService) locationArrayMapper(entLocations []*ent.Location) []Location {
	var locations []Location
	for _, entLocation := range entLocations {
		locations = append(locations, *r.locationMapper(entLocation))
	}
	return locations
}

func (r *EntityService) locationMapper(entLocation *ent.Location) *Location {
	l := Location{
		Id:          entLocation.ID,
		Name:        entLocation.Name,
		Description: entLocation.Description,
		Tags:        nil,
		Floor:       entLocation.Floor,
		PathNode:    navigation.PathNode{},
	}

	t, err := entLocation.Edges.TagsOrErr()
	if err == nil {
		l.Tags = r.tagsArrayMapper(t)
	}

	pn, err := entLocation.Edges.PathnodeOrErr()
	if err == nil {
		l.PathNode = *r.pathNodeMapper(pn, []*navigation.PathNode{}, false)
	}

	return &l
}

func (r *EntityService) mapSectionArray(sections []Section) ([]*ent.Section, error) {

	var entSections []*ent.Section

	for _, s := range sections {
		entS, err := r.mapSection(&s)
		if err != nil {
			return nil, err
		}
		entSections = append(entSections, entS)
	}

	return entSections, nil

}

