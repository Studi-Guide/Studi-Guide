package services

import (
	"studi-guide/pkg/navigation"
)

type PathNodeProvider interface {
	GetAllPathNodes() ([]navigation.PathNode, error)
	GetPathNode(name string) (navigation.PathNode, error)
}