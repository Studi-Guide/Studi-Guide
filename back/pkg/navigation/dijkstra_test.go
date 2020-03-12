package navigation

import (
	"reflect"
	"testing"
)

func TestNewDijkstraNavigation(t *testing.T) {

}

func TestDijkstraNavigation_GetRoute1(t *testing.T) {
	var pathNodes []PathNode

	// P0(1,2,3) -> P1(3,2,3) -> P2(3,4,3)
	//          `-> P3(1,4,3) -^
	pathNodes = append(pathNodes, PathNode{Id: 0, Coordinate: Coordinate{X: 1, Y: 2, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 1, Coordinate: Coordinate{X: 3, Y: 2, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 2, Coordinate: Coordinate{X: 3, Y: 4, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 3, Coordinate: Coordinate{X: 1, Y: 4, Z: 3,}, Group: nil, ConnectedNodes: nil })

	pathNodes[0].ConnectedNodes = []*PathNode{&pathNodes[1], &pathNodes[3]}
	pathNodes[1].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[2]}
	pathNodes[2].ConnectedNodes = []*PathNode{&pathNodes[1], &pathNodes[3]}
	pathNodes[3].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[2]}

	expected1 := []PathNode{pathNodes[0], pathNodes[1], pathNodes[2]}
	expected2 := []PathNode{pathNodes[0], pathNodes[3], pathNodes[2]}

	navigation := NewDijkstraNavigation(pathNodes)
	route, distance, err := navigation.GetRoute(pathNodes[0], pathNodes[2])

	if distance != 4  {
		t.Error("expected distance:", 0, "; got:", distance)
	}

	if err != nil {
		t.Error("expected nil, got:", err)
	}

	if !reflect.DeepEqual(expected1, route) && !reflect.DeepEqual(expected2, route) {
		t.Error("expected route", expected1, "or", expected2, "; got: ", route)
	}

}

func TestDijkstraNavigation_GetRoute2(t *testing.T) {

	var pathNodes []PathNode

	// P0(1,2,3) -> P1(8,2,3) -> P2(8,10,3) ----------------> P3(8,10,6)
	//          `-> P4(10,2,3)-> P5(10,2,6) -> P6(10,10,6) -^
	pathNodes = append(pathNodes, PathNode{Id: 0, Coordinate: Coordinate{X: 1, Y: 2, Z: 3}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 1, Coordinate: Coordinate{X: 8, Y: 2, Z: 3}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 2, Coordinate: Coordinate{X: 8, Y: 10, Z: 3}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 3, Coordinate: Coordinate{X: 8, Y: 10, Z: 6}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 4, Coordinate: Coordinate{X: 10, Y: 2, Z: 3}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 5, Coordinate: Coordinate{X: 10, Y: 2, Z: 6}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 6, Coordinate: Coordinate{X: 10, Y: 10, Z: 6}, Group: nil, ConnectedNodes: nil})

	pathNodes[0].ConnectedNodes = []*PathNode{&pathNodes[1], &pathNodes[4]}
	pathNodes[1].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[2]}
	pathNodes[2].ConnectedNodes = []*PathNode{&pathNodes[1], &pathNodes[3]}
	pathNodes[3].ConnectedNodes = []*PathNode{&pathNodes[2], &pathNodes[6]}
	pathNodes[4].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[5]}
	pathNodes[5].ConnectedNodes = []*PathNode{&pathNodes[4], &pathNodes[6]}
	pathNodes[6].ConnectedNodes = []*PathNode{&pathNodes[5], &pathNodes[3]}

	expected := []PathNode{pathNodes[0], pathNodes[1], pathNodes[2], pathNodes[3]}

	navigation := NewDijkstraNavigation(pathNodes)
	route, distance, err := navigation.GetRoute(pathNodes[0], pathNodes[3])

	if distance != 18  {
		t.Error("expected distance:", 0, "; got:", distance)
	}

	if err != nil {
		t.Error("expected nil, got:", err)
	}

	if len(route) != len(expected) {
		t.Error("expected route len: ", len(expected), "; got:", len(route))
	}

	for i, _ := range route {
		if expected[i].Id != route[i].Id {
			t.Error("expected Id:", expected[i].Id, "; got:", route[i].Id)
		}
	}

}

func TestDijkstraNavigation_GetRoute3(t *testing.T) {

	var pathNodes []PathNode

	// P1(3,3,3) -> P2(3,3,3)
	pathNodes = append(pathNodes, PathNode{Id: 0, Coordinate: Coordinate{X: 3, Y: 3, Z: 3}, Group: nil, ConnectedNodes: nil})
	pathNodes = append(pathNodes, PathNode{Id: 1, Coordinate: Coordinate{X: 3, Y: 3, Z: 3}, Group: nil, ConnectedNodes: nil})

	pathNodes[0].ConnectedNodes = []*PathNode{&pathNodes[1]}
	pathNodes[1].ConnectedNodes = []*PathNode{&pathNodes[0]}

	navigation := NewDijkstraNavigation(pathNodes)

	route, distance, err := navigation.GetRoute(pathNodes[0], pathNodes[1])

	if distance != 0  {
		t.Error("expected distance:", 0, "; got:", distance)
	}

	if err != nil && route != nil {
		t.Error("expected nil, got:", err, route)
	}

}