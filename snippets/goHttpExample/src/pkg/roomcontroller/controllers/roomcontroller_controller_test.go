package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"httpExample/pkg/roomcontroller/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RoomMockService struct {
	RoomList []models.Room
}

func NewRoomMockService() (*RoomMockService) {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN01", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN02", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN03", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, models.Room{Name: "RoomN04", Description: "Dummy"})

	return &rms
}

func (r *RoomMockService) GetAllRooms() ([]models.Room, error) {
	return r.RoomList, nil
}

func (r *RoomMockService) GetRoom(name string) (models.Room, error) {

	for _, room := range(r.RoomList) {
		if room.Name == name {
			return room, nil
		}
	}

	return models.Room{}, errors.New("no such room")
}

func (r* RoomMockService) QueryRooms(query string) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}



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
	if(string(expected) != rec.Body.String()) {
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
	if (expected != rec.Body.String()) {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}