package mock

import (
	"errors"
	"strings"
	"studi-guide/pkg/building/model"
	"studi-guide/pkg/rooom/mock"
)

type MockBuildingProvider struct {
	BuildingList []model.Building
	RoomProvider *mock.RoomMockService
}

func NewMockBuildingProvider() *MockBuildingProvider {
	mock := MockBuildingProvider{
		BuildingList: []model.Building{{
			Id:   1,
			Name: "main",
		},
		{
			Id: 2,
			Name: "sub",
		},
		},
		RoomProvider: mock.NewRoomMockService(),
	}

	return &mock
}


func (r *MockBuildingProvider) GetAllBuildings() ([]model.Building, error) {
	if r.BuildingList == nil {
		return nil, errors.New("no buildings available")
	}

	return r.BuildingList, nil
}

func (r *MockBuildingProvider) GetBuilding(name string) (model.Building, error) {
	if r.BuildingList == nil {
		return model.Building{}, errors.New("no buildings available")
	}

	for _, bd := range r.BuildingList {
		if bd.Name == name {
			return bd, nil
		}
	}

	return model.Building{}, nil
}

func (r *MockBuildingProvider) FilterBuildings(name string) ([]model.Building, error) {
	var buildings []model.Building
	if r.BuildingList == nil {
		return buildings, errors.New("no buildings available")
	}

	for _, bd := range r.BuildingList {
		if strings.Contains(bd.Name, name) {
			buildings = append(buildings, bd)
		}
	}

	return buildings, nil
}

func (r *MockBuildingProvider) GetFloorsFromBuilding(building model.Building) ([]string, error) {
	var floors []string
	if r.BuildingList == nil {
		return nil, errors.New("no buildings available")
	}

	for _, room := range r.RoomProvider.RoomList {
		floor := room.MapItem.Floor
		if !contains(floors , floor) {
			floors = append(floors, floor)
		}
	}

	return floors, errors.New("no room found")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}