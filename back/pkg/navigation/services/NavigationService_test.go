package services

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
	"testing"
)

type MockRouteCalculator struct {
}

func NewMockRoutecalCulator() (navigation.RouteCalculator, error) {
	return &MockRouteCalculator{}, nil
}

func (l *MockRouteCalculator) Initialize(pathnodes []navigation.PathNode) {
}

func (l *MockRouteCalculator) GetRoute(start, end navigation.PathNode) ([]navigation.PathNode, int64, error) {
	// Create dummy route

	distance := start.Coordinate.DistanceTo(end.Coordinate)
	node2 := navigation.PathNode{
		Id: 0,
		Coordinate: navigation.Coordinate{
			X: start.Coordinate.X + (distance / 3),
			Y: start.Coordinate.Y + (distance / 3),
			Z: start.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	node3 := navigation.PathNode{
		Id: 0,
		Coordinate: navigation.Coordinate{
			X: start.Coordinate.X + (2 * distance / 3),
			Y: start.Coordinate.Y + (2 * distance / 3),
			Z: start.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	nodes := []navigation.PathNode{start, node2, node3, end}

	overalldistance := start.Coordinate.DistanceTo(node2.Coordinate) + node2.Coordinate.DistanceTo(node3.Coordinate) + node3.Coordinate.DistanceTo(end.Coordinate)
	return nodes, int64(overalldistance), nil
}

func TestNavigationService_CalculateFromString(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"

	loc1 := entitymapper.Location{
		Id:          1,
		Name:        "RoomN01",
		Description: "",
		Tags:        nil,
		Building:    "TestBuilding",
		Floor:       "2",
		PathNode: navigation.PathNode{
			Id: 1,
			Coordinate: navigation.Coordinate{
				X: 200,
				Y: 300,
				Z: 2,
			},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}
	loc2 := entitymapper.Location{
		Id:          2,
		Name:        "RoomN02",
		Building:    "TestBuilding",
		Description: "",
		Tags:        nil,
		Floor:       "2",
		PathNode: navigation.PathNode{
			Id: 2,
			Coordinate: navigation.Coordinate{
				X: 2,
				Y: 3,
				Z: 2,
			},
			Group:          nil,
			ConnectedNodes: []*navigation.PathNode{&loc1.PathNode},
		},
	}

	loc1.PathNode.ConnectedNodes = []*navigation.PathNode{&loc2.PathNode}

	expecteddistance := loc1.PathNode.Coordinate.DistanceTo(loc2.PathNode.Coordinate)
	node2 := navigation.PathNode{
		Id: 0,
		Coordinate: navigation.Coordinate{
			X: loc1.PathNode.Coordinate.X + (expecteddistance / 3),
			Y: loc1.PathNode.Coordinate.Y + (expecteddistance / 3),
			Z: loc1.PathNode.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	node3 := navigation.PathNode{
		Id: 0,
		Coordinate: navigation.Coordinate{
			X: loc1.PathNode.Coordinate.X + (2 * expecteddistance / 3),
			Y: loc1.PathNode.Coordinate.Y + (2 * expecteddistance / 3),
			Z: loc1.PathNode.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockPathNodeProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return([]navigation.PathNode{loc1.PathNode, loc2.PathNode}, nil)
	mock.EXPECT().GetRoutePoint("RoomN01").Return(navigation.RoutePoint{
		Node:  loc1.PathNode,
		Floor: loc1.Floor,
	}, nil)
	mock.EXPECT().GetRoutePoint("RoomN02").Return(navigation.RoutePoint{
		Node:  loc2.PathNode,
		Floor: loc2.Floor,
	}, nil)

	mock.EXPECT().GetPathNodeLocationData(loc1.PathNode).Times(1).Return(navigation.LocationData{
		Building: loc1.Building,
		Floor:    loc1.Floor,
		Campus:   "",
	}, nil)

	mock.EXPECT().GetPathNodeLocationData(loc2.PathNode).Times(1).Return(navigation.LocationData{
		Building: loc2.Building,
		Floor:    loc2.Floor,
		Campus:   "",
	}, nil)

	mock.EXPECT().GetPathNodeLocationData(node2).Return(navigation.LocationData{
		Building: loc2.Building,
		Floor:    loc2.Floor,
		Campus:   "",
	}, nil)

	mock.EXPECT().GetPathNodeLocationData(node3).Return(navigation.LocationData{
		Building: loc2.Building,
		Floor:    loc2.Floor,
		Campus:   "",
	}, nil)

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err != nil {
		t.Error(err)
	}

	expected, distance, _ := calculator.GetRoute(loc1.PathNode, loc2.PathNode)
	expectedRoute := navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       expected,
			Description: "",
			Distance:    distance,
			Building:    loc2.Building,
			Floor:       loc2.Floor,
		}},
		Start: navigation.RoutePoint{
			Node:  loc1.PathNode,
			Floor: loc1.Floor,
		},
		End: navigation.RoutePoint{
			Node:  loc2.PathNode,
			Floor: loc2.Floor,
		},
		Distance: distance,
	}

	expectedAsString, _ := json.Marshal(expectedRoute)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_CalculateFromString_Negative(t *testing.T) {
	startroomname := "RoomN00"
	endroomname := "RoomN02"

	loc1 := entitymapper.Location{
		Id:          1,
		Name:        "RoomN00",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}
	loc2 := entitymapper.Location{
		Id:          2,
		Name:        "RoomN02",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockPathNodeProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return([]navigation.PathNode{loc1.PathNode, loc2.PathNode}, nil)
	mock.EXPECT().GetRoutePoint("RoomN00").Return(navigation.RoutePoint{}, errors.New("error text"))

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	_, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err == nil {
		t.Error(err)
	}
}

func TestNavigationService_Calculate(t *testing.T) {
	loc1 := entitymapper.Location{
		Id:          1,
		Name:        "RoomN01",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}
	loc2 := entitymapper.Location{
		Id:          2,
		Name:        "RoomN02",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockPathNodeProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return([]navigation.PathNode{loc1.PathNode, loc2.PathNode}, nil)
	mock.EXPECT().GetRoutePoint("loc1").Return(navigation.RoutePoint{
		Node:  loc1.PathNode,
		Floor: loc1.Floor,
	}, nil)
	mock.EXPECT().GetRoutePoint("loc2").Return(navigation.RoutePoint{
		Node:  loc2.PathNode,
		Floor: loc2.Floor,
	}, nil)
	mock.EXPECT().GetPathNodeLocationData(loc1.PathNode).Times(2).Return(navigation.LocationData{
		Building: loc1.Building,
		Floor:    loc1.Floor,
		Campus:   "",
	}, nil)

	mock.EXPECT().GetPathNodeLocationData(loc2.PathNode).Times(2).Return(navigation.LocationData{
		Building: loc2.Building,
		Floor:    loc2.Floor,
		Campus:   "",
	}, nil)

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromString("loc1", "loc2")

	if err != nil {
		t.Error(err)
	}

	expected, distance, _ := calculator.GetRoute(loc1.PathNode, loc2.PathNode)
	expectedRoute := navigation.NavigationRoute{
		RouteSections: []navigation.RouteSection{{
			Route:       expected,
			Description: "",
			Distance:    0,
			Building:    "",
			Floor:       "",
		}},
		Distance: distance,
	}
	expectedAsString, _ := json.Marshal(expectedRoute)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_CalculateStromString_Negative2(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN0001"

	loc1 := entitymapper.Location{
		Id:          1,
		Name:        "RoomN01",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockPathNodeProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return([]navigation.PathNode{loc1.PathNode}, nil)
	mock.EXPECT().GetRoutePoint("RoomN01").Return(navigation.RoutePoint{
		Node:  loc1.PathNode,
		Floor: loc1.Floor,
	}, nil)

	mock.EXPECT().GetRoutePoint("RoomN0001").Return(navigation.RoutePoint{}, errors.New("error text"))

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	_, err := navigationservice.CalculateFromString(startroomname, endroomname)
	if err == nil {
		t.Error(err)
	}
}

func TestNavigationService_CalculateFromCoordinate(t *testing.T) {

	loc1 := entitymapper.Location{
		Id:          1,
		Name:        "RoomN01",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}
	loc2 := entitymapper.Location{
		Id:          2,
		Name:        "RoomN02",
		Description: "",
		Tags:        nil,
		PathNode:    navigation.PathNode{},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockPathNodeProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return([]navigation.PathNode{loc1.PathNode, loc2.PathNode}, nil)

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromCoordinate(loc1.PathNode.Coordinate, loc2.PathNode.Coordinate)

	if err != nil {
		t.Error(err)
	}

	if nodes != nil {
		t.Error(err)
	}

	//TODO implement unit test for feature

	//expected, _ := calculator.GetRoute(startroom.PathNodes, endroom.PathNodes)
	//expectedAsString, _ := json.Marshal(expected)
	//resultAsString, _ := json.Marshal(nodes)
	//if string(expectedAsString) != string(resultAsString) {
	//	t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	//}
}
