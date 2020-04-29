package entityservice

import (
	"studi-guide/ent"
	entbuilding "studi-guide/ent/building"
	"studi-guide/pkg/utils"
)

func (r *EntityService) GetAllBuildings() ([]Building, error) {
	buildings, err := r.client.Building.Query().All(r.context)
	if err != nil {
		return nil, err
	}
	return r.buildingArrayMapper(buildings)
}

func (r *EntityService) GetBuilding(name string) (Building, error) {
	b, err := r.client.Building.Query().Where(entbuilding.NameEQ(name)).First(r.context)
	if err != nil {
		return Building{}, err
	}
	bding, err := r.buildingMapper(b)
	if err != nil {
		return Building{}, err
	}
	return *bding, nil
}

func (r *EntityService) FilterBuildings(name string) ([]Building, error) {
	buildings, err := r.client.Building.Query().Where(entbuilding.NameContains(name)).All(r.context)
	if err != nil {
		return nil, err
	}
	return r.buildingArrayMapper(buildings)
}

func (r *EntityService) buildingArrayMapper(entBuildings []*ent.Building) ([]Building, error) {
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

func (r *EntityService) buildingMapper(entBuilding *ent.Building) (*Building, error) {
	floors, _ := r.getFloorsFromBuilding(entBuilding)
	return &Building{
		Id:     entBuilding.ID,
		Name:   entBuilding.Name,
		Floors: floors,
	}, nil
}

func (r *EntityService) getFloorsFromBuilding(building *ent.Building) ([]string, error) {
	floors, err := building.QueryMapitems().
		Select("Floor").Strings(r.context)

	if err != nil {
		return nil, err
	}

	return utils.Distinct(floors), nil
}

func (r *EntityService) mapBuildingArray(buildings []Building) ([]*ent.Building, error) {
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

func (r *EntityService) mapBuilding(buildingName string) (*ent.Building, error) {
	entBuilding, _ := r.client.Building.Query().Where(entbuilding.NameEQ(buildingName)).First(r.context)
	if entBuilding != nil {
		return entBuilding, nil
	}

	return r.client.Building.Create().
		SetName(buildingName).
		Save(r.context)
}
