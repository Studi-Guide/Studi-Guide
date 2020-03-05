package models

import "studi-guide/pkg/navigation"

type SectionProvider interface {
	Draw() string
}

type LineSection struct {
	Id    int                   `json:"id" xml:"id" db:"ID"`
	Start navigation.Coordinate `json:"start" xml:"start" db:"start"`
	End   navigation.Coordinate `json:"end" xml:"end" db:"end"`
}

func (s *LineSection) Draw() string {
	return ""
}

type ArcSection struct {
	Id     int                   `json:"id" xml:"id" db:"ID"`
	Center navigation.Coordinate `json:"center" xml:"center" db:"center"`
	Start  navigation.Coordinate `json:"start" xml:"start" db:"start"`
	End    navigation.Coordinate `json:"end" xml:"end" db:"end"`
}

func (s *ArcSection) Draw() string {
	return ""
}
