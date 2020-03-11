package models

import "studi-guide/pkg/navigation"

type Door struct {
	Id       int
	Section  Section
	PathNode navigation.PathNode
}
