package osm

import (
	"encoding/json"
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