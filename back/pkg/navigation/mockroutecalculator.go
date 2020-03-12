package navigation

type MockRouteCalculator struct {
}

func NewMockRoutecalCulator() (RouteCalculator, error) {
	return &MockRouteCalculator{}, nil
}

func (l *MockRouteCalculator) GetRoute(start, end PathNode) ([]PathNode, error) {
	// Create dummy route

	distance := start.Coordinate.DistanceTo(end.Coordinate)
	node2 := PathNode{
		Id: 0,
		Coordinate: Coordinate{
			X: start.Coordinate.X + (distance / 3),
			Y: start.Coordinate.Y + (distance / 3),
			Z: start.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	node3 := PathNode{
		Id: 0,
		Coordinate: Coordinate{
			X: start.Coordinate.X + (2 * distance / 3),
			Y: start.Coordinate.Y + (2 * distance / 3),
			Z: start.Coordinate.Z,
		},
		Group:          nil,
		ConnectedNodes: nil,
	}

	nodes := []PathNode{start, node2, node3, end}
	return nodes, nil
}
