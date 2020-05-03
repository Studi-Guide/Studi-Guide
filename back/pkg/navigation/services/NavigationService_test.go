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
	return nodes, int64(distance), nil
}

func TestNavigationService_CalculateFromString(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"


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
	mock.EXPECT().GetPathNode("RoomN01").Return(loc1.PathNode, nil)
	mock.EXPECT().GetPathNode("RoomN02").Return(loc2.PathNode, nil)

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err != nil {
		t.Error(err)
	}


	expected, distance, _ := calculator.GetRoute(loc1.PathNode, loc2.PathNode)
	expectedRoute := navigation.NavigationRoute{
		Route: expected,
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
	mock.EXPECT().GetPathNode("RoomN00").Return(navigation.PathNode{}, errors.New("error text"))

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
	mock.EXPECT().GetPathNode("loc1").Return(loc1.PathNode, nil)
	mock.EXPECT().GetPathNode("loc2").Return(loc2.PathNode, nil)

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromString("loc1", "loc2")

	if err != nil {
		t.Error(err)
	}

	expected, distance, _ := calculator.GetRoute(loc1.PathNode, loc2.PathNode)
	expectedRoute := navigation.NavigationRoute{
		Route: expected,
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
	mock.EXPECT().GetPathNode("RoomN01").Return(loc1.PathNode, nil)
	mock.EXPECT().GetPathNode("RoomN0001").Return(navigation.PathNode{}, errors.New("error text"))


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
