package info

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/building/location"
	"studi-guide/pkg/building/room/mock"
	"studi-guide/pkg/navigation"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestBuildingController_GetAllBuildings(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider)

	building := []*ent.Building{{
		ID:   1,
		Name: "main",
	},
		{
			ID:   2,
			Name: "sub",
		},
	}

	buildingprovider.EXPECT().GetAllBuildings().Return(building, nil)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(building)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetBuildings_Error(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider)

	buildingprovider.EXPECT().GetAllBuildings().Return(nil, errors.New("bla"))

	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}

func TestBuildingController_GetBuildings_Filter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider)

	building := ent.Building{
		ID:   1,
		Name: "main",
	}
	buildingprovider.EXPECT().GetBuilding("main").Return(&building, nil)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(building)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetBuildings_Filter_Error(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/random", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider)

	buildingprovider.EXPECT().GetBuilding("random").Return(nil, errors.New("bla"))
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}

func TestBuildingController_GetLocationsFromBuildingFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors/1/locations", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	excepectedLocations := []entitymapper.Location{{
		Id:          1,
		Name:        "test",
		Description: "",
		Tags:        nil,
		Floor:       "1",
		PathNode:    navigation.PathNode{},
	},
		{
			Id:          2,
			Name:        "test",
			Description: "",
			Tags:        nil,
			Floor:       "1",
			PathNode:    navigation.PathNode{},
		},
	}

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	locationProviderMock.EXPECT().
		FilterLocations("", "", "1", "main", "").
		Return(excepectedLocations, nil)

	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(excepectedLocations)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetRoomsFromBuildingFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors/1/rooms", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	testbuilding := ent.Building{
		ID:   1,
		Name: "main",
	}

	router.ServeHTTP(rec, req)

	rooms, _ := roomProvider.FilterRooms("1", "", "", "", testbuilding.Name, "")

	expected, _ := json.Marshal(rooms)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetRoomsFromBuildingFloor_Exception(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors/1/rooms", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	roomProvider.RoomList = nil
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}

func TestBuildingController_GetFloorsFromBuilding(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	testbuilding := ent.Building{
		ID:   1,
		Name: "main",
	}

	floorValue := []string{"1", "2", "3"}
	buildingprovider.EXPECT().GetBuilding("main").Return(&testbuilding, nil)
	buildingprovider.EXPECT().GetFloorsFromBuilding(&testbuilding).Return(floorValue, nil)

	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	router.ServeHTTP(rec, req)
	expected, _ := json.Marshal(floorValue)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetFloorsFromBuilding_BuildingNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	buildingprovider.EXPECT().GetBuilding("main").Return(nil, errors.New("not found"))

	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}

func TestBuildingController_GetFloorsFromBuilding_FloorError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	testbuilding := ent.Building{
		ID:   1,
		Name: "main",
	}

	buildingprovider.EXPECT().GetBuilding("main").Return(&testbuilding, nil)

	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock)
	buildingprovider.EXPECT().GetFloorsFromBuilding(&testbuilding).Return(nil, errors.New("not found"))
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}
