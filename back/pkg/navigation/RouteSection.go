package navigation

type RouteSection struct {
	Route 		[]PathNode
	Description string
	Distance 	int64
	Building 	string
	Floor 		string
}
