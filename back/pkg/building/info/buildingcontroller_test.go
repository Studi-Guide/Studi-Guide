package info

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/building/location"
	maps "studi-guide/pkg/building/map"
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
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider, mapsProvider)

	building := []entitymapper.Building{{
		Id:     1,
		Name:   "main",
		Floors: []string{"1", "3"},
	},
		{
			Id:     2,
			Name:   "sub",
			Floors: []string{"1", "3"},
		},
	}

	buildingprovider.EXPECT().GetAllBuildings().Return(building, nil)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(building)
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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider, mapsProvider)

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
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider, mapsProvider)

	building := entitymapper.Building{
		Id:     1,
		Name:   "main",
		Floors: []string{"1", "3"},
	}
	buildingprovider.EXPECT().GetBuilding("main").Return(building, nil)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(building)
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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProvider := location.NewMockLocationProvider(ctrl)
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProvider, mapsProvider)

	buildingprovider.EXPECT().GetBuilding("random").Return(entitymapper.Building{}, errors.New("bla"))
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
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock, mapsProvider)
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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock, mapsProvider)
	testbuilding := entitymapper.Building{
		Id:     1,
		Name:   "main",
		Floors: []string{"1", "3"},
	}

	router.ServeHTTP(rec, req)

	rooms, _ := roomProvider.FilterRooms("1", "", "", "", testbuilding.Name, "")

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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock, mapsProvider)
	roomProvider.RoomList = nil
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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)

	mockmaps := []entitymapper.MapItem{{
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
	roomProvider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	mapsProvider.EXPECT().FilterMapItems("1", "main", "").Return(mockmaps, nil)
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock, mapsProvider)
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

	buildingprovider := NewMockBuildingProvider(ctrl)
	locationProviderMock := location.NewMockLocationProvider(ctrl)
	roomProvider := mock.NewRoomMockService()
	mapsProvider := maps.NewMockMapServiceProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/buildings")
	NewBuildingController(mapRouter, buildingprovider, roomProvider, locationProviderMock, mapsProvider)
	mapsProvider.EXPECT().FilterMapItems("1", "main", "").Return(nil, errors.New("mock exception"))
	router.ServeHTTP(rec, req)
	if http.StatusBadRequest != rec.Code {
		t.Error("expected ", http.StatusOK)
	}
}
