package models

import "studi-guide/pkg/navigation"

type Door struct {
	Id          int                     `json:"id" xml:"id" db:"ID"`
	Coordinates []navigation.Coordinate `json:"coordinates" xml:"coordinates" db:"coordinates"`
}
