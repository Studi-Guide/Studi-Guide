package models

import (
	"context"
	"database/sql"
	"github.com/ahmetb/go-linq/v3"
	fbsql "github.com/facebookincubator/ent/dialect/sql"
	"log"
	"os"
	"reflect"
	"strconv"
	"studi-guide/ent"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
	"testing"
)

var testRooms []Room
var testConnectors []ConnectorSpace

func setupTestRoomDbService() (RoomServiceProvider, *sql.DB) {
	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e := env.NewEnv()

	//testRooms = append(testRooms, Room{Id: 1, Name: "01", Description: "d"})
	//testRooms = append(testRooms, Room{Id: 2, Name: "02", Description: "d"})
	//testRooms = append(testRooms, Room{Id: 3, Name: "03", Description: "d"})

	drv, err := fbsql.Open(e.DbDriverName(), "file:"+e.DbDataSource()+"?_fk=1")
	if err != nil {
		return nil, nil
	}

	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	//defer client.Close()
	// run the auto migration tool.
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	testRooms = []Room{}
	testConnectors = []ConnectorSpace{}
	for i := 1; i < 4; i++ {

		sequence, err := client.Section.Create().Save(ctx)
		if err != nil {
			log.Println("error creating sequence:", err)
		}

		door, err := client.Door.Create().SetSectionID(sequence.ID).Save(ctx)
		if err != nil {
			log.Println("error creating door: ", err)
		}

		pathNode, err := client.PathNode.
			Create().
			SetXCoordinate(i).
			SetYCoordinate(i).
			SetZCoordinate(i).Save(ctx)

		if err != nil {
			log.Println("error creating pathnode:", err)
		}

		entRoom, err := client.Room.Create().
			SetName(strconv.Itoa(i)).
			SetPathNode(pathNode).
			AddDoorIDs(door.ID).
			SetFloor(i).
			Save(ctx)
		if err != nil {
			log.Println("error creating room:", err)
		}

		testRooms = append(testRooms, Room{
			Id:          entRoom.ID,
			MapItem:MapItem{
				Name:        entRoom.Name,
				Description: entRoom.Description,
				Alias:       nil,
				Doors: []Door{{
					Id:       door.ID,
					Section:  Section{Id: sequence.ID},
					PathNode: navigation.PathNode{},
				}},
				Color:    "",
				Sections: nil,
				Floor:    i,
			},

			PathNode: navigation.PathNode{Id: pathNode.ID, Coordinate:navigation.Coordinate{
				X: pathNode.XCoordinate,
				Y: pathNode.YCoordinate,
				Z: pathNode.ZCoordinate,
			}},
		})

		pathNode2, err := client.PathNode.Create().
			SetXCoordinate(i).
			SetYCoordinate(i).
			SetZCoordinate(i).Save(ctx)

		if err != nil {
			log.Println("error creating pathnode:", err)
		}

		sequence2, err := client.Section.Create().Save(ctx)
		if err != nil {
			log.Println("error creating sequence:", err)
		}

		entConnector, err := client.ConnectorSpace.Create().
			SetName(string(i)).
			AddConnectorPathNodeIDs(pathNode2.ID).
			AddConnectorSectionIDs(sequence2.ID).
			SetFloor(i).
			Save(ctx)
		if err != nil {
			log.Println("error creating connector:", err)
		}

		testConnectors = append(testConnectors, ConnectorSpace{
			Id:          entConnector.ID,
			MapItem:MapItem{
				Name:        entConnector.Name,
				Description: entConnector.Description,
				Alias:       nil,
				Color:    "",
				Sections: []Section {{Id:sequence2.ID}},
				Floor:    i,
			},

			PathNodes: []navigation.PathNode {{Id: pathNode2.ID, Coordinate:navigation.Coordinate{
				X: pathNode2.XCoordinate,
				Y: pathNode2.YCoordinate,
				Z: pathNode2.ZCoordinate,
			}}},
		})
	}

	dbService := RoomEntityService{client: client, table: "", context: ctx}

	return &dbService, drv.DB()
}

func TestNewRoomDbService(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "some_driver")
	os.Setenv("DB_DATA_SOURCE", ":some_source")

	e := env.NewEnv()

	dbService, err := NewRoomEntityService(e)
	if err == nil {
		t.Error("expected error; got: ", err)
	}
	if !reflect.ValueOf(dbService).IsNil() {
		t.Error("expected: ", nil, "; got: ", dbService)
	}

	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e = env.NewEnv()
	dbService, err = NewRoomEntityService(e)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
	if dbService == nil {
		t.Error("expected dbService; got: ", dbService)
	}

}

func TestGetRoomAllRooms(t *testing.T) {

	dbService, db := setupTestRoomDbService()

	getRooms, err := dbService.GetAllRooms()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []Room, b []Room) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	if !compare(testRooms, getRooms) {
		t.Error("expected: ", testRooms, "; got: ", getRooms)
	}

	db.Exec("drop table rooms")

	getRooms, err = dbService.GetAllRooms()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareRooms []Room
	if !compare(compareRooms, getRooms) {
		t.Error("expected: ", compareRooms, "; got: ", getRooms)
	}

}

func TestGetRoom(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	room, err := dbService.GetRoom(strconv.Itoa(2))
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testRooms[1], room) {
		t.Error("expected: ", testRooms[1], "; got: ", room)
	}

	room, err = dbService.GetRoom("4")
	if err == nil {
		t.Error("expected: ", nil, "; got: ", err)
	}
	var noneRoom Room
	if !reflect.DeepEqual(room, noneRoom) {
		t.Error("expected: ", noneRoom, "; got: ", room)
	}
}

func TestAddRoom(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	testRoom := Room{
		Id: 4,
		MapItem:MapItem{
			Name:        "04",
			Description: "description",
		},
	}
	
	err := dbService.AddRoom(testRoom)
	if err == nil {
		t.Error("expected: error", "; got: ", err)
	}

	err = dbService.AddRoom(testRoom)
	if err == nil {
		t.Error("expected error; got nil")
	}
}

func TestAddRooms(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	var newRooms []Room
	newRooms = append(newRooms, Room{
		Id: 4,
		MapItem:MapItem{
			Name:        "04",
			Description: "d",
		},
		
	})
	
	newRooms = append(newRooms, Room{
		Id: 4, 
		MapItem:MapItem{
			Name:        "04",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id: 5,
		MapItem:MapItem{
			Name:        "05",
			Description: "d",
		},
	})

	err := dbService.AddRooms(newRooms)
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	newRooms = newRooms[:0]
	newRooms = append(newRooms, Room{
		Id: 6,
		MapItem:MapItem{
			Name:        "06",
			Description: "d",
		},
	})
	
	newRooms = append(newRooms, Room{
		Id: 7, 
		MapItem:MapItem{
			Name:        "07",
			Description: "d",
		},
	})
	
	newRooms = append(newRooms, Room{
		Id: 8, 
		MapItem:MapItem{
			Name:        "08",
			Description: "d",
		},
	})

	err = dbService.AddRooms(newRooms)
	if err == nil {
		t.Error("expected: error", "; got: ", err)
	}
}

func TestRoomEntityService_GetAllPathNodes(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getNodes, err := dbService.GetAllPathNodes()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	checkNodes := func(a []Room, b []navigation.PathNode) bool {
		for i, _ := range a {
			found := linq.From(b).
				AnyWith(
					func(p interface{}) bool {
						return p.(navigation.PathNode).Id == a[i].PathNode.Id
					},
				)

			if !found {
				return false
			}
		}
		return true
	}

	if !checkNodes(testRooms, getNodes) {
		t.Error("expected: ", testRooms, "; got: ", getNodes)
	}
}

func TestRoomEntityService_GetAllConnectorSpaces(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	getConnectors, err := dbService.GetAllConnectorSpaces()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []ConnectorSpace, b []ConnectorSpace) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	expected := testConnectors
	if !compare(expected, getConnectors) {
		t.Error("expected: ", expected, "; got: ", getConnectors)
	}

	db.Exec("drop table connector_spaces")

	getConnectors, err = dbService.GetAllConnectorSpaces()
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareConnectors []ConnectorSpace
	if !compare(compareConnectors, getConnectors) {
		t.Error("expected: ", compareConnectors, "; got: ", getConnectors)
	}
}

func TestRoomEntityService_GetConnectorsFromFloor(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	getConnectors, err := dbService.FilterConnectorSpaces("1", "", "", "", "", nil, nil)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []ConnectorSpace, b []ConnectorSpace) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	var expected []ConnectorSpace
	linq.From(testConnectors).Where(func(p interface{}) bool { return p.(ConnectorSpace).MapItem.Floor == 1}).ToSlice(&expected)

	if !compare(expected, getConnectors) {
		t.Error("expected: ", expected, "; got: ", getConnectors)
	}

	db.Exec("drop table connector_spaces")

	getConnectors, err = dbService.FilterConnectorSpaces("1", "", "", "", "", nil, nil)
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareConnectors []ConnectorSpace
	if !compare(compareConnectors, getConnectors) {
		t.Error("expected: ", compareConnectors, "; got: ", getConnectors)
	}
}

func TestRoomEntityService_GetConnectorsFromFloor_FilterCoordinate(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	getConnectors, err := dbService.FilterConnectorSpaces("", "", "", "", "", &navigation.Coordinate{
		X: 1,
		Y: 1,
		Z: 1,
	}, &navigation.Coordinate{
		X: 0,
		Y: 0,
		Z: 0,
	})

	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []ConnectorSpace, b []ConnectorSpace) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	var expected []ConnectorSpace
	linq.From(testConnectors).Where(func(p interface{}) bool {
		return p.(ConnectorSpace).PathNodes[0].Coordinate.X == 1 &&
		p.(ConnectorSpace).PathNodes[0].Coordinate.Y == 1 &&
			p.(ConnectorSpace).PathNodes[0].Coordinate.Z == 1 }).ToSlice(&expected)

	if !compare(expected, getConnectors) {
		t.Error("expected: ", expected, "; got: ", getConnectors)
	}

	db.Exec("drop table connector_spaces")

	getConnectors, err = dbService.FilterConnectorSpaces("1", "", "", "", "", nil, nil)
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareConnectors []ConnectorSpace
	if !compare(compareConnectors, getConnectors) {
		t.Error("expected: ", compareConnectors, "; got: ", getConnectors)
	}
}

func TestRoomEntityService_GetRoomsFromFloor(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	getConnectors, err := dbService.FilterRooms("1", "", "", "")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []Room, b []Room) bool {
		if len(a) != len(b) {
			return false
		}
		for i, _ := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	var expected []Room
	linq.From(testRooms).Where(func(p interface{}) bool { return p.(Room).MapItem.Floor == 1}).ToSlice(&expected)

	if !compare(expected, getConnectors) {
		t.Error("expected: ", expected, "; got: ", getConnectors)
	}

	db.Exec("drop table rooms")

	getConnectors, err = dbService.FilterRooms("1", "", "", "")
	if err == nil {
		t.Error("expected error; got: ", err)
	}

	var compareRooms []Room
	if !compare(compareRooms, getConnectors) {
		t.Error("expected: ", compareRooms, "; got: ", getConnectors)
	}
}

func TestRoomEntityService_FilterRooms(t *testing.T) {

	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("1", "", "", "") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("abcd", "", "", "")
	if err == nil {
		t.Error("expect error", err, " got nil")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_RoomFilterParam(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("", "", "", "1") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("", "", "", "abcd")
	if err != nil {
		t.Error("expect no error", err, " got not nil")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}