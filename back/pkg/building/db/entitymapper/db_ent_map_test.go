package entitymapper

import (
	"os"
	"reflect"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/location"
	"studi-guide/pkg/building/db/ent/mapitem"
	"studi-guide/pkg/building/db/ent/room"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
	"testing"
)

func setupRoomEntityService() (*EntityMapper, error) {
	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e := env.NewEnv()

	return newEntityMapper(e)
}

func TestMapRoom(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	node1 := navigation.PathNode{
		Id: 2,
		Coordinate: navigation.Coordinate{
			X: 34,
			Y: 35,
			Z: 36,
		},
	}

	node2 := navigation.PathNode{
		Id: 1,
		Coordinate: navigation.Coordinate{
			X: 7,
			Y: 8,
			Z: 9,
		},
		Group:          nil,
		ConnectedNodes: []*navigation.PathNode{&node1},
	}

	node1.ConnectedNodes = []*navigation.PathNode{&node2}

	ro := Room{
		Id: 0,
		MapItem: MapItem{
			Doors: []Door{
				{
					Id: 0,
					Section: Section{
						Id: 0,
						Start: navigation.Coordinate{
							X: 1,
							Y: 2,
							Z: 3,
						},
						End: navigation.Coordinate{
							X: 4,
							Y: 5,
							Z: 6,
						},
					},
					PathNode: node2,
				},
			},
			Color: "#5682a3",
			Sections: []Section{
				{
					Id: 0,
					Start: navigation.Coordinate{
						X: 10,
						Y: 11,
						Z: 12,
					},
					End: navigation.Coordinate{
						X: 13,
						Y: 14,
						Z: 15,
					},
				},
				{
					Id: 0,
					Start: navigation.Coordinate{
						X: 16,
						Y: 17,
						Z: 18,
					},
					End: navigation.Coordinate{
						X: 19,
						Y: 20,
						Z: 21,
					},
				},
				{
					Id: 0,
					Start: navigation.Coordinate{
						X: 22,
						Y: 23,
						Z: 24,
					},
					End: navigation.Coordinate{
						X: 25,
						Y: 26,
						Z: 27,
					},
				},
				{
					Id: 0,
					Start: navigation.Coordinate{
						X: 28,
						Y: 29,
						Z: 30,
					},
					End: navigation.Coordinate{
						X: 31,
						Y: 32,
						Z: 33,
					},
				},
			},
			Floor:    "1",
			Building: "main",
		},

		Location: Location{
			Name:        "RoomN01",
			Building:    "main",
			Description: "Room Number 1 Special Description",
			Tags: []string{
				"Tag1",
				"#Tag2",
			},
			Floor:    "1",
			PathNode: node1,
		}}

	err = r.storeRooms([]Room{ro})
	if err != nil {
		t.Error("expected no error, got:", err)
	}

	ro.Id = 1

	err = r.storeRooms([]Room{ro})
	if err == nil {
		t.Error("expected error, got:", err)
	}

	checkRoom := Room{
		Id: 1,
		MapItem: MapItem{
			Doors: []Door{
				{
					Id: 1,
					Section: Section{
						Id: 5,
						Start: navigation.Coordinate{
							X: 1,
							Y: 2,
							Z: 3,
						},
						End: navigation.Coordinate{
							X: 4,
							Y: 5,
							Z: 6,
						},
					},
					PathNode: navigation.PathNode{
						Id:             0,
						Coordinate:     navigation.Coordinate{},
						Group:          nil,
						ConnectedNodes: nil,
					},
				},
			},
			Color: "#5682a3",
			Sections: []Section{
				{
					Id: 1,
					Start: navigation.Coordinate{
						X: 10,
						Y: 11,
						Z: 12,
					},
					End: navigation.Coordinate{
						X: 13,
						Y: 14,
						Z: 15,
					},
				},
				{
					Id: 2,
					Start: navigation.Coordinate{
						X: 16,
						Y: 17,
						Z: 18,
					},
					End: navigation.Coordinate{
						X: 19,
						Y: 20,
						Z: 21,
					},
				},
				{
					Id: 3,
					Start: navigation.Coordinate{
						X: 22,
						Y: 23,
						Z: 24,
					},
					End: navigation.Coordinate{
						X: 25,
						Y: 26,
						Z: 27,
					},
				},
				{
					Id: 4,
					Start: navigation.Coordinate{
						X: 28,
						Y: 29,
						Z: 30,
					},
					End: navigation.Coordinate{
						X: 31,
						Y: 32,
						Z: 33,
					},
				},
			},
			Floor:    "1",
			Building: "main",
		},

		Location: Location{
			Id:          1,
			Name:        "RoomN01",
			Description: "Room Number 1 Special Description",
			Tags: []string{
				"Tag1",
				"#Tag2",
			},
			Floor:    "1",
			Building: "main",
			PathNode: navigation.PathNode{
				Id: 2,
				Coordinate: navigation.Coordinate{
					X: 34,
					Y: 35,
					Z: 36,
				},
			},
		}}

	entRoom, _ := r.client.Room.Query().Where(room.HasLocationWith(location.NameEQ(ro.Location.Name))).First(r.context)
	retRoom := r.roomMapper(entRoom)

	// exclude pathnodes since they are reference types
	retRoom.PathNodes = nil
	retRoom.Doors[0].PathNode = navigation.PathNode{}
	if !reflect.DeepEqual(checkRoom, *retRoom) {
		t.Error("expected room equality. expected: ", checkRoom, ", actual: ", retRoom)
	}

	ro.Id = 2
	err = r.storeRooms([]Room{ro})
	if err == nil {
		t.Error("expected error, got:", err)
	}

	checkRoom = Room{
		Id: 2,
		MapItem: MapItem{
			Doors: []Door{
				{
					Id: 2,
					Section: Section{
						Id: 10,
						Start: navigation.Coordinate{
							X: 1,
							Y: 2,
							Z: 3,
						},
						End: navigation.Coordinate{
							X: 4,
							Y: 5,
							Z: 6,
						},
					},
					PathNode: navigation.PathNode{
						Id: 20,
						Coordinate: navigation.Coordinate{
							X: 7,
							Y: 8,
							Z: 9,
						},
						Group:          nil,
						ConnectedNodes: nil,
					},
				},
			},
			Color: "#5682a3",
			Sections: []Section{
				{
					Id: 6,
					Start: navigation.Coordinate{
						X: 10,
						Y: 11,
						Z: 12,
					},
					End: navigation.Coordinate{
						X: 13,
						Y: 14,
						Z: 15,
					},
				},
				{
					Id: 7,
					Start: navigation.Coordinate{
						X: 16,
						Y: 17,
						Z: 18,
					},
					End: navigation.Coordinate{
						X: 19,
						Y: 20,
						Z: 21,
					},
				},
				{
					Id: 8,
					Start: navigation.Coordinate{
						X: 22,
						Y: 23,
						Z: 24,
					},
					End: navigation.Coordinate{
						X: 25,
						Y: 26,
						Z: 27,
					},
				},
				{
					Id: 9,
					Start: navigation.Coordinate{
						X: 28,
						Y: 29,
						Z: 30,
					},
					End: navigation.Coordinate{
						X: 31,
						Y: 32,
						Z: 33,
					},
				},
			},
			Floor:    "1",
			Building: "main",
		},
		Location: Location{
			Id:          2,
			Name:        "Fancy Room",
			Building:    "main",
			Description: "Room Number 1 Special Description",
			Tags: []string{
				"Tag1",
				"Tag3",
			},
			PathNode: navigation.PathNode{
				Id: 499,
				Coordinate: navigation.Coordinate{
					X: 34,
					Y: 35,
					Z: 36,
				},
				Group:          nil,
				ConnectedNodes: nil,
			},
		},
	}
	ro = checkRoom
	ro.Id = 0
	ro.Name = "Fancy Room"
	ro.Location.Id = 0
	ro.Tags = []string{"Tag1", "Tag3"}
	var sections []Section
	for _, section := range ro.Sections {
		section.Id = 0
		sections = append(sections, section)
	}
	ro.Sections = sections

	var doors []Door
	for _, door := range ro.Doors {
		door.Id = 0
		door.Section.Id = 0
		doors = append(doors, door)
	}

	ro.Doors = doors
	err = r.storeRooms([]Room{ro})
	if err != nil {
		t.Error("expected no error, got:", err)
	}

	entRoom, _ = r.client.Room.Query().Where(room.HasLocationWith(location.NameEQ(ro.Name))).First(r.context)
	retRoom = r.roomMapper(entRoom)
	if !reflect.DeepEqual(checkRoom, *retRoom) {
		t.Error("expected room equality. expected: ", checkRoom, ", actual: ", retRoom)
	}
}

func TestMapSectionArray(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	sections := []Section{
		{
			Id:    0,
			Start: navigation.Coordinate{},
			End:   navigation.Coordinate{},
		},
	}

	entSections, err := r.mapSectionArray(sections)
	if err != nil || entSections == nil {
		t.Error("expected no error, got:", err, entSections)
	}

	sections[0].Id = 1
	entSections, err = r.mapSectionArray(sections)
	if err != nil || entSections == nil {
		t.Error("expected no error, got:", err, entSections)
	}

	sections[0].Id = 2
	entSections, err = r.mapSectionArray(sections)
	if err == nil || entSections != nil {
		t.Error("expected error, got: ", err, entSections)
	}

}

func TestMapSection(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	s := Section{
		Id:    0,
		Start: navigation.Coordinate{X: 1, Y: 2, Z: 3},
		End:   navigation.Coordinate{X: 4, Y: 5, Z: 6},
	}

	entSection, err := r.mapSection(&s)
	if err != nil || entSection == nil {
		t.Error("expected section, got err:", err, ", section: ", entSection)
	}

	s.Id = 1
	retSection := r.sectionMapper(entSection)
	if !reflect.DeepEqual(s, *retSection) {
		t.Error("expected equal sections: expected:", s, ", actual:", retSection)
	}

	entSection, err = r.mapSection(&s)
	if err != nil || entSection == nil {
		t.Error("expected section, got err:", err, ", section: ", entSection)
	}

	s.Id = 2
	entSection, err = r.mapSection(&s)
	if err == nil || entSection != nil {
		t.Error("expected error, got: err:", err, ", section: ", entSection)
	}
}

func TestMapDoorArray(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	doors := []Door{
		{
			Id:       0,
			Section:  Section{},
			PathNode: navigation.PathNode{Id: 1},
		},
		{
			Id:       0,
			Section:  Section{},
			PathNode: navigation.PathNode{Id: 2},
		},
	}

	entDoors, err := r.mapDoorArray(doors)
	if err != nil || entDoors == nil {
		t.Error("expected no error, got:", err, entDoors)
	}

	doors[0].Id = 1
	doors[1].Id = 3

	entDoors, err = r.mapDoorArray(doors)
	if err == nil || entDoors != nil {
		t.Error("expected error, got: ", err, entDoors)
	}
}

func TestMapDoor(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	d := Door{
		Id: 0,
		Section: Section{
			Id:    0,
			Start: navigation.Coordinate{X: 1, Y: 2, Z: 3},
			End:   navigation.Coordinate{X: 4, Y: 5, Z: 6},
		},
		PathNode: navigation.PathNode{
			Id:             1,
			Coordinate:     navigation.Coordinate{X: 7, Y: 8, Z: 9},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}

	entDoor, err := r.mapDoor(&d)
	if err != nil || entDoor == nil {
		t.Error("expected no error, got:", err, entDoor)
	}

	d.Id = 1
	d.Section.Id = 1
	retDoor := r.doorMapper(entDoor)
	if !reflect.DeepEqual(d, *retDoor) {
		t.Error("expected door equality. expected:", d, ", actual:", retDoor)
	}

	entDoor, err = r.mapDoor(&d)
	if err != nil || entDoor == nil {
		goto E
	}

	d.Id = 0
	d.Section.Id = 0
	d.PathNode.Id = 3

	entDoor, err = r.mapDoor(&d)
	if err != nil || entDoor == nil {
		goto E
	}
	//if entDoor.ID != 1 {
	//	t.Error("expected door ID 1, got: ", entDoor)
	//}

	return

E:
	t.Error("expected no error, got:", err, entDoor)

}

func TestMapPathNodeArray(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	pathNodes := []*navigation.PathNode{
		&navigation.PathNode{
			Id:             1,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
		&navigation.PathNode{
			Id:             2,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}

	entPathNodes, err := r.mapPathNodeArray(pathNodes)
	if err != nil {
		t.Error("expected no error, got: ", err, entPathNodes)
	}

	pathNodes[0].Id = 1
	pathNodes[1].Id = 2

	entPathNodes, err = r.mapPathNodeArray(pathNodes)
	if err != nil {
		t.Error("expected no error, got: ", err, entPathNodes)
	}

	//pathNodes[0].Id = 3
	//pathNodes[1].Id = 0

	//entPathNodes, err = r.mapPathNodeArray(pathNodes)
	//if err == nil {
	//	t.Error("expected error, got: ", err, entPathNodes)
	//}

	if len(entPathNodes) != 2 {
		t.Error("expected 2 pathNode, got: ", len(entPathNodes), entPathNodes)
	}

	if entPathNodes[0].ID != 1 {
		t.Error("expected pathnode id ", pathNodes[1].Id, "got: ", entPathNodes[0].ID)
	}

}

func TestMapPathNode_Exception(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	entPathNode, err := r.mapPathNode(&navigation.PathNode{
		Id:             1,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})

	r.client.PathNode.DeleteOneID(entPathNode.ID).Exec(r.context)
	result := r.pathNodeMapper(entPathNode, []*navigation.PathNode{}, true)

	if !reflect.DeepEqual(*result, navigation.PathNode{}) {
		t.Error("expected no error and  pathnode; got:", err, entPathNode)
	}
}

func TestMapPathNode(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	entPathNode, err := r.mapPathNode(&navigation.PathNode{
		Id:             1,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})

	if err != nil || entPathNode == nil {
		t.Error("expected no error and  pathnode; got:", err, entPathNode)
	}

	pathNode1 := navigation.PathNode{
		Id:             2,
		Coordinate:     navigation.Coordinate{X: 1, Y: 2, Z: 3},
		Group:          nil,
		ConnectedNodes: nil,
	}
	entPathNode, err = r.mapPathNode(&pathNode1)

	if err != nil || entPathNode == nil {
		t.Error("expected nil and pathnode, got error:", err, entPathNode)
	}

	retPathNode := r.pathNodeMapper(entPathNode, []*navigation.PathNode{}, true)
	if !reflect.DeepEqual(pathNode1, *retPathNode) {
		t.Error("expected pathnode equality, expected:", pathNode1, ", actual: ", retPathNode)
	}

	retPathNode2 := r.pathNodeMapper(entPathNode, []*navigation.PathNode{}, true)
	if !reflect.DeepEqual(*retPathNode2, *retPathNode) {
		t.Error("expected pathnode equality, expected:", pathNode1, ", actual: ", retPathNode)
	}

	entPathNode, err = r.mapPathNode(&navigation.PathNode{
		Id:             1,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})

	if err != nil || entPathNode == nil {
		t.Error("expected pathnode, got:", err, entPathNode)
	}

}

func TestMapColor(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	color1 := "#007bff"

	color, err := r.mapColor(color1)
	if err != nil || color == nil || color.Color != color1 {
		t.Error("expected nil and color, got:", err, ", ", color)
	}

	color2 := "#007BFF"

	color, err = r.mapColor(color2)
	if err != nil || color == nil || color.Color != color2 {
		t.Error("expected nil and color, got:", err, ", ", color)
	}

	color3 := "#zabx56"

	color, err = r.mapColor(color3)
	if err == nil || color != nil {
		t.Error("expected error, got: ", err, color)
	}

	color4 := "#1234567"

	color, err = r.mapColor(color4)
	if err == nil || color != nil {
		t.Error("expected error, got: ", err, color)
	}

	color5 := "#1aB"

	color, err = r.mapColor(color5)
	if err != nil || color == nil || color.Color != color5 {
		t.Error("expected nil and color, got:", err, ", ", color)
	}

	color6 := "#0aF9"

	color, err = r.mapColor(color6)
	if err == nil || color != nil {
		t.Error("expected error, got: ", err, color)
	}

}

func TestEntityMapper_GetMapItemByPathNodeID(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil {
		t.Error(err)
	}
	defer r.client.Close()

	// client create test campus
	address, err := r.client.Address.Create().
		SetCity("Munich").
		SetCountry("Germany").
		SetStreet("Am Platzl").
		SetNumber("1").
		SetPLZ(80331).Save(r.context)

	if err != nil {
		t.Error(err)
	}

	campus, err := r.client.Campus.Create().
		SetName("testcampus").
		SetShortName("TC").
		SetLatitude(0).
		SetLongitude(0).
		Save(r.context)

	if err != nil {
		t.Error(err)
	}

	pNode, err := r.client.PathNode.Create().Save(r.context)
	if err != nil {
		t.Error(err)
	}

	building, err := r.client.Building.Create().SetName("building").SetAddress(address).SetCampus(campus).Save(r.context)
	if err != nil {
		t.Error(err)
	}

	mItem, err := r.client.MapItem.Create().
		AddPathNodes(pNode).
		SetBuilding(building).
		Save(r.context)
	if err != nil {
		t.Error(err)
	}

	mItem, err = r.client.MapItem.Query().
		WithBuilding(
			func(buildingQuery *ent.BuildingQuery) { buildingQuery.WithCampus().WithAddress() }).
		WithPathNodes().
		Where(mapitem.ID(mItem.ID)).
		First(r.context)
	if err != nil {
		t.Error(err)
	}
	mappedItem := *r.mapItemMapper(mItem)

	checkItem, err := r.GetMapItemByPathNodeID(pNode.ID)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(mappedItem.Building, checkItem.Building) {
		t.Error(mappedItem.Building, checkItem.Building)
	}

	if !reflect.DeepEqual(mappedItem.PathNodes, checkItem.PathNodes) {
		t.Error(mappedItem.PathNodes, checkItem.PathNodes)
	}

	floors, err := r.getFloorsFromBuilding(mItem.Edges.Building)
	if len(floors) == 0 {
		t.Error("expected something and got:", err)
	}

	if floors[0] != "0" {
		t.Error("expected '0' and got:", floors[0])
	}

	_, err = r.GetAllBuildings()
	if err != nil {
		t.Error(err)
	}

	testbuilding, err := r.GetBuilding("building")
	if err != nil {
		t.Error(err)
	}

	if testbuilding.Name != building.Name {
		t.Error("expected ", building.Name, " and got:", testbuilding.Name)
	}
	_, err = r.FilterBuildings("building")
	if err != nil {
		t.Error(err)
	}
}
