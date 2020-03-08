package navigation

import (
	"math"
	"studi-guide/ent"
)

func Distance(a, b ent.PathNode) (int64) {

	p := ent.PathNode{XCoordinate: a.XCoordinate - b.XCoordinate, YCoordinate: a.YCoordinate - b.YCoordinate, ZCoordinate: a.ZCoordinate - b.ZCoordinate}

	distance := math.Sqrt(math.Pow(float64(p.XCoordinate), 2) + math.Pow(float64(p.YCoordinate), 2) + math.Pow(float64(p.ZCoordinate), 2))

	return int64(math.Round(distance))
}