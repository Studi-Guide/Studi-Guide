package entityservice

import (
	"studi-guide/ent"
	"studi-guide/ent/location"
	"studi-guide/ent/tag"
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

	b, err := entLocation.Edges.BuildingOrErr()
	if err == nil {
		l.Building = b.Name
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

func (r *EntityService) GetAllLocations() ([]Location, error) {
	entLoactions, err := r.client.Location.Query().
		WithTags().
		WithBuilding().
		WithPathnode().
		All(r.context)
	if err != nil {
		return nil, err
	}

	return r.locationArrayMapper(entLoactions), nil
}

func (r *EntityService) GetLocation(name, building, campus string) (Location, error) {

	q := r.client.Location.Query().WithPathnode().WithBuilding().WithTags().Where(location.NameEQ(name))

	if len(building) > 0 {
		// TODO implement building
	}

	if len(campus) > 0 {
		// TODO implement campus
	}

	entLocation, err := q.First(r.context)
	if err != nil {
		return Location{}, err
	}
	return *r.locationMapper(entLocation), nil
}

func (r *EntityService) FilterLocations(name, tagStr, floor, building, campus string) ([]Location, error) {

	query := r.client.Location.Query().
		WithPathnode().WithTags()

	if len(name) > 0 {
		query = query.Where(location.NameContains(name))
	}

	if len(tagStr) > 0 {
		query = query.Where(location.HasTagsWith(tag.NameContains(tagStr)))
	}

	if len(floor) > 0 {
		query = query.Where(location.FloorContains(floor))
	}

	if len(building) > 0 {
		// Todo query building
	}

	if len(campus) > 0 {
		// Todo query campus
	}

	entLocations, err := query.All(r.context)
	if err != nil {
		return nil, err
	}
	return r.locationArrayMapper(entLocations), nil
}

