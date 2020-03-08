package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"studi-guide/ent"
	"testing"
)

type RoomMockService struct {
	RoomList []*ent.Room
}

func NewRoomMockService() *RoomMockService {
	var rms RoomMockService

	rms.RoomList = append(rms.RoomList, &ent.Room{Name: "RoomN01", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, &ent.Room{Name: "RoomN02", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, &ent.Room{Name: "RoomN03", Description: "Dummy"})
	rms.RoomList = append(rms.RoomList, &ent.Room{Name: "RoomN04", Description: "Dummy"})

	return &rms
}

func (r *RoomMockService) GetAllRooms() ([]*ent.Room, error) {
	return r.RoomList, nil
}

func (r *RoomMockService) GetRoom(name string) (*ent.Room, error) {

	for _, room := range r.RoomList {
		if room.Name == name {
			return room, nil
		}
	}

	return &ent.Room{}, errors.New("no such room")
}

func (r *RoomMockService) AddRoom(room ent.Room) error {
	r.RoomList = append(r.RoomList, &room)
	return nil
}

func (r *RoomMockService) AddRooms(rooms []ent.Room) error {
	for _, room := range rooms {
		_ = r.AddRoom(room)
	}
	return nil
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
