package location

import (
	"encoding/json"
	"errors"
	"github.com/prometheus/common/log"
	"os"
	"path/filepath"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/navigation"
)

type Importer interface {
	RunImport() error
}

type JsonImporter struct {
	dbService LocationProvider
	file      string
	locations []entitymapper.Location
}

type importPathNode struct {
	Id                 int
	X                  int
	Y                  int
	Z                  int
	ConnectedPathNodes []int
}

type importLocation struct {
	Id          int
	Name        string
	Description string
	Tags        []string
	Floor       string
	Building    string
	PathNode    importPathNode
}



func NewLocationImporter(file string, dbService LocationProvider) (Importer, error) {
	var i Importer = nil
	ext := filepath.Ext(file)
	if ext == ".json" {
		i = &JsonImporter{dbService: dbService, file: file}
	} else {
		return nil, errors.New("unknown extension")
	}

	return i, nil
}

func (r* JsonImporter) createRealLocations() error {
	file, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var items []importLocation
	err = json.NewDecoder(file).Decode(&items)
	if err != nil {
		return err
	}

	for _, importLoc := range items {
		loc := entitymapper.Location{
			Id:          importLoc.Id,
			Name:        importLoc.Name,
			Description: importLoc.Description,
			Tags:        importLoc.Tags,
			Floor:       importLoc.Floor,
			Building:    importLoc.Building,
			PathNode:    navigation.PathNode{
				Id:             importLoc.PathNode.Id,
				Coordinate:     navigation.Coordinate{
					X: importLoc.PathNode.X,
					Y: importLoc.PathNode.Y,
					Z: importLoc.PathNode.Z,
				},
				Group:          nil,
				ConnectedNodes: nil,
			},
		}

		for _, id := range importLoc.PathNode.ConnectedPathNodes {
			loc.PathNode.ConnectedNodes = append(loc.PathNode.ConnectedNodes, &navigation.PathNode{
				Id:             id,
				Coordinate:     navigation.Coordinate{},
				Group:          nil,
				ConnectedNodes: nil,
			})
		}

		r.locations = append(r.locations, loc)
	}

	return nil
}

func (r *JsonImporter) RunImport() error {

	if err := r.createRealLocations(); err != nil {
		return err
	}

	for _, l := range r.locations {
		err := r.dbService.AddLocation(l)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
