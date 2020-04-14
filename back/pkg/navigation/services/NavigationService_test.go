package services

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/models"
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


	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return()

	roomprovider := models.NewRoomMockService()
	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	nodes, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err != nil {
		t.Error(err)
	}

	startroom, _ := roomprovider.GetRoom(startroomname)
	endroom, _ := roomprovider.GetRoom(endroomname)
	expected, _, _ := calculator.GetRoute(startroom.PathNode, endroom.PathNode)
	expectedAsString, _ := json.Marshal(expected)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_CalculateFromString_Negative(t *testing.T) {
	startroomname := "RoomN00"
	endroomname := "RoomN02"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return()

	roomprovider := models.NewRoomMockService()

	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	_, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err == nil {
		t.Error(err)
	}
}

func TestNavigationService_Calculate(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return()

	roomprovider := models.NewRoomMockService()
	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	startroom, _ := roomprovider.GetRoom(startroomname)
	endroom, _ := roomprovider.GetRoom(endroomname)

	nodes, err := navigationservice.Calculate(startroom, endroom)

	if err != nil {
		t.Error(err)
	}

	expected, _, _ := calculator.GetRoute(startroom.PathNode, endroom.PathNode)
	expectedAsString, _ := json.Marshal(expected)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_CalculateStromString_Negative2(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN0001"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return()


	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	_, err := navigationservice.CalculateFromString(startroomname, endroomname)
	if err == nil {
		t.Error(err)
	}
}

func TestNavigationService_CalculateFromCoordinate(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockLocationProvider(ctrl)
	mock.EXPECT().GetAllPathNodes().Return()

	roomprovider := models.NewRoomMockService()
	calculator, _ := NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, mock)

	startroom, _ := roomprovider.GetRoom(startroomname)
	endroom, _ := roomprovider.GetRoom(endroomname)

	nodes, err := navigationservice.CalculateFromCoordinate(startroom.PathNode.Coordinate, endroom.PathNode.Coordinate)

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
