package campus

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/db/ent"
	"testing"
)

func TestCampusController_GetCampusByName(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/campus/K", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	campusprovider := NewMockCampusProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/campus")
	NewCampusController(mapRouter, campusprovider)

	address := ent.Address{
		ID:      1,
		Street:  "Kesslerplatz",
		Number:  "1",
		PLZ:     1111,
		City:    "Nürnberg",
		Country: "Deutschland",
		Edges:   ent.AddressEdges{},
	}

	campus := ent.Campus{
		ID:        1,
		ShortName: "K",
		Name:      "Kesslerplatz",
		Longitude: 100.5,
		Latitude:  200.8,
		Edges: ent.CampusEdges{
			Address: &address,
		},
	}

	campusprovider.EXPECT().GetCampus("K").Return(campus, nil)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(campus)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestCampusController_GetCampus_All(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/campus", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	campusprovider := NewMockCampusProvider(ctrl)
	router := gin.Default()
	mapRouter := router.Group("/campus")
	NewCampusController(mapRouter, campusprovider)

	address := ent.Address{
		ID:      1,
		Street:  "Kesslerplatz",
		Number:  "1",
		PLZ:     1111,
		City:    "Nürnberg",
		Country: "Deutschland",
		Edges:   ent.AddressEdges{},
	}

	address2 := ent.Address{
		ID:      1,
		Street:  "Am Platzl",
		Number:  "1",
		PLZ:     80331,
		City:    "Munich",
		Country: "Germany",
		Edges:   ent.AddressEdges{},
	}

	campus := []ent.Campus{
		{
			ID:        1,
			ShortName: "K",
			Name:      "Kesslerplatz",
			Longitude: 100.5,
			Latitude:  200.8,
			Edges: ent.CampusEdges{
				Address: &address,
			},
		},
		{
			ID:        2,
			ShortName: "B",
			Name:      "Hofbräu",
			Longitude: 500.5,
			Latitude:  1100.8,
			Edges: ent.CampusEdges{
				Address: &address2,
			},
		},
	}

	campusprovider.EXPECT().GetAllCampus().Return(campus, nil)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(campus)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}
