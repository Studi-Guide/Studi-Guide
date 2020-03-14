package navigation

type PathNode struct {
	Id             int
	Coordinate     Coordinate
	Group          *PathNodeGroup
	ConnectedNodes []*PathNode 		`json:"-"` // Not needed in frontend => stackOverflow in JSON
}
