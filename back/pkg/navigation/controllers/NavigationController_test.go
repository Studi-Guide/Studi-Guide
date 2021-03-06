package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"studi-guide/pkg/building/db/entitymapper"
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

func (m MockNavigationService) CalculateFromString(startRoomName string, endRoomName string) (*navigation.NavigationRoute, error) {
	log.Print("Calculating entered")
	if !(startRoomName == m.startroom) || !(endRoomName == m.endroom) {
		return nil, errors.New("wrong rooms")
	}

	return &navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       m.nodes,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},

		Distance: int64(m.nodes[1].Coordinate.DistanceTo(m.nodes[0].Coordinate)),
	}, nil
}

func (m MockNavigationService) Calculate(startRoom entitymapper.Location, endRoom entitymapper.Location) (*navigation.NavigationRoute, error) {
	if !(startRoom.Name == m.startroom) || !(endRoom.Name == m.endroom) {
		return nil, errors.New("wrong rooms")
	}

	return &navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       m.nodes,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},
		Distance: int64(m.nodes[1].Coordinate.DistanceTo(m.nodes[0].Coordinate)),
	}, nil
}

func (m MockNavigationService) CalculateFromCoordinate(startCoordinate navigation.Coordinate, endCoordinate navigation.Coordinate) (*navigation.NavigationRoute, error) {
	return &navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       m.nodes,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},
		Distance: int64(m.nodes[1].Coordinate.DistanceTo(m.nodes[0].Coordinate)),
	}, nil
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
	q.Add("start", startroomname)
	q.Add("end", endroomname)

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expectedRoute := navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       provider.nodes,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},
		Distance: 2,
	}

	expected, err := json.Marshal(expectedRoute)
	if err != nil {
		t.Error(err)
	}

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
	q.Add("start", startroomname)
	q.Add("end", endroomname)

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)
	rec.Body.String()
	router.ServeHTTP(rec, req)

	expectedRoute := navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       provider.nodes,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},
		Distance: 2,
	}

	expected, err := json.Marshal(expectedRoute)
	if err != nil {
		t.Error(err)
	}

	actual := rec.Body.String()
	if string(expected) != actual {
		t.Errorf("expected = %v; actual = %v", string(expected), rec.Body.String())
	}
}

func TestNavigationCalculatefromString_Negativ(t *testing.T) {
	startroomname := "dummystart"
	endroomname := "dummyend"
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/navigation/dir", nil)
	q := req.URL.Query()
	q.Add("start", startroomname)
	q.Add("end", "differentroom")

	req.URL.RawQuery = q.Encode()

	provider := NewRoomMockService(startroomname, endroomname)
	router := gin.Default()
	roomRouter := router.Group("/navigation")
	NewNavigationController(roomRouter, provider)

	router.ServeHTTP(rec, req)

	expected := http.StatusBadRequest
	if rec.Code != expected {
		t.Errorf("expected = %v; actual = %v", string(rune(expected)), rec.Body.String())
	}
}
