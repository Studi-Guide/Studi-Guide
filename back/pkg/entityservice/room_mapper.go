package entityservice

import (
	"studi-guide/ent"
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

