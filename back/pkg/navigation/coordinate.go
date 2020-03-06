package navigation

import (
	"math"
)

type Coordinate struct {
	X int64 `json:"x" xml:"x" db:"X"`
	Y int64 `json:"y" xml:"y" db:"Y"`
	Z int64 `json:"z" xml:"z" db:"Z"`
}

func (c Coordinate) DistanceTo(other Coordinate) (int64) {

	p := Coordinate{X: c.X - other.X, Y: c.Y - other.Y, Z: c.Z - other.Z}

	distance := math.Sqrt(math.Pow(float64(p.X), 2) + math.Pow(float64(p.Y), 2) + math.Pow(float64(p.Z), 2))

	return int64(math.Round(distance))
}