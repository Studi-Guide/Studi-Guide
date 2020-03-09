package models

import "studi-guide/pkg/navigation"

type Door struct {
	Id          int
	Sequence    Sequence
	PathNode    navigation.PathNode
}
