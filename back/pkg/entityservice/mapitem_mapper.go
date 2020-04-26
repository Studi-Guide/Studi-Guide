package entityservice

import (
	"strconv"
	"studi-guide/ent"
	"studi-guide/ent/building"
	"studi-guide/ent/mapitem"
	"studi-guide/pkg/navigation"
)

func (r *EntityService) mapItemArrayMapper(entMapItems []*ent.MapItem) []MapItem {
	var mapItems []MapItem
	for _, entMapItem := range entMapItems {
		mapItems = append(mapItems, *r.mapItemMapper(entMapItem))
	}
	return mapItems
}

func (r *EntityService) mapItemMapper(entMapItem *ent.MapItem) *MapItem {
	m := MapItem{
		Doors:     []Door{},
		Color:     "",
		Sections:  []Section{},
		Floor:     entMapItem.Floor,
		PathNodes: []*navigation.PathNode{},
	}

	d, err := entMapItem.Edges.DoorsOrErr()
	if err == nil {
		m.Doors = r.doorArrayMapper(d)
	}

	s, err := entMapItem.Edges.SectionsOrErr()
	if err == nil {
		m.Sections = r.sectionArrayMapper(s)
	}

	p, err := entMapItem.Edges.PathNodesOrErr()
	if err == nil {
		m.PathNodes = r.pathNodeArrayMapper(p, []*navigation.PathNode{})
	}

	c, err := entMapItem.Edges.ColorOrErr()
	if err == nil {
		m.Color = c.Color
	}

	b, err := entMapItem.Edges.BuildingOrErr()
	if err == nil {
		m.Building = b.Name
	}

	return &m
}

func (r *EntityService) GetAllMapItems() ([]MapItem, error) {
	entMapItems, err := r.client.MapItem.Query().
		WithPathNodes().
		WithColor().
		WithBuilding().
		WithDoors().
		WithSections().
		All(r.context)
	if err != nil {
		return nil, err
	}

	return r.mapItemArrayMapper(entMapItems), nil
}

func (r *EntityService) FilterMapItems(floor, buildingFilter, campus string) ([]MapItem, error) {

	mapQuery := r.client.MapItem.Query()

	if len(buildingFilter) > 0 {
		mapQuery.Where(mapitem.HasBuildingWith(building.Name(buildingFilter)))
	}

	if len(floor) > 0 {
		iFloor, err := strconv.Atoi(floor)
		if err != nil {
			return nil, err
		} else {
			mapQuery = mapQuery.Where(mapitem.Floor(iFloor))
		}
	}

	// TODO Missing items: campus
	entMapItems, err := mapQuery.
		WithBuilding().
		WithDoors().
		WithSections().
		WithPathNodes().
		All(r.context)
	if err != nil {
		return nil, err
	}
	return r.mapItemArrayMapper(entMapItems), nil
}

