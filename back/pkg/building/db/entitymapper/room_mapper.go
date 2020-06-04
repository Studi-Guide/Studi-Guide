package entitymapper

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/db/ent/building"
	"studi-guide/pkg/building/db/ent/location"
	"studi-guide/pkg/building/db/ent/mapitem"
	"studi-guide/pkg/building/db/ent/room"
)

func (r *EntityMapper) roomArrayMapper(entRooms []*ent.Room) []Room {
	var rooms []Room

	for _, roomPtr := range entRooms {
		rooms = append(rooms, *r.roomMapper(roomPtr))
	}

	return rooms
}

func (r *EntityMapper) roomMapper(entRoom *ent.Room) *Room {

	entRoom, err := r.client.Room.Query().Where(room.ID(entRoom.ID)).
		WithMapitem(func(q *ent.MapItemQuery) {
			q.WithPathNodes().WithColor().WithBuilding().WithDoors(func(p *ent.DoorQuery) { p.WithPathNode().WithSection() }).WithSections()
		}).
		WithLocation(func(q *ent.LocationQuery) {
			q.WithPathnode().WithTags().WithBuilding()
		}).
		First(r.context)
	if err != nil || entRoom == nil {
		return nil
	}

	rm := Room{
		Id:       entRoom.ID,
		MapItem:  *r.mapItemMapper(entRoom.Edges.Mapitem),
		Location: *r.locationMapper(entRoom.Edges.Location),
	}

	return &rm
}

func (r *EntityMapper) GetAllRooms() ([]Room, error) {

	roomsPtr, err := r.client.Room.Query().WithMapitem().All(r.context)
	if err != nil {
		return nil, err
	}

	var rooms []Room

	for _, roomPtr := range roomsPtr {
		rooms = append(rooms, *r.roomMapper(roomPtr))
	}

	return rooms, nil
}

func (r *EntityMapper) GetRoom(roomName, buildingName, campusName string) (Room, error) {

	q := r.client.Room.Query().Where(room.HasLocationWith(location.NameEQ(roomName)))

	if len(buildingName) > 0 {
		q = q.Where(room.HasMapitemWith(mapitem.HasBuildingWith(building.NameEQ(buildingName))))
	}

	if len(campusName) > 0 {
		// TODO implement campus
	}

	entRoom, err := q.First(r.context)

	if err != nil {
		return Room{}, err
	}

	return *r.roomMapper(entRoom), nil
}

func (r *EntityMapper) AddRoom(room Room) error {

	return r.storeRooms([]Room{room})
}

func (r *EntityMapper) AddRooms(rooms []Room) error {
	return r.storeRooms(rooms)
}

func (r *EntityMapper) FilterRooms(floorFilter, nameFilter, aliasFilter, roomFilter, buildingFilter, campus string) ([]Room, error) {

	var entRooms []*ent.Room
	var err error = nil

	q := r.client.Room.Query()
	if len(roomFilter) > 0 {
		q = q.Where(
			room.Or(
				room.HasLocationWith(location.NameContains(roomFilter)),
				room.HasLocationWith(location.DescriptionContains(roomFilter))))
	} else {

		if len(nameFilter) > 0 {
			q = q.Where(room.HasLocationWith(location.NameContains(nameFilter)))
		}

		if len(buildingFilter) > 0 {
			q = q.Where(room.HasMapitemWith(mapitem.HasBuildingWith(building.NameContains(buildingFilter))))
		}
	}

	if len(floorFilter) > 0 {
		q = q.Where(room.HasMapitemWith(mapitem.Floor(floorFilter)))
	}

	// alias is missing here ...
	entRooms, err = q.WithMapitem().All(r.context)
	if err != nil {
		return nil, err
	}

	return r.roomArrayMapper(entRooms), nil
}

func (r *EntityMapper) storeRooms(rooms []Room) error {
	var errorStr []string

	for _, rm := range rooms {

		log.Printf("Adding room %s", rm.Name)
		if rm.Id != 0 {
			_, err := r.client.Room.Get(r.context, rm.Id)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}

			continue
		}

		entBuilding, err := r.mapBuilding(rm.Location.Building)
		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.PathNode.Id))
		}

		var entNodes []*ent.PathNode
		for _, node := range rm.PathNodes {

			entPathNode, err := r.mapPathNode(node)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}
			entNodes = append(entNodes, entPathNode)
		}

		entNode, err := r.mapPathNode(&rm.Location.PathNode)
		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.PathNode.Id))
		}

		entSections, err := r.mapSectionArray(rm.MapItem.Sections)
		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		}

		entDoors, err := r.mapDoorArray(rm.MapItem.Doors)
		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		}

		entColor, err := r.mapColor(rm.MapItem.Color)
		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		}

		entMapItem, err := r.client.MapItem.Create().
			AddDoors(entDoors...).
			SetColor(entColor).
			SetBuilding(entBuilding).
			AddSections(entSections...).
			AddPathNodes(entNodes...).
			SetFloor(rm.MapItem.Floor).
			Save(r.context)

		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			continue
		}

		entLocation, err := r.client.Location.Create().
			SetName(rm.Location.Name).
			SetDescription(rm.Location.Description).
			SetPathnode(entNode).
			SetBuilding(entBuilding).
			SetFloor(rm.Location.Floor).
			Save(r.context)

		if err != nil {
			log.Fatal("Error adding room:", rm.Name, " Error:", err)
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		}

		entRoom, err := r.client.Room.Create().
			SetMapitem(entMapItem).
			SetLocation(entLocation).
			Save(r.context)

		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		} else {
			log.Println("Added room:", rm, " as:", entRoom)
		}

		if len(rm.Tags) > 0 {
			entTags, err := r.mapTagArray(rm.Tags, entLocation)
			if err == nil && entTags != nil {
				_, err = entLocation.Update().AddTags(entTags...).Save(r.context)
			}
		}

		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		}
	}

	// link pathnodes
	for _, rm := range rooms {
		for _, node := range rm.PathNodes {
			err := r.linkPathNode(node)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" Room: "+rm.Name+" PathNode:"+strconv.Itoa(node.Id))
			}
		}
		for _, door := range rm.MapItem.Doors {
			err := r.linkPathNode(&door.PathNode)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" Room: "+rm.Name+" PathNode:"+strconv.Itoa(door.PathNode.Id))
			}
		}
	}

	var err error
	if len(errorStr) > 0 {
		err = errors.New(strings.Join(errorStr, "; "))
	}

	return err
}
