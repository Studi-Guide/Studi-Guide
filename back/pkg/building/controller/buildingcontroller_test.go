package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	mock2 "studi-guide/pkg/building/mock"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
	"studi-guide/pkg/navigation"
	"testing"
)

func TestBuildingController_GetAllBuildings(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)
	router.ServeHTTP(rec, req)

	buildings,_ := buildingprovider.GetAllBuildings()

	expected, _ := json.Marshal(buildings)
	expected = append(expected, '\n')
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)

	buildingprovider.BuildingList = nil
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)
	router.ServeHTTP(rec, req)

	buildings,_ := buildingprovider.GetBuilding("main")

	expected, _ := json.Marshal(buildings)
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetBuildings_Filter_Negative(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/random", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)
	router.ServeHTTP(rec, req)

	buildings, _ := buildingprovider.GetBuilding("random")

	expected, _ := json.Marshal(buildings)
	expected = append(expected, '\n')
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)

	buildingprovider.BuildingList = nil
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)
	router.ServeHTTP(rec, req)

	building,_ := buildingprovider.GetBuilding("main");
	floors,_ := buildingprovider.GetFloorsFromBuilding(building)

	expected, _ := json.Marshal(floors)
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetFloorsFromBuilding_Exception(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	buildingprovider.BuildingList = nil
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProvider, mapsProvider)
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

	excepectedLocations := []entityservice.Location{{
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	locationProviderMock.EXPECT().
		FilterLocations("", "", "1", "main", "").
		Return(excepectedLocations, nil)

	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProviderMock, mapsProvider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(excepectedLocations)
	expected = append(expected, '\n')
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProviderMock, mapsProvider)
	router.ServeHTTP(rec, req)

	building,_ := buildingprovider.GetBuilding("main");
	rooms,_ := buildingprovider.RoomProvider.FilterRooms("1", "", "", "",building.Name, "")


	expected, _ := json.Marshal(rooms)
	expected = append(expected, '\n')
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

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProviderMock, mapsProvider)
	buildingprovider.RoomProvider.RoomList = nil
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}

func TestBuildingController_GetMapsFromBuildingFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors/1/maps", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	mockmaps := []entityservice.MapItem{{
		Doors:     nil,
		Color:     "",
		Sections:  nil,
		Campus:    "",
		Building:  "main",
		PathNodes: nil,
		Floor:     "1",
		},
		{
			Doors:     nil,
			Color:     "",
			Sections:  nil,
			Campus:    "",
			Building:  "main",
			PathNodes: nil,
			Floor:     "1",
		},
		{
			Doors:     nil,
			Color:     "",
			Sections:  nil,
			Campus:    "",
			Building:  "foobar",
			PathNodes: nil,
			Floor:     "3",
		},
		{
			Doors:     nil,
			Color:     "",
			Sections:  nil,
			Campus:    "",
			Building:  "main",
			PathNodes: nil,
			Floor:     "2",
		},
	}

	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	mapsProvider.EXPECT().FilterMapItems("1", "main", "").Return(mockmaps, nil)
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProviderMock, mapsProvider)
	router.ServeHTTP(rec, req)
	expected, _ := json.Marshal(mockmaps)
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestBuildingController_GetMapsFromBuildingFloor_Exception(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/buildings/main/floors/1/maps", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	buildingprovider := mock2.NewMockBuildingProvider()
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, buildingprovider.RoomProvider, locationProviderMock, mapsProvider)
	mapsProvider.EXPECT().FilterMapItems("1","main","").Return(nil, errors.New("mock exception"))
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}