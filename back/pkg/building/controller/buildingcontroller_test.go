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

