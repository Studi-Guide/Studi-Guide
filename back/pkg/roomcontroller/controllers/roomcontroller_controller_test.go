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

func TestGetRoom(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/roomlist/RoomN01", nil)

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
	req, _ := http.NewRequest("GET", "/roomlist/abcdefg", nil)

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
