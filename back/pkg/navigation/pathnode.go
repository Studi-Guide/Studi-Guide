package navigation

type PathNode struct {
	Id             int            `db:"Id"`
	Coordinate     Coordinate     `db:"Coordinate"`
	Group          *PathNodeGroup `db:"Group"`
	ConnectedNodes []PathNode     `db:"ConnectedNodes"`
}
