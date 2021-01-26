package entitymapper

import (
	"studi-guide/pkg/building/db/ent"
	entbuilding "studi-guide/pkg/building/db/ent/building"
	"studi-guide/pkg/utils"
)

func (r *EntityMapper) GetAllBuildings() ([]*ent.Building, error) {
	buildings, err := r.client.Building.Query().WithCampus().WithAddress().WithBody().All(r.context)
	if err != nil {
		return nil, err
	}
	return buildings, nil
}

func (r *EntityMapper) GetBuilding(name string) (*ent.Building, error) {
	b, err := r.client.Building.
		Query().
		WithBody().
		WithCampus().
		WithAddress().
		Where(entbuilding.NameEqualFold(name)).First(r.context)

	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *EntityMapper) FilterBuildings(name string) ([]*ent.Building, error) {
	buildings, err := r.client.Building.Query().
		WithBody().
		WithCampus().
		WithAddress().
		Where(entbuilding.NameEqualFold(name)).All(r.context)
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

func (r *EntityMapper) GetFloorsFromBuilding(building *ent.Building) ([]string, error) {
	floors, err := building.QueryMapitems().
		Select("floor").Strings(r.context)

	if err != nil {
		return nil, err
	}

	return utils.Distinct(floors), nil
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
