package navigation

type PathNodeGroup struct {
	Id    int        `db:"Id"`
	Name  string     `db:"Name"`
	Nodes []PathNode `db:"Nodes"`
}
