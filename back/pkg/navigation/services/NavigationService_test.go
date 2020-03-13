package services

import (
	"encoding/json"
	"studi-guide/pkg/navigation"
	"studi-guide/pkg/roomcontroller/controllers"
	"testing"
)

func TestNavigationService_CalculateFromString(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"
	roomprovider := controllers.NewRoomMockService()
	calculator, _ := navigation.NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, roomprovider)

	nodes, err := navigationservice.CalculateFromString(startroomname, endroomname)

	if err != nil {
		t.Error(err)
	}

	startroom, _ := roomprovider.GetRoom(startroomname)
	endroom, _ := roomprovider.GetRoom(endroomname)
	expected, _ := calculator.GetRoute(startroom.PathNode, endroom.PathNode)
	expectedAsString, _ := json.Marshal(expected)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_Calculate(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"
	roomprovider := controllers.NewRoomMockService()
	calculator, _ := navigation.NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, roomprovider)

	startroom, _ := roomprovider.GetRoom(startroomname)
	endroom, _ := roomprovider.GetRoom(endroomname)

	nodes, err := navigationservice.Calculate(startroom, endroom)

	if err != nil {
		t.Error(err)
	}

	expected, _ := calculator.GetRoute(startroom.PathNode, endroom.PathNode)
	expectedAsString, _ := json.Marshal(expected)
	resultAsString, _ := json.Marshal(nodes)
	if string(expectedAsString) != string(resultAsString) {
		t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	}
}

func TestNavigationService_CalculateFromCoordinate(t *testing.T) {
	startroomname := "RoomN01"
	endroomname := "RoomN02"
	roomprovider := controllers.NewRoomMockService()
	calculator, _ := navigation.NewMockRoutecalCulator()
	navigationservice, _ := NewNavigationService(calculator, roomprovider)

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

	//expected, _ := calculator.GetRoute(startroom.PathNode, endroom.PathNode)
	//expectedAsString, _ := json.Marshal(expected)
	//resultAsString, _ := json.Marshal(nodes)
	//if string(expectedAsString) != string(resultAsString) {
	//	t.Errorf("expected = %v; actual = %v", string(expectedAsString), string(resultAsString))
	//}
}
