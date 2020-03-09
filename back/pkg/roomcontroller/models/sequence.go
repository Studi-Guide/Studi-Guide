package models

import "studi-guide/pkg/navigation"

type Sequence struct {
	Id    int
	Start navigation.Coordinate
	End   navigation.Coordinate
}