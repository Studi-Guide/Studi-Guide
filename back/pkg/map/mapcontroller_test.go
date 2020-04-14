package maps

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/roomcontroller/models"
	"testing"
)

func TestMapController_GetMapItems(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(GetExpectedJson(provider.RoomList))
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItems_RoomError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItems_ConnectorError(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_Filter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/?floor=1", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	rooms,_ := provider.FilterRooms("1", "", "", "")

	expected, _ := json.Marshal(GetExpectedJson(rooms))
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItemsFromFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/floor/1", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	rooms,_ := provider.FilterRooms("1", "", "", "")

	expected, _ := json.Marshal(GetExpectedJson(rooms))
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestMapController_GetMapItemsFromFloor_BadInteger(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/floor/test", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_EmptyRoomlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/floor/1", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestMapController_GetMapItemsFromFloor_EmptyConnectorlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/map/floor/1", nil)

	provider :=  models.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/map")
	NewMapController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

// Helper method
func GetExpectedJson(rooms []entityservice.Room) ([]entityservice.MapItem)	 {
	var mapItems []entityservice.MapItem
	for _, room := range rooms {
		mapItems = append(mapItems, room.MapItem)
	}

	return mapItems
}
