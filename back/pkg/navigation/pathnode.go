package navigation

type PathNode struct {
	Id             int            `json:"id" xml:"id" db:"ID"`
	Coordinate     Coordinate     `json:"coordinate" xml:"coordinate" db:"coordinate"`
	Group          *PathNodeGroup `json:"group" xml:"group" db:"group"`
	ConnectedNodes []PathNode     `json:"connectedNodes" xml:"connectedNodes" db:"connectedNodes"`
}
