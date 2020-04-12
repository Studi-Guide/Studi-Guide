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
	for i := 1; i < 4; i++ {

		sequence, err := client.Section.Create().Save(ctx)
		if err != nil {
			log.Println("error creating sequence:", err)
		}

		pathNode, err := client.PathNode.
			Create().
			SetID(i).
			SetXCoordinate(i).
			SetYCoordinate(i).
			SetZCoordinate(i).Save(ctx)

		if err != nil {
			log.Println("error creating pathnode:", err)
		}

		door, err := client.Door.Create().SetSectionID(sequence.ID).SetPathNode(pathNode).Save(ctx)
		if err != nil {
			log.Println("error creating door: ", err)
		}

		entMapItem, err := client.MapItem.Create().
			AddPathNodes(pathNode).
			AddDoorIDs(door.ID).
			SetFloor(i).
			Save(ctx)

		if err != nil {
			log.Println("error creating map item:", err)
		}

		entLocation, err := client.Location.Create().
			SetName(strconv.Itoa(i)).
			SetPathnode(pathNode).
			Save(ctx)

		if err != nil {
			log.Println("error creating location:", err)
		}

		entRoom, err := client.Room.Create().
			SetLocation(entLocation).
			SetMapitem(entMapItem).
			Save(ctx)
		if err != nil {
			log.Println("error creating room:", err)
		}

		patnode := navigation.PathNode{
			Id:             pathNode.ID,
			Coordinate:navigation.Coordinate{
				X: pathNode.XCoordinate,
				Y: pathNode.YCoordinate,
				Z: pathNode.ZCoordinate,
			}}

		testRooms = append(testRooms, Room{
			Id:          entRoom.ID,
			MapItem: MapItem{
				Doors: []Door{{
					Id:       door.ID,
					Section:  Section{Id: sequence.ID},
					PathNode: patnode,
				}},
				Color:    "",
				Sections: nil,
				Floor:    i,
				PathNodes: []*navigation.PathNode{&patnode},
			},
			Location: Location{
				Name:        entLocation.Name,
				Description: entLocation.Description,
				Tags:       nil,
				PathNode: patnode,
			},
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

	expected := testRooms[1]
	if !reflect.DeepEqual(expected, room) {
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
		MapItem: MapItem{
		},
		Location: Location{
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
		MapItem: MapItem{
		},
		Location: Location{
			Name:        "04",
			Description: "d",
		},
		
	})
	
	newRooms = append(newRooms, Room{
		Id: 4, 
		MapItem: MapItem{
		},
		Location: Location{
			Name:        "04",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id: 5,
		MapItem: MapItem{
		},
		Location: Location{
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
		MapItem: MapItem{
		},
		Location: Location{
			Name:        "06",
			Description: "d",
		},
	})
	
	newRooms = append(newRooms, Room{
		Id: 7, 
		MapItem: MapItem{
		},
		Location: Location{
			Name:        "07",
			Description: "d",
		},
	})
	
	newRooms = append(newRooms, Room{
		Id: 8, 
		MapItem: MapItem{
		},
		Location: Location{
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
						for _, node := range a[i].PathNodes {
							if p.(navigation.PathNode).Id == node.Id {
								return true
							}
						}

						return false
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

func TestRoomEntityService_FilterRooms_RoomFilterParam_FloorFilterParam(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("1", "", "", "1") // no floor 0 in test data

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

func TestRoomEntityService_FilterRooms_RoomFilterParam_FloorFilterParam_BadInteger(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("1", "", "", "1") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("resr", "", "", "1")
	if err == nil {
		t.Error("expect error")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_DbCrash(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	db.Exec("drop table rooms")

	_, err := dbService.FilterRooms("1", "", "", "1") // no floor 0 in test data

	if err == nil {
		t.Error("expect no error, got:", err)
	}
}