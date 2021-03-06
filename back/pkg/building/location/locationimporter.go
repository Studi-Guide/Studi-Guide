package location

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/file"
	"studi-guide/pkg/navigation"
)

//Importer general importer interface
type Importer interface {
	RunImport() error
}

//JsonImporter Json implementation of the Importer interface
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
	Images      []file.File
	Icon        string
}

//NewLocationImporter creates a new location importer
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

func (r *JsonImporter) createRealLocations() error {
	fileHandle, err := os.Open(r.file)
	if err != nil {
		return err
	}

	var items []importLocation
	err = json.NewDecoder(fileHandle).Decode(&items)
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
			PathNode: navigation.PathNode{
				Id: importLoc.PathNode.Id,
				Coordinate: navigation.Coordinate{
					X: importLoc.PathNode.X,
					Y: importLoc.PathNode.Y,
					Z: importLoc.PathNode.Z,
				},
				Group:          nil,
				ConnectedNodes: nil,
			},
			Images: importLoc.Images,
			Icon:   importLoc.Icon,
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

//RunImport runs the importer
func (r *JsonImporter) RunImport() error {

	if err := r.createRealLocations(); err != nil {
		return err
	}

	for _, l := range r.locations {
		err := r.dbService.AddLocation(l)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
