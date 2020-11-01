package entitymapper

import "studi-guide/pkg/navigation"

type Building struct {
	Id     int
	Name   string
	Floors []string
	Campus string
	Body   []navigation.GpsCoordinate
	Color  string
}
