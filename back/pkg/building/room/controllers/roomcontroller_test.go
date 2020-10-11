package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/room/mock"
	"testing"
)

func TestRoomlistIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/rooms")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(provider.RoomList)
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomlistIndex_Negativ(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/rooms")
	NewRoomController(roomRouter, provider)

	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestGetRoom(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms/RoomN01", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/rooms")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(provider.RoomList[0])
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestGetRoomNotExists(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms/abcdefg", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/rooms")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected := "\"no such room\""
	if expected != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomController_GetRoomsFromFloor_Filter(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms?floor=1", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	rooms, _ := provider.FilterRooms("1", "", "", "", "", "")

	expected, _ := json.Marshal(rooms)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomController_GetRoomList_FilterFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms?floor=0", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestRoomController_GetRoomList_BadFilterFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms?floor=first", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestRoomController_GetAllRoom(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms", nil)

	provider := mock.NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/rooms")
	NewRoomController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	rooms := provider.RoomList

	expected, _ := json.Marshal(rooms)
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}
