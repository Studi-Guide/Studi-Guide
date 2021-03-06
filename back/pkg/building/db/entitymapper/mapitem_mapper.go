package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/building"
	entcampus "studi-guide/pkg/building/db/ent/campus"
	"studi-guide/pkg/building/db/ent/mapitem"
	"studi-guide/pkg/building/db/ent/pathnode"
	"studi-guide/pkg/navigation"
)

func (r *EntityMapper) mapItemArrayMapper(entMapItems []*ent.MapItem) []MapItem {
	var mapItems []MapItem
	for _, entMapItem := range entMapItems {
		mapItems = append(mapItems, *r.mapItemMapper(entMapItem))
	}
	return mapItems
}

func (r *EntityMapper) mapItemMapper(entMapItem *ent.MapItem) *MapItem {
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

func (r *EntityMapper) GetAllMapItems() ([]MapItem, error) {
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

func (r *EntityMapper) GetMapItemByPathNodeID(pathNodeID int) (MapItem, error) {

	mapQuery := r.client.MapItem.Query()
	entMapItem, err := mapQuery.Where(mapitem.HasPathNodesWith(pathnode.ID(pathNodeID))).
		WithPathNodes().
		WithColor().
		WithBuilding().
		WithDoors().
		WithSections().
		First(r.context)
	if err != nil {
		return MapItem{}, err
	}

	return *r.mapItemMapper(entMapItem), nil
}

func (r *EntityMapper) FilterMapItems(floor, buildingFilter, campus string) ([]MapItem, error) {

	mapQuery := r.client.MapItem.Query()

	if len(buildingFilter) > 0 {
		mapQuery.Where(mapitem.HasBuildingWith(building.NameEqualFold(buildingFilter)))
	}

	if len(floor) > 0 {
		mapQuery = mapQuery.Where(mapitem.FloorEqualFold(floor))
	}

	if len(campus) > 0 {
		mapQuery = mapQuery.Where(
			mapitem.HasBuildingWith(
				building.HasCampusWith(
					entcampus.Or(
						entcampus.ShortNameEqualFold(campus),
						entcampus.NameEqualFold(campus)))))
	}

	entMapItems, err := mapQuery.
		WithBuilding().
		WithDoors().
		WithSections().
		WithColor().
		WithPathNodes().
		All(r.context)
	if err != nil {
		return nil, err
	}
	return r.mapItemArrayMapper(entMapItems), nil
}
