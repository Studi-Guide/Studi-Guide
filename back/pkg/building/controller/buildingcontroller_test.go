package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	mock2 "studi-guide/pkg/building/mock"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
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

	buildings,_ := buildingprovider.GetBuilding("random")

	expected, _ := json.Marshal(buildings)
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}
/*
func TestRoomController_GetRoomsFromFloor_BadInteger(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms/building/main/floor/bla", nil)

	provider :=  mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

/*

func TestRoomController_GetRoomFromFloor_EmptyRoomlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms/building/main/floor/1", nil)

	provider :=  mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

 */