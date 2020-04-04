package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoomlistIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/", nil)

	provider := NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/roomlist")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(provider.RoomList)
	expected = append(expected, '\n')
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomlistIndex_Negativ(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/", nil)

	provider := NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/roomlist")
	NewRoomController(roomRouter, provider)

	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestGetRoom(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/room/RoomN01", nil)

	provider := NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/roomlist")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, _ := json.Marshal(provider.RoomList[0])
	expected = append(expected, '\n')
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestGetRoomNotExists(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/room/abcdefg", nil)

	provider := NewRoomMockService()
	router := gin.Default()
	roomRouter := router.Group("/roomlist")
	NewRoomController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected := "\"no such room\"\n"
	if expected != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomController_GetRoomsFromFloor(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/floor/1", nil)

	provider :=  NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/roomlist")
	NewRoomController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	rooms,_ := provider.GetRoomsFromFloor(1)

	expected, _ := json.Marshal(rooms)
	expected = append(expected, '\n')
	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestRoomController_GetRoomsFromFloor_BadInteger(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/floor/bla", nil)

	provider :=  NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/roomlist")
	NewRoomController(mapRouter, provider)
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}

func TestRoomController_GetRoomFromFloor_EmptyRoomlist(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/floor/1", nil)

	provider :=  NewRoomMockService()
	router := gin.Default()
	mapRouter := router.Group("/roomlist")
	NewRoomController(mapRouter, provider)
	provider.RoomList = nil
	router.ServeHTTP(rec, req)

	if http.StatusBadRequest != rec.Code {
		t.Errorf("expected = %v; actual = %v", http.StatusBadRequest, rec.Code)
	}
}