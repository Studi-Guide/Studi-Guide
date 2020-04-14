package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/entityservice"
	"studi-guide/pkg/navigation"
	"testing"
)

type MockNavigationService struct {
	nodes     []navigation.PathNode
	startroom string
	endroom   string
}

func NewRoomMockService(startroom, endrooom string) MockNavigationService {
	var rms MockNavigationService
	rms.nodes = rms.getDummyValues()
	rms.startroom = startroom
	rms.endroom = endrooom
	return rms
}

func (m MockNavigationService) CalculateFromString(startRoomName string, endRoomName string) ([]navigation.PathNode, error) {
	log.Print("Calculating entered")
	if !(startRoomName == m.startroom) || !(endRoomName == m.endroom) {
		return nil, errors.New("wrong rooms")
	}

	return m.nodes, nil
}

func (m MockNavigationService) Calculate(startRoom entityservice.Location, endRoom entityservice.Location) ([]navigation.PathNode, error) {
	if !(startRoom.Name == m.startroom) || !(endRoom.Name == m.endroom) {
		return nil, errors.New("wrong rooms")
	}

	return m.nodes, nil
}

func (m MockNavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) ([]navigation.PathNode, error) {
	return m.nodes, nil
}

func (m MockNavigationService) getDummyValues() []navigation.PathNode {
	// return dummy value
	node1 := navigation.PathNode{
		Id: 0,
		Coordinate: navigation.Coordinate{
			X: 1,
			Y: 1,
			Z: 1,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	node2 := navigation.PathNode{
		Id: 1,
		Coordinate: navigation.Coordinate{
			X: 2,
			Y: 2,
			Z: 2,
		},
		Group:          nil,
		ConnectedNodes: []*navigation.PathNode{&node1},
	}

	node1.ConnectedNodes = []*navigation.PathNode{&node2}
	nodes := []navigation.PathNode{node2, node1}
	return nodes
}

func TestNavigationCalculatefromString_NoRooms(t *testing.T) {
	startroomname := "dummystart"
	endroomname := "dummyend"
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/navigation/dir", nil)
	q := req.URL.Query()
	q.Add("startroom", startroomname)
	q.Add("endroom", endroomname)

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, err := json.Marshal(provider.nodes)
	if err != nil {
		t.Error(err)
	}

	expected = append(expected, '\n')
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestNavigationCalculatefromString(t *testing.T) {
	startroomname := "dummystart"
	endroomname := "dummyend"
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/navigation/dir", nil)
	q := req.URL.Query()
	q.Add("startroom", startroomname)
	q.Add("endroom", endroomname)

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected, err := json.Marshal(provider.nodes)
	if err != nil {
		t.Error(err)
	}

	expected = append(expected, '\n')
	if string(expected) != rec.Body.String() {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestNavigationCalculatefromString_Negativ(t *testing.T) {
	startroomname := "dummystart"
	endroomname := "dummyend"
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/navigation/dir", nil)
	q := req.URL.Query()
	q.Add("startroom", startroomname)
	q.Add("endroom", "differentroom")

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected := http.StatusBadRequest
	if rec.Code != expected {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}
