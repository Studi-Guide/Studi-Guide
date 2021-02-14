package navigation

//PathNode contains all information to be part of the navigation network
type PathNode struct {
	Id             int
	Coordinate     Coordinate
	Group          *PathNodeGroup
	ConnectedNodes []*PathNode `json:"-"` // Not needed in frontend => stackOverflow in JSON
}
