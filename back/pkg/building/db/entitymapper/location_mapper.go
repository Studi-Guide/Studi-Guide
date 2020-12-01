package entitymapper

import (
	"fmt"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/building"
	entcampus "studi-guide/pkg/building/db/ent/campus"
	"studi-guide/pkg/building/db/ent/location"
	"studi-guide/pkg/building/db/ent/tag"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/utils"
)

func (r *EntityMapper) locationArrayMapper(entLocations []*ent.Location) []Location {
	var locations []Location
	for _, entLocation := range entLocations {
		locations = append(locations, *r.locationMapper(entLocation))
	}
	return locations
}

func (r *EntityMapper) locationMapper(entLocation *ent.Location) *Location {
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

func (r *EntityMapper) mapSectionArray(sections []Section) ([]*ent.Section, error) {

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

func (r *EntityMapper) GetAllLocations() ([]Location, error) {
	entLocations, err := r.getLocationQuery().
		All(r.context)
	if err != nil {
		return nil, err
	}

	return r.locationArrayMapper(entLocations), nil
}

func (r *EntityMapper) GetLocation(name, buildingStr, campus string) (Location, error) {

	q := r.getLocationQuery().Where(location.NameEqualFold(name))

	if len(buildingStr) > 0 {
		q = q.Where(location.HasBuildingWith(building.NameEqualFold(buildingStr)))

		if len(campus) > 0 {
			q = q.Where(location.HasBuildingWith(building.HasCampusWith(entcampus.NameEqualFold(campus))))
		}
	}

	entLocation, err := q.First(r.context)
	if err != nil {
		return Location{}, &utils.QueryError{
			Query: fmt.Sprintf("Location %v: not found!", name),
			Err:   err,
		}
	}
	return *r.locationMapper(entLocation), nil
}

func (r *EntityMapper) FilterLocations(searchStr, tagStr, floor, buildingStr, campusStr string) ([]Location, error) {
	query := r.getLocationQuery()

	if len(searchStr) > 0 {
		query = query.Where(location.Or(
			location.NameContainsFold(searchStr),
			location.DescriptionContainsFold(searchStr),
			location.HasTagsWith(tag.NameContainsFold(searchStr)),
		))
	}

	if len(tagStr) > 0 {
		query = query.Where(location.HasTagsWith(tag.NameEqualFold(tagStr)))
	}

	if len(floor) > 0 {
		query = query.Where(location.FloorContains(floor))
	}

	if len(buildingStr) > 0 {
		query = query.Where(location.HasBuildingWith(building.NameEqualFold(buildingStr)))
	}

	if len(campusStr) > 0 {
		// Todo query campus
	}

	return r.queryLocations(query)
}

func (r *EntityMapper) GetRoutePoint(name string) (navigation.RoutePoint, error) {
	loc, err := r.GetLocation(name, "", "")
	if err != nil {
		return navigation.RoutePoint{}, err
	}
	return navigation.RoutePoint{
		Node:  loc.PathNode,
		Floor: loc.Floor}, err
}

func (r *EntityMapper) queryLocations(query *ent.LocationQuery) ([]Location, error) {
	entLocations, err := query.All(r.context)
	if err != nil {
		return nil, err
	}
	return r.locationArrayMapper(entLocations), nil
}

func (r *EntityMapper) getLocationQuery() *ent.LocationQuery {
	return r.client.Location.Query().
		WithPathnode().WithBuilding().WithTags()
}
