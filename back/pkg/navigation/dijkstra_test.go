package navigation

import (
	"fmt"
	"testing"
)

func TestNewDijkstraNavigation(t *testing.T) {

}

func TestDijkstraNavigation_GetRoute1(t *testing.T) {
	var pathNodes []PathNode

	// P0(1,2,3) -> P1(3,2,3) -> P2(3,4,3)
	//           -> P3(1,4,3) ->
	pathNodes = append(pathNodes, PathNode{Id: 0, Coordinate: Coordinate{X: 1, Y: 2, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 1, Coordinate: Coordinate{X: 3, Y: 2, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 2, Coordinate: Coordinate{X: 3, Y: 4, Z: 3,}, Group: nil, ConnectedNodes: nil })
	pathNodes = append(pathNodes, PathNode{Id: 3, Coordinate: Coordinate{X: 1, Y: 4, Z: 3,}, Group: nil, ConnectedNodes: nil })

	pathNodes[0].ConnectedNodes = []*PathNode{&pathNodes[1]}
	pathNodes[1].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[2]}
	pathNodes[2].ConnectedNodes = []*PathNode{&pathNodes[1]}

	pathNodes[3].ConnectedNodes = []*PathNode{&pathNodes[0], &pathNodes[2]}

	navigation := NewDijkstraNavigation(pathNodes)

	route := navigation.GetRoute(pathNodes[0], pathNodes[2])
	fmt.Println(route)

}

func TestDijkstraNavigation_GetRoute2(t *testing.T) {

}

func TestDijkstraNavigation_GetRoute3(t *testing.T) {

}