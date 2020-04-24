package entityservice

import (
	"studi-guide/ent"
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
