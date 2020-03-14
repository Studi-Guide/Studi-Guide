package models

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/door"
	"studi-guide/ent/pathnode"
	"studi-guide/ent/room"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
)

type RoomEntityService struct {
	client  *ent.Client
	context context.Context
	table   string
}

func NewRoomEntityService(env *env.Env) (RoomServiceProvider, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx, err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func (r *RoomEntityService) GetAllRooms() ([]Room, error) {

	roomsPtr, err := r.client.Room.Query().WithSections().WithDoors().WithColor().WithPathNode().All(r.context)
	if err != nil {
		return nil, err
	}

	var rooms []Room

	for _, roomPtr := range roomsPtr {
		rooms = append(rooms, *r.roomMapper(roomPtr))
	}

	return rooms, nil
}

func (r *RoomEntityService) GetRoom(name string) (Room, error) {

	room, err := r.client.Room.Query().Where(room.Name(name)).WithSections().WithDoors().WithColor().WithPathNode().First(r.context)
	if err != nil {
		return Room{}, err
	}

	return *r.roomMapper(room), nil
}

func (r *RoomEntityService) AddRoom(room Room) error {

	_, err := r.client.Room.Create().
		SetFloor(room.Floor).
		SetName(room.Name).
		SetDescription(room.Description).
		AddDoors().
		Save(r.context)
	// Add other stuff like sequence etc.

	if err != nil {
		return err
	}

	return nil
}

func (r *RoomEntityService) AddRooms(rooms []Room) error {
	var errorStr []string
	for _, room := range rooms {
		if err := r.AddRoom(room); err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(room.Id))
			log.Println(err, "room:", room)
		} else {
			log.Println("add room:", room)
		}
	}

	if len(errorStr) > 0 {
		return errors.New(strings.Join(errorStr, "; "))
	}
	return nil
}

func openDB(dbDriverName string, dbSourceName string) (*ent.Client, context.Context, error) {
	client, err := ent.Open(dbDriverName, "file:"+dbSourceName+"?cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}
	//defer client.Close()
	// run the auto migration tool.
	ctx := context.Background()

	// SQLite was developed only for testing, and it does not support the incremental updates for tables.
	// https://entgo.io/docs/dialects/#sqlite
	if _, err := os.Stat(dbSourceName); dbDriverName != "sqlite3" || (dbDriverName == "sqlite3" && os.IsNotExist(err)) {
		log.Println("running one time migration")
		if err := client.Schema.Create(ctx); err != nil {
			return nil, nil, err
		}
	}

	return client, ctx, err
}

func (r *RoomEntityService) roomMapper(entRoom *ent.Room) *Room {

	entRoom, err := r.client.Room.Query().Where(room.ID(entRoom.ID)).WithPathNode().WithColor().WithDoors().WithSections().First(r.context)

	rm := Room{
		Id:          entRoom.ID,
		Name:        entRoom.Name,
		Description: entRoom.Description,
		Alias:       nil,
		Doors:       nil,
		Color:       "",
		Sections:    nil,
		Floor:       entRoom.Floor,
		PathNode:    navigation.PathNode{},
	}

	c, err := entRoom.Edges.ColorOrErr()
	if err == nil {
		rm.Color = c.Color
	}

	d, err := entRoom.Edges.DoorsOrErr()
	if err == nil {
		rm.Doors = r.doorArrayMapper(d)
	}

	s, err := entRoom.Edges.SectionsOrErr()
	if err == nil {
		rm.Sections = r.sectionArrayMapper(s)
	}

	p, err := entRoom.Edges.PathNodeOrErr()
	if err == nil {
		rm.PathNode = *r.pathNodeMapper(p)
	}

	return &rm
}

func (r *RoomEntityService) sectionArrayMapper(sections []*ent.Section) []Section {
	var s []Section
	for _, seq := range sections {
		s = append(s, Section{
			Id:    seq.ID,
			Start: navigation.Coordinate{},
			End:   navigation.Coordinate{},
		})
	}
	return s
}

func (r *RoomEntityService) sectionMapper(s *ent.Section) *Section {
	return &Section{
		Id:    s.ID,
		Start: navigation.Coordinate{X: s.XStart, Y: s.YStart, Z: 0},
		End:   navigation.Coordinate{X: s.XEnd, Y: s.YEnd, Z: 0},
	}
}

func (r *RoomEntityService) doorArrayMapper(doors []*ent.Door) []Door {
	var d []Door
	for _, door := range doors {
		d = append(d, *r.doorMapper(door))
	}
	return d
}

func (r *RoomEntityService) doorMapper(entDoor *ent.Door) *Door {

	entDoor, err := r.client.Door.Query().WithPathNode().WithOwner().WithSection().Where(door.ID(entDoor.ID)).First(r.context)
	if err != nil {
		return &Door{}
	}

	d := Door{
		Id:       entDoor.ID,
		Section:  Section{},
		PathNode: navigation.PathNode{},
	}

	s, err := entDoor.Edges.SectionOrErr()
	if err == nil {
		d.Section = *r.sectionMapper(s)
	}

	p, err := entDoor.Edges.PathNodeOrErr()
	if err == nil {
		d.PathNode = *r.pathNodeMapper(p)
	}

	return &d
}

func (r *RoomEntityService) pathNodeArrayMapper(pathNodePtr []*ent.PathNode) []*navigation.PathNode {
	var pathNodes []*navigation.PathNode
	for _, node := range pathNodePtr {
		pathNodes = append(pathNodes, r.pathNodeMapper(node))
	}
	return pathNodes
}

func (r *RoomEntityService) pathNodeMapper(entPathNode *ent.PathNode) *navigation.PathNode {

	entPathNode, err := r.client.PathNode.Query().WithPathGroups().WithLinkedTo().Where(pathnode.ID(entPathNode.ID)).First(r.context)
	if err != nil {
		return &navigation.PathNode{}
	}

	p := navigation.PathNode{
		Id:             entPathNode.ID,
		Coordinate:     navigation.Coordinate{X: entPathNode.XCoordinate, Y: entPathNode.YCoordinate, Z: entPathNode.ZCoordinate},
		Group:          nil,
		ConnectedNodes: r.pathNodeArrayMapper(entPathNode.Edges.LinkedTo),
	}

	return &p
}
