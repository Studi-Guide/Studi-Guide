package entityservice

import "studi-guide/pkg/navigation"

type Section struct {
	Id    int
	Start navigation.Coordinate
	End   navigation.Coordinate
}
