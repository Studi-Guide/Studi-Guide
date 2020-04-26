package entityservice

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/building"
	"studi-guide/ent/location"
	"studi-guide/ent/mapitem"
	"studi-guide/ent/room"
)

func (r *EntityService) roomArrayMapper(entRooms []*ent.Room) []Room {
	var rooms []Room

	for _, roomPtr := range entRooms {
		rooms = append(rooms, *r.roomMapper(roomPtr))
	}

	return rooms
}

func (r *EntityService) roomMapper(entRoom *ent.Room) *Room {

	entRoom, err := r.client.Room.Query().Where(room.ID(entRoom.ID)).
		WithMapitem(func(q *ent.MapItemQuery) {
			q.WithPathNodes().WithColor().WithDoors(func(p *ent.DoorQuery) { p.WithPathNode().WithSection() }).WithSections()
		}).
		WithLocation(func(q *ent.LocationQuery) {
			q.WithPathnode().WithTags()
		}).
		First(r.context)
	if err != nil || entRoom == nil {
		return nil
	}

	entMapItem, err := r.client.Room.QueryMapitem(entRoom).WithPathNodes().
		WithColor().
		WithDoors().
		WithBuilding().
		WithSections().
		First(r.context)
	if err != nil || entMapItem == nil {
		return nil
	}

	entLocation, err := r.client.Room.QueryLocation(entRoom).WithTags().WithPathnode().First(r.context)
	if err != nil || entLocation == nil {
		return nil
	}

	rm := Room{
		Id:       entRoom.ID,
		MapItem:  *r.mapItemMapper(entMapItem),
		Location: *r.locationMapper(entLocation),
	}

	return &rm
}

func (r *EntityService) GetAllRooms() ([]Room, error) {

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

func (r *EntityService) GetRoom(roomName, buildingName, campusName string) (Room, error) {

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

func (r *EntityService) AddRoom(room Room) error {

	return r.storeRooms([]Room{room})
}

func (r *EntityService) AddRooms(rooms []Room) error {
	return r.storeRooms(rooms)
}

func (r *EntityService) FilterRooms(floorFilter, nameFilter, aliasFilter, roomFilter, buildingFilter, campus string) ([]Room, error) {

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

	if floor, err := strconv.Atoi(floorFilter); len(floorFilter) > 0 && err != nil {
		return nil, err
	} else {
		// Just use query when its available
		if len(floorFilter) > 0 {
			q = q.Where(room.HasMapitemWith(mapitem.FloorEQ(floor)))
		}
	}

	// alias is missing here ...
	entRooms, err = q.WithMapitem().All(r.context)
	if err != nil {
		return nil, err
	}

	return r.roomArrayMapper(entRooms), nil
}

func (r *EntityService) storeRooms(rooms []Room) error {
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

		entBuilding, err := r.mapBuilding(rm.Building)
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
			SetFloor(rm.Location.Floor).
			Save(r.context)

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
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}
		}
		for _, door := range rm.MapItem.Doors {
			err := r.linkPathNode(&door.PathNode)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}
		}
	}

	var err error
	if len(errorStr) > 0 {
		err = errors.New(strings.Join(errorStr, "; "))
	}

	return err
}

