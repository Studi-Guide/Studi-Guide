package models

import (
	"os"
	"reflect"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
	"testing"
)

func setupRoomEntityService() (*RoomEntityService, error) {
	os.Setenv("DB_DRIVER_NAME", "sqlite3")
	os.Setenv("DB_DATA_SOURCE", ":memory:")

	e := env.NewEnv()

	return newRoomEntityService(e)
}

func TestMapRoom(t *testing.T) {
	r, err := setupRoomEntityService()
	if err != nil || r == nil {
		t.Error("error setting up room entity service:", err, r)
		return
	}
	defer r.client.Close()

	ro := Room{
		Id:          0,
		MapItem:MapItem{
			Name:        "RoomN01",
			Description: "Room Number 1 Special Description",
			Tags:       []string{
				"Tag1",
				"#Tag2",
			},
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
					PathNode: navigation.PathNode{
						Id: 0,
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
			Floor: 1,
		},

		PathNodes: navigation.PathNode{
			Id: 0,
			Coordinate: navigation.Coordinate{
				X: 34,
				Y: 35,
				Z: 36,
			},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}

	entRoom, err := r.mapRoom(&ro)
	if err != nil || entRoom == nil {
		t.Error("expected no error, got:", err, entRoom)
	}

	ro.Id = 1

	entRoom, err = r.mapRoom(&ro)
	if err != nil || entRoom == nil {
		t.Error("expected no error, got:", err, entRoom)
	}

	checkRoom := Room{
		Id:          1,
		MapItem:MapItem{
			Name:        "RoomN01",
			Description: "Room Number 1 Special Description",
			Tags:       []string{
				"Tag1",
				"#Tag2",
			},
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
						Id: 1,
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
			Floor: 1,
		},

		PathNodes: navigation.PathNode{
			Id: 2,
			Coordinate: navigation.Coordinate{
				X: 34,
				Y: 35,
				Z: 36,
			},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}

	retRoom := r.roomMapper(entRoom)
	if !reflect.DeepEqual(checkRoom, *retRoom) {
		t.Error("expected room equality. expected: ", checkRoom, ", actual: ", retRoom)
	}

	ro.Id = 2
	entRoom, err = r.mapRoom(&ro)
	if err == nil || entRoom != nil {
		t.Error("expected error, got:", err, entRoom)
	}

	ro.Id = 0
	ro.Name = "Fancy Room"
	ro.Tags = []string{"Tag1", "Tag3"}
	checkRoom = Room{
		Id:          2,
		MapItem:MapItem{
			Name:        "Fancy Room",
			Description: "Room Number 1 Special Description",
			Tags:       []string{
				"Tag1",
				"Tag3",
			},
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
						Id: 3,
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
			Floor: 1,
		},

		PathNodes: navigation.PathNode{
			Id: 4,
			Coordinate: navigation.Coordinate{
				X: 34,
				Y: 35,
				Z: 36,
			},
			Group:          nil,
			ConnectedNodes: nil,
		},
	}
	entRoom, err = r.mapRoom(&ro)
	if err != nil || entRoom == nil {
		t.Error("expected no error, got:", err, entRoom)
	}


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
			PathNode: navigation.PathNode{},
		},
		{
			Id:       0,
			Section:  Section{},
			PathNode: navigation.PathNode{},
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
			Id:             0,
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
	d.PathNode.Id = 1
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
	d.PathNode.Id = 0

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
			Id:             0,
			Coordinate:     navigation.Coordinate{},
			Group:          nil,
			ConnectedNodes: nil,
		},
		&navigation.PathNode{
			Id:             0,
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

	pathNodes[0].Id = 3
	pathNodes[1].Id = 0

	entPathNodes, err = r.mapPathNodeArray(pathNodes)
	if err == nil {
		t.Error("expected error, got: ", err, entPathNodes)
	}

	if len(entPathNodes) != 1 {
		t.Error("expected 1 pathNode, got: ", len(entPathNodes), entPathNodes)
	}

	if entPathNodes[0].ID != 3 {
		t.Error("expected pathnode id ", pathNodes[1].Id, "got: ", entPathNodes[0].ID)
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

	if err == nil || entPathNode != nil {
		t.Error("expected error and nil pathnode; got:", err, entPathNode)
	}

	pathNode1 := navigation.PathNode{
		Id:             0,
		Coordinate:     navigation.Coordinate{X: 1, Y: 2, Z: 3},
		Group:          nil,
		ConnectedNodes: nil,
	}
	entPathNode, err = r.mapPathNode(&pathNode1)

	if err != nil || entPathNode == nil {
		t.Error("expected nil and pathnode, got error:", err, entPathNode)
	}

	pathNode1.Id = 1
	retPathNode := r.pathNodeMapper(entPathNode)
	if !reflect.DeepEqual(pathNode1, *retPathNode) {
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