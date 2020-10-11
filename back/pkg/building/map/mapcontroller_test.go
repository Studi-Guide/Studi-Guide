package maps

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
	"testing"
)

func TestMapController_GetMapItems(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps", nil)

	expectedMapItems := []entitymapper.MapItem{{
		Doors: []entitymapper.Door{entitymapper.Door{
			Id:       1,
			Section:  entitymapper.Section{},
			PathNode: navigation.PathNode{},
		}},
		Color:     "",
		Floor:     "1",
		Sections:  nil,
		Campus:    "",
		Building:  "",
		PathNodes: nil,
	}}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().GetAllMapItems().Return(expectedMapItems, nil)

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(expectedMapItems)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItems_RoomError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().GetAllMapItems().Return(nil, errors.New("error text"))

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItems_ConnectorError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().GetAllMapItems().Return(nil, errors.New("error text"))
	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_Filter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps?floor=1", nil)

	expectedMapItems := []entitymapper.MapItem{{
		Doors: []entitymapper.Door{entitymapper.Door{
			Id:       1,
			Section:  entitymapper.Section{},
			PathNode: navigation.PathNode{},
		}},
		Color:     "",
		Floor:     "1",
		Sections:  nil,
		Campus:    "",
		Building:  "",
		PathNodes: nil,
	}}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().FilterMapItems("1", "", "").Return(expectedMapItems, nil)

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(expectedMapItems)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItems_PathNodeID(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps?pathnodeid=1", nil)

	expectedMapItem := entitymapper.MapItem{
		Doors: []entitymapper.Door{entitymapper.Door{
			Id:       1,
			Section:  entitymapper.Section{},
			PathNode: navigation.PathNode{},
		}},
		Color:     "",
		Floor:     "1",
		Sections:  nil,
		Campus:    "",
		Building:  "",
		PathNodes: nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().GetMapItemByPathNodeID(1).Return(expectedMapItem, nil)

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal([]entitymapper.MapItem{expectedMapItem})
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}
