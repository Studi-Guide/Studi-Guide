package entityservice

import "studi-guide/pkg/navigation"

type Location struct {
	Id          int
	Name        string
	Description string
	Tags        []string
	Floor       string
	Building 	string
	PathNode    navigation.PathNode
}
