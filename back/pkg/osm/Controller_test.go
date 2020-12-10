package osm

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"os"
	env2 "studi-guide/pkg/env"
	"studi-guide/pkg/osm/latlng"
	"testing"
)

func TestController_GetBounds(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/bounds", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	bounds := latlng.LatLngBounds{
		SouthWest: latlng.LatLngLiteral{
			Lat: 49.4126,
			Lng: 11.0111,
		},
		NorthEast: latlng.LatLngLiteral{
			Lat: 49.5118,
			Lng: 11.2167,
		},
	}

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(bounds)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestController_GetRoute_1(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/route", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Error("expected status bad request as code")
	}
}

func TestController_GetRoute_2(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/route?start=49.1,11.0&end=xyz", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Error("expected status bad request as code")
	}
}

func TestController_GetRoute_3(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/route?start=49.45,11.1&end=49.50,11.2", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	navProvider.EXPECT().GetRoute(latlng.LatLngLiteral{
		Lat: 49.45,
		Lng: 11.1,
	}, latlng.LatLngLiteral{
		Lat: 49.50,
		Lng: 11.2,
	}, "en-US").Return([]Route{{}}, nil)
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal([]Route{{}})
	if rec.Body.String() != string(expected) {
		t.Error("expected 'no route lol'")
	}
}

func TestController_GetRoute_4(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/route?start=49.45,11.1&end=49.50,11.2", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	navProvider.EXPECT().GetRoute(latlng.LatLngLiteral{
		Lat: 49.45,
		Lng: 11.1,
	}, latlng.LatLngLiteral{
		Lat: 49.50,
		Lng: 11.2,
	}, "en-US").Return(nil, errors.New("new error"))
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Error("expected internal server error")
	}
}

func TestController_GetRoute_5(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/osm/route?start=43.45,11.1&end=49.50,11.2", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	navProvider := NewMockOpenStreetMapNavigationProvider(ctrl)
	os.Setenv("OPENSTREETMAP_BOUNDS", "49.4126,11.0111;49.5118,11.2167")
	env := env2.NewEnv()
	router := gin.Default()
	osmRouter := router.Group("/osm")
	_ = NewOpenStreetMapController(osmRouter, navProvider, env)

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Error("expected bad request code")
	}
}
