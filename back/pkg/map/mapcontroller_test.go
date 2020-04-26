package maps

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
	"testing"
)

func TestMapController_GetMapItems(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps", nil)

	expectedMapItems := []entityservice.MapItem{{
		Doors:     []entityservice.Door{entityservice.Door{
			Id:       1,
			Section:  entityservice.Section{},
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
	expected = append(expected, '\n')
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

	expectedMapItems := []entityservice.MapItem{{
		Doors:     []entityservice.Door{entityservice.Door{
			Id:       1,
			Section:  entityservice.Section{},
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
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItemsFromFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps/floor/1", nil)

	expectedMapItems := []entityservice.MapItem{{
		Doors:     []entityservice.Door{entityservice.Door{
			Id:       1,
			Section:  entityservice.Section{},
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
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItemsFromFloor_BadInteger(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps/floor/test", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().FilterMapItems("test", "", "").Return(nil, errors.New("error text"))

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_EmptyRoomlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps/floor/1", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().FilterMapItems("1", "", "").Return(nil, errors.New("error text"))

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_EmptyConnectorlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/maps/floor/1", nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	provider := NewMockMapServiceProvider(ctrl)
	provider.EXPECT().FilterMapItems("1", "", "").Return(nil, errors.New("error text"))

	router := gin.Default()
	mapRouter := router.Group("/maps")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

