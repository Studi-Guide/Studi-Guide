package navigation

type PathNodeGroup struct {
	Id    int        `json:"id" xml:"id" db:"ID"`
	Name  string     `json:"name" xml:"name" db:"name"`
	Nodes []PathNode `json:"nodes" xml:"nodes" db:"nodes"`
}
