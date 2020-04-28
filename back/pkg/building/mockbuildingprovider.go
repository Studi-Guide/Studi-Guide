package building

import (
	"errors"
	"strings"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/room/mock"
)

type MockBuildingProvider struct {
	BuildingList []entityservice.Building
	RoomProvider *mock.RoomMockService
}

func NewMockBuildingProvider() *MockBuildingProvider {
	mock := MockBuildingProvider{
		BuildingList: []entityservice.Building{{
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


func (r *MockBuildingProvider) GetAllBuildings() ([]entityservice.Building, error) {
	if r.BuildingList == nil {
		return nil, errors.New("no buildings available")
	}

	return r.BuildingList, nil
}

func (r *MockBuildingProvider) GetBuilding(name string) (entityservice.Building, error) {
	if r.BuildingList == nil {
		return entityservice.Building{}, errors.New("no buildings available")
	}

	for _, bd := range r.BuildingList {
		if bd.Name == name {
			return bd, nil
		}
	}

	return entityservice.Building{}, nil
}

func (r *MockBuildingProvider) FilterBuildings(name string) ([]entityservice.Building, error) {
	var buildings []entityservice.Building
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

func (r *MockBuildingProvider) GetFloorsFromBuilding(building entityservice.Building) ([]string, error) {
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

	return floors, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}