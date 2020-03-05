package navigation

type Coordinate struct {
	X int64 `json:"x" xml:"x" db:"X"`
	Y int64 `json:"y" xml:"y" db:"Y"`
	Z int64 `json:"z" xml:"z" db:"Z"`
}
