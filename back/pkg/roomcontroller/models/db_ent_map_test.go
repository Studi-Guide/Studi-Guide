package models

import (
	"os"
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
		Name:        "RoomN01",
		Description: "Room Number 1 Special Description",
		Alias:       nil,
		Doors: []Door{
			{
				Id: 0,
				Section: Section{
					Id: 0,
					Start: navigation.Coordinate{
						X: 0,
						Y: 0,
						Z: 0,
					},
					End: navigation.Coordinate{
						X: 0,
						Y: 0,
						Z: 0,
					},
				},
				PathNode: navigation.PathNode{
					Id: 0,
					Coordinate: navigation.Coordinate{
						X: 0,
						Y: 0,
						Z: 0,
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
					X: 0,
					Y: 0,
					Z: 0,
				},
				End: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
			{
				Id: 0,
				Start: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
				End: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
			{
				Id: 0,
				Start: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
				End: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
			{
				Id: 0,
				Start: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
				End: navigation.Coordinate{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
		},
		Floor: 0,
		PathNode: navigation.PathNode{
			Id: 0,
			Coordinate: navigation.Coordinate{
				X: 0,
				Y: 0,
				Z: 0,
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

	ro.Id = 2
	entRoom, err = r.mapRoom(&ro)
	if err == nil || entRoom != nil {
		t.Error("expected error, got:", err, entRoom)
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
		Start: navigation.Coordinate{},
		End:   navigation.Coordinate{},
	}

	entSection, err := r.mapSection(&s)
	if err != nil || entSection == nil {
		t.Error("expected section, got err:", err, ", section: ", entSection)
	}

	s.Id = 1

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
		Id:       0,
		Section:  Section{},
		PathNode: navigation.PathNode{},
	}

	entDoor, err := r.mapDoor(&d)
	if err != nil || entDoor == nil {
		t.Error("expected no error, got:", err, entDoor)
	}

	d.Id = 1

	entDoor, err = r.mapDoor(&d)
	if err != nil || entDoor == nil {
		goto E
	}

	d.Id = 0

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

	entPathNode, err = r.mapPathNode(&navigation.PathNode{
		Id:             0,
		Coordinate:     navigation.Coordinate{},
		Group:          nil,
		ConnectedNodes: nil,
	})

	if err != nil || entPathNode == nil {
		t.Error("expected nil and pathnode, got error:", err, entPathNode)
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
	if err != nil || color == nil {
		t.Error("expected nil and color, got:", err, ", ", color)
	}

	color2 := "#007BFF"

	color, err = r.mapColor(color2)
	if err != nil || color == nil {
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
	if err != nil || color == nil {
		t.Error("expected nil and color, got:", err, ", ", color)
	}

	color6 := "#0aF9"

	color, err = r.mapColor(color6)
	if err == nil || color != nil {
		t.Error("expected error, got: ", err, color)
	}

}
