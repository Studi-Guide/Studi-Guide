package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	entbuilding "studi-guide/pkg/building/db/ent/building"
	"studi-guide/pkg/utils"
)

func (r *EntityMapper) GetAllBuildings() ([]Building, error) {
	buildings, err := r.client.Building.Query().All(r.context)
	if err != nil {
		return nil, err
	}
	return r.buildingArrayMapper(buildings)
}

func (r *EntityMapper) GetBuilding(name string) (Building, error) {
	b, err := r.client.Building.Query().Where(entbuilding.NameEqualFold(name)).First(r.context)
	if err != nil {
		return Building{}, err
	}
	bding, err := r.buildingMapper(b)
	if err != nil {
		return Building{}, err
	}
	return *bding, nil
}

func (r *EntityMapper) FilterBuildings(name string) ([]Building, error) {
	buildings, err := r.client.Building.Query().Where(entbuilding.NameEqualFold(name)).All(r.context)
	if err != nil {
		return nil, err
	}
	return r.buildingArrayMapper(buildings)
}

func (r *EntityMapper) buildingArrayMapper(entBuildings []*ent.Building) ([]Building, error) {
	var buildings []Building
	for _, b := range entBuildings {
		bding, err := r.buildingMapper(b)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, *bding)
	}
	return buildings, nil
}

func (r *EntityMapper) buildingMapper(entBuilding *ent.Building) (*Building, error) {
	floors, _ := r.getFloorsFromBuilding(entBuilding)
	building := Building{
		Id:     entBuilding.ID,
		Name:   entBuilding.Name,
		Floors: floors,
	}

	campus, _ := entBuilding.Edges.CampusOrErr()
	if campus != nil {
		building.Campus = campus.ShortName
	}

	return &building, nil
}

func (r *EntityMapper) getFloorsFromBuilding(building *ent.Building) ([]string, error) {
	floors, err := building.QueryMapitems().
		Select("floor").Strings(r.context)

	if err != nil {
		return nil, err
	}

	return utils.Distinct(floors), nil
}

func (r *EntityMapper) mapBuildingArray(buildings []Building) ([]*ent.Building, error) {
	var entBuildings []*ent.Building
	for _, b := range buildings {
		entBuilding, err := r.mapBuilding(b.Name)
		if err != nil {
			return nil, err
		}
		entBuildings = append(entBuildings, entBuilding)
	}
	return entBuildings, nil
}

func (r *EntityMapper) mapBuilding(buildingName string) (*ent.Building, error) {
	entBuilding, _ := r.client.Building.Query().Where(entbuilding.NameEQ(buildingName)).First(r.context)
	if entBuilding != nil {
		return entBuilding, nil
	}

	return r.client.Building.Create().
		SetName(buildingName).
		Save(r.context)
}
