package entitymapper

import (
	"context"
	"database/sql"
	"github.com/ahmetb/go-linq/v3"
	fbsql "github.com/facebook/ent/dialect/sql"
	"log"
	"os"
	"reflect"
	"strconv"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/env"
	"studi-guide/pkg/file"
	"studi-guide/pkg/navigation"
	"testing"
)

var testRooms []Room

func setupTestRoomDbService() (*EntityMapper, *sql.DB) {
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

	_, err = client.RssFeed.
		Create().
		SetName("testfeed").
		SetURL("http://www.testfeed.de/rss.xml").
		Save(ctx)

	address, err := client.Address.Create().
		SetCity("Munich").
		SetCountry("Germany").
		SetStreet("Am Platzl").
		SetNumber("1").
		SetPLZ(80331).Save(ctx)

	campus, err := client.Campus.Create().
		SetName("testcampus").
		SetShortName("TC").
		SetLongitude(0).
		SetLatitude(0).
		Save(ctx)

	if err != nil {
		log.Fatal(err)
	}

	building, _ := client.Building.Create().SetName("main").SetAddress(address).SetCampus(campus).Save(ctx)
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
			SetFloor(strconv.Itoa(i)).
			SetBuilding(building).
			Save(ctx)

		if err != nil {
			log.Println("error creating map item:", err)
		}

		entLocation, err := client.Location.Create().
			SetName(strconv.Itoa(i)).
			SetPathnode(pathNode).
			SetFloor(strconv.Itoa(i)).
			SetBuilding(building).
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
			Id: pathNode.ID,
			Coordinate: navigation.Coordinate{
				X: pathNode.XCoordinate,
				Y: pathNode.YCoordinate,
				Z: pathNode.ZCoordinate,
			}}

		testRooms = append(testRooms, Room{
			Id: entRoom.ID,
			MapItem: MapItem{
				Doors: []Door{{
					Id:       door.ID,
					Section:  Section{Id: sequence.ID},
					PathNode: patnode,
				}},
				Color:     "",
				Sections:  nil,
				Floor:     strconv.Itoa(i),
				Building:  "main",
				PathNodes: []*navigation.PathNode{&patnode},
			},
			Location: Location{
				Id:          entLocation.ID,
				Name:        entLocation.Name,
				Description: entLocation.Description,
				Tags:        nil,
				PathNode:    patnode,
				Building:    "main",
				Floor:       strconv.Itoa(i),
			},
		})
	}

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	client.Campus.Create().
		SetShortName("HB").
		SetName("HB Hofbräu Haus").
		SetLatitude(48.1378).
		SetLongitude(11.5797).Save(ctx)

	dbService := EntityMapper{client: client, env: env.NewEnv(), context: ctx}

	return &dbService, drv.DB()
}

func TestNewRoomDbService(t *testing.T) {
	os.Setenv("DB_DRIVER_NAME", "some_driver")
	os.Setenv("DB_DATA_SOURCE", ":some_source")

	e := env.NewEnv()

	dbService, err := NewEntityMapper(e)
	if err == nil {
		t.Error("expected error; got: ", err)
	}
	if !reflect.ValueOf(dbService).IsNil() {
		t.Error("expected: ", nil, "; got: ", dbService)
	}

	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e = env.NewEnv()
	dbService, err = NewEntityMapper(e)
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
		for i := range a {
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

	room, err := dbService.GetRoom(strconv.Itoa(2), "", "")
	if err != nil {
		t.Error(err)
	}

	expected := testRooms[1]
	if !reflect.DeepEqual(expected, room) {
		t.Error("expected: ", testRooms[1], "; got: ", room)
	}

	room, err = dbService.GetRoom("4", "", "")
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
		Id:      4,
		MapItem: MapItem{},
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
		Id:      4,
		MapItem: MapItem{},
		Location: Location{
			Name:        "04",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id:      4,
		MapItem: MapItem{},
		Location: Location{
			Name:        "04",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id:      5,
		MapItem: MapItem{},
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
		Id:      6,
		MapItem: MapItem{},
		Location: Location{
			Name:        "06",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id:      7,
		MapItem: MapItem{},
		Location: Location{
			Name:        "07",
			Description: "d",
		},
	})

	newRooms = append(newRooms, Room{
		Id:      8,
		MapItem: MapItem{},
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
		for i := range a {
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

	getConnectors, err := dbService.FilterRooms("1", "", "", "", "", "")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	compare := func(a []Room, b []Room) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if !reflect.DeepEqual(a[i], b[i]) {
				return false
			}
		}
		return true
	}

	var expected []Room
	linq.From(testRooms).Where(func(p interface{}) bool { return p.(Room).MapItem.Floor == "1" }).ToSlice(&expected)

	if !compare(expected, getConnectors) {
		t.Error("expected: ", expected, "; got: ", getConnectors)
	}

	db.Exec("drop table rooms")

	getConnectors, err = dbService.FilterRooms("1", "", "", "", "", "")
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

	rooms, err := dbService.FilterRooms("1", "", "", "", "", "") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("abcd", "", "", "", "", "")
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_RoomFilterParam(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("", "", "", "1", "", "") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("", "", "", "abcd", "", "")
	if err != nil {
		t.Error("expect no error", err, " got not nil")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_NameFilterParam(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("", "1", "", "", "", "") // no floor 0 in test data
	if err != nil {
		t.Error("expect no error, got:", err)
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	if !linq.From(rooms).All(func(p interface{}) bool { return p.(Room).Location.Name == "1" }) {
		t.Error("expect room array with name == 1 and not anything else")
	}

	rooms, err = dbService.FilterRooms("", "foobar", "", "", "", "")
	if err != nil {
		t.Error("expect no error", err, " got not nil")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_RoomFilterParam_FloorFilterParam(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rooms, err := dbService.FilterRooms("1", "", "", "1", "", "") // no floor 0 in test data

	if err != nil {
		t.Error("expect no error, got:", err)

		rooms, err = dbService.FilterRooms("abcd", "", "", "", "", "")
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	if err != nil {
		t.Error("expect no error, got:", err)

		rooms, err = dbService.FilterRooms("ABCD", "", "", "", "", "")
	}
	if rooms == nil {
		t.Error("expect room array but is nil")
	}

	rooms, err = dbService.FilterRooms("", "", "", "abcd", "", "")
	if err != nil {
		t.Error("expect no error", err, " got not nil")
	}
	if rooms != nil {
		t.Error("expect nil room array, got: ", rooms)
	}
}

func TestRoomEntityService_FilterRooms_DbCrash(t *testing.T) {
	dbService, db := setupTestRoomDbService()

	db.Exec("drop table rooms")

	_, err := dbService.FilterRooms("1", "", "", "1", "", "") // no floor 0 in test data

	if err == nil {
		t.Error("expect no error, got:", err)
	}
}

func TestEntityService_GetAllLocations(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getLocations, err := dbService.GetAllLocations()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var testLocations []Location
	for _, room := range testRooms {
		testLocations = append(testLocations, room.Location)
	}

	if !reflect.DeepEqual(testLocations, getLocations) {
		t.Error("expected: ", testLocations, "; got: ", getLocations)
	}
}

func TestEntityService_FilterLocations(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getLocations, err := dbService.FilterLocations("1", "", "1", "main", "testcampus")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var testLocations []Location
	testLocations = append(testLocations, testRooms[0].Location)
	if !reflect.DeepEqual(testLocations, getLocations) {
		t.Error("expected: ", testLocations, "; got: ", getLocations)
	}
}

func TestEntityService_GetLocation(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getLocation, err := dbService.GetLocation("1", "", "")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if !reflect.DeepEqual(testRooms[0].Location, getLocation) {
		t.Error("expected: ", testRooms[0].Location, "; got: ", getLocation)
	}
}

func TestEntityMapper_AddLocation(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	loc := Location{
		Id:          999,
		Name:        "loc",
		Description: "location in some place",
		Tags:        []string{"a", "b"},
		Floor:       "f",
		Building:    "b",
		PathNode: navigation.PathNode{
			Id: 1111,
			Coordinate: navigation.Coordinate{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}

	if dbService.AddLocation(loc) == nil {
		t.Error("expected error because of building")
	}

	loc.Building = "main"

	if err := dbService.AddLocation(loc); err != nil {
		t.Error(err)
	}

	_, err := dbService.client.PathNode.Create().Save(dbService.context)
	if err != nil {
		t.Error("err creating pathnode", err)
	}

	loc.Name = "another location"
	loc.Id = 998
	loc.PathNode.Id = 1112
	loc.PathNode.ConnectedNodes = append(loc.PathNode.ConnectedNodes, &navigation.PathNode{
		Id:             1111,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})
	loc.Icon = "my-icon"

	if err := dbService.AddLocation(loc); err != nil {
		t.Error(err)
	}

	loc.Name = "anothernother location"
	loc.Id = 997
	loc.PathNode.Id = 1113
	loc.PathNode.ConnectedNodes = append(loc.PathNode.ConnectedNodes, &navigation.PathNode{
		Id:             1112,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})
	loc.Images = append(loc.Images, file.File{
		Name: "myfile.png",
		Path: "/file.png",
	})

	if err := dbService.AddLocation(loc); err != nil {
		t.Error(err)
	}

	loc.Name = "bla location"
	loc.Id = 996
	loc.PathNode.Id = 1114
	loc.PathNode.ConnectedNodes = append(loc.PathNode.ConnectedNodes, &navigation.PathNode{
		Id:             1113,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})
	loc.Images = append(loc.Images, file.File{
		Name: "myfile.png",
		Path: "/file.png",
	})

	if err := dbService.AddLocation(loc); err != nil {
		t.Error(err)
	}
}

func TestEntityService_GetAllMapItems(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getMapItems, err := dbService.GetAllMapItems()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var expectMapItems []MapItem
	for _, room := range testRooms {
		expectMapItems = append(expectMapItems, room.MapItem)
	}

	if !reflect.DeepEqual(expectMapItems, getMapItems) {
		t.Error("expected: ", expectMapItems, "; got: ", getMapItems)
	}
}

func TestEntityService_FilterMapItems(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getMapItems, err := dbService.FilterMapItems("2", "", "")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var expectMapItems []MapItem
	expectMapItems = append(expectMapItems, testRooms[1].MapItem)

	if !reflect.DeepEqual(expectMapItems, getMapItems) {
		t.Error("expected: ", expectMapItems, "; got: ", getMapItems)
	}
}

func TestEntityService_FilterMapItems_Building(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getMapItems, err := dbService.FilterMapItems("2", "main", "")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var expectMapItems []MapItem
	expectMapItems = append(expectMapItems, testRooms[1].MapItem)

	if !reflect.DeepEqual(expectMapItems, getMapItems) {
		t.Error("expected: ", expectMapItems, "; got: ", getMapItems)
	}
}

func TestEntityService_FilterMapItems_Campus(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getMapItems, err := dbService.FilterMapItems("", "", "testcampus")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	var expectMapItems []MapItem
	expectMapItems = append(expectMapItems, testRooms[0].MapItem, testRooms[1].MapItem, testRooms[2].MapItem)

	if !reflect.DeepEqual(expectMapItems, getMapItems) {
		t.Error("expected: ", expectMapItems, "; got: ", getMapItems)
	}
}

func TestEntityService_FilterMapItems_Building_Negativ(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	getMapItems, _ := dbService.FilterMapItems("2", "main2", "")
	if len(getMapItems) != 0 {
		t.Error("expected: ", nil, "; got: ", getMapItems)
	}
}

func TestEntityService_CampusEntity(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	campusArray, err := dbService.GetAllCampus()
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if len(campusArray) == 0 {
		t.Error("expected something got: nothing")
	}

	if campusArray[1].Name != "HB Hofbräu Haus" {
		t.Error("expected: ", "HB Hofbräu Haus", "; got: ", campusArray[0].Name)
	}

	campus, err := dbService.GetCampus("HB")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if campus.Name != "HB Hofbräu Haus" {
		t.Error("expected: ", "HB Hofbräu Haus", "; got: ", campus.Name)
	}

	campusArray, err = dbService.FilterCampus("HB")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if len(campusArray) == 0 {
		t.Error("expected something got: nothing")
	}

	if campusArray[0].Name != "HB Hofbräu Haus" {
		t.Error("expected: ", "HB Hofbräu Haus", "; got: ", campusArray[0].Name)
	}
}
func TestEntityService_CampusEntity_Negative(t *testing.T) {
	dbService, _ := setupTestRoomDbService()
	_, err := dbService.GetCampus("HBB")
	if err == nil {
		t.Error("expected error got: ", nil)
	}

	campusArray, err := dbService.FilterCampus("HBB")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if len(campusArray) != 0 {
		t.Error("expected: ", nil, "; got: ", campusArray)
	}
}

func TestEntityMapper_AddCampus(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	address := ent.Address{
		Street:  "BlaStreet",
		Number:  "10",
		PLZ:     11111,
		City:    "BlaTown",
		Country: "BlaLand",
	}

	invalidAddress := ent.Address{
		City:    "BlaTown",
		Country: "BlaLand",
	}

	testcampus := ent.Campus{
		ShortName: "Test",
		Name:      "TESTTEST",
		Longitude: 12180840.92938,
		Latitude:  120480124.29323,
		Edges: ent.CampusEdges{
			Buildings: []*ent.Building{{
				ID:   1,
				Name: "Test",
				Edges: ent.BuildingEdges{
					Body: []*ent.Coordinate{{
						Latitude:  20,
						Longitude: 10,
					}},
					Address: &address,
				},
			},
				{
					ID:   2,
					Name: "Test2",
					Edges: ent.BuildingEdges{
						Address: &address,
					},
				},
				{
					ID:   1,
					Name: "Test3",
					Edges: ent.BuildingEdges{
						Address: &address,
					},
				}},
		},
	}

	testcampus2 := ent.Campus{
		ShortName: "Test2",
		Name:      "TESTTEST2",
		Longitude: 12180840.92938,
		Latitude:  120480124.29323,
		Edges: ent.CampusEdges{
			Buildings: []*ent.Building{{
				ID:   2,
				Name: "Test2",
				Edges: ent.BuildingEdges{
					Body: []*ent.Coordinate{{
						Latitude:  20,
						Longitude: 10,
					}},
					Address: &invalidAddress,
				},
			}},
		},
	}

	err := dbService.AddCampus(testcampus)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	err = dbService.AddCampus(testcampus)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	realValue, err := dbService.GetCampus("Test")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if realValue.ShortName != testcampus.ShortName ||
		realValue.Name != testcampus.Name ||
		realValue.Latitude != testcampus.Latitude ||
		realValue.Longitude != testcampus.Longitude ||
		realValue.Edges.Buildings[0].Edges.Address.Street != testcampus.Edges.Buildings[0].Edges.Address.Street ||
		realValue.Edges.Buildings[0].Edges.Address.City != testcampus.Edges.Buildings[0].Edges.Address.City ||
		realValue.Edges.Buildings[0].Edges.Address.Country != testcampus.Edges.Buildings[0].Edges.Address.Country ||
		realValue.Edges.Buildings[0].Edges.Address.PLZ != testcampus.Edges.Buildings[0].Edges.Address.PLZ {
		t.Error("expected equal but got ", testcampus, "; real: ", realValue)
	}

	err = dbService.AddCampus(testcampus2)
	if err == nil {
		t.Error("expected error; got: ", nil)
	}
}

func TestEntityService_RssFeedEntity(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	rssfeed, err := dbService.GetRssFeed("testfeed")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if rssfeed.Name != "testfeed" {
		t.Error("expected: ", "testfeed", "; got: ", rssfeed.Name)
	}

	if rssfeed.URL != "http://www.testfeed.de/rss.xml" {
		t.Error("expected: ", "http://www.testfeed.de/rss.xml", "; got: ", rssfeed.URL)
	}
}
func TestEntityService_RssFeed_Negative(t *testing.T) {
	dbService, _ := setupTestRoomDbService()
	_, err := dbService.GetRssFeed("tesutatoha")
	if err == nil {
		t.Error("expected error got: ", nil)
	}
}

func TestEntityMapper_AddRssFeed(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	testfeed := ent.RssFeed{
		URL:  "http://www.testtestfeed.de/rss.xml",
		Name: "Tesitjaoistgnaj",
	}

	err := dbService.AddRssFeed(testfeed)
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	rssfeed, err := dbService.GetRssFeed("Tesitjaoistgnaj")
	if err != nil {
		t.Error("expected: ", nil, "; got: ", err)
	}

	if rssfeed.Name != "Tesitjaoistgnaj" {
		t.Error("expected: ", "Tesitjaoistgnaj", "; got: ", rssfeed.Name)
	}

	if rssfeed.URL != "http://www.testtestfeed.de/rss.xml" {
		t.Error("expected: ", "http://www.testtestfeed.de/rss.xml", "; got: ", rssfeed.URL)
	}
}

func TestEntityMapper_AddRssFeed_InvalidAddress(t *testing.T) {
	dbService, _ := setupTestRoomDbService()

	testfeed := ent.RssFeed{
		URL:  "http://invalid.com/perl.cgi?key= | http://web-site.com/cgi-bin/perl.cgi?key1=value1&key2",
		Name: "Tesitjaoistgnaj",
	}

	err := dbService.AddRssFeed(testfeed)
	if err == nil {
		t.Error("expected error got: ", nil)
	}
}

func TestEntityMapper_mapFile(t *testing.T) {
	entL := ent.File{
		ID:    0,
		Name:  "Name of File",
		Path:  "/Path/to/file",
		Edges: ent.FileEdges{},
	}

	dbService, _ := setupTestRoomDbService()

	f := dbService.mapFile(&entL)

	if f.Name != entL.Name || f.Path != entL.Path {
		t.Error("map file test failed, name or path not equal")
	}
}
