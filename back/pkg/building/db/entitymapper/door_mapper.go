package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/door"
	"studi-guide/pkg/navigation"
)

func (r *EntityMapper) doorArrayMapper(doors []*ent.Door) []Door {
	var d []Door
	for _, door := range doors {
		d = append(d, *r.doorMapper(door))
	}
	return d
}

func (r *EntityMapper) doorMapper(entDoor *ent.Door) *Door {

	entDoor, err := r.client.Door.Query().WithPathNode().WithOwner().WithSection().Where(door.ID(entDoor.ID)).First(r.context)
	if err != nil {
		return &Door{}
	}

	d := Door{
		Id:       entDoor.ID,
		Section:  Section{},
		PathNode: navigation.PathNode{},
	}

	s, err := entDoor.Edges.SectionOrErr()
	if err == nil {
		d.Section = *r.sectionMapper(s)
	}

	p, err := entDoor.Edges.PathNodeOrErr()
	if err == nil {
		d.PathNode = *r.pathNodeMapper(p, []*navigation.PathNode{}, true)
	}

	return &d
}

func (r *EntityMapper) mapDoorArray(doors []Door) ([]*ent.Door, error) {
	var entDoors []*ent.Door

	for _, d := range doors {
		entD, err := r.mapDoor(&d)
		if err != nil {
			return nil, err
		}
		entDoors = append(entDoors, entD)
	}

	return entDoors, nil
}

func (r *EntityMapper) mapDoor(d *Door) (*ent.Door, error) {

	if d.Id != 0 {
		return r.client.Door.Query().Where(door.ID(d.Id)).First(r.context)
	}

	sec, err := r.mapSection(&d.Section)
	if err != nil {
		return nil, err
	}

	pNode, err := r.mapPathNode(&d.PathNode)
	if err != nil {
		return nil, err
	}

	return r.client.Door.Create().
		SetPathNodeID(pNode.ID).
		SetSectionID(sec.ID).
		Save(r.context)
}
