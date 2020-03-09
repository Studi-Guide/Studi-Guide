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
	"studi-guide/ent/room"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
)

type RoomEntityService struct {
	client    *ent.Client
	context   context.Context
	table string
}

func NewRoomEntityService(env *env.Env) (RoomServiceProvider, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx,  err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func (r *RoomEntityService) GetAllRooms() ([]Room, error) {

	roomsPtr, err := r.client.Room.Query().WithColor().WithDoors().WithPathNode().WithSequences().All(r.context)
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

	room, err := r.client.Room.Query().Where(room.Name(name)).WithColor().WithDoors().WithPathNode().WithSequences().First(r.context)
	if err != nil {
		return Room{}, err
	}



	return *r.roomMapper(room), nil
}

func (r *RoomEntityService) AddRoom(room Room) error {

	_, err :=r.client.Room.Create().
		SetFloor(room.Floor).
		SetName(room.Name).
		SetDescription(room.Description).
		Save(r.context)
	// Add other stuff like sequence etc.

	if err != nil {
		return  err
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

func (r *RoomEntityService)  roomMapper(room *ent.Room) (*Room) {

	rm := Room{
		Id:          room.ID,
		Name:        room.Name,
		Description: room.Description,
		Alias:       nil,
		Doors:       r.doorArrayMapper(room.Edges.Doors),
		Color:       "",
		Sections:    r.sequenceArrayMapper(room.Edges.Sequences),
		Floor:       room.Floor,
	}

	if room.Edges.Color != nil {
		rm.Color = room.Edges.Color.Color
	}

	return &rm
}

func (r *RoomEntityService)  sequenceMapper(sequence *ent.Sequence) (*Sequence) {
	return &Sequence{
		Id:    sequence.ID,
		Start: navigation.Coordinate{X: sequence.XStart, Y: sequence.YStart, Z: 0},
		End:   navigation.Coordinate{X: sequence.XEnd, Y: sequence.YEnd, Z: 0},
	}
}

func (r *RoomEntityService)  sequenceArrayMapper(sequences []*ent.Sequence) ([]Sequence) {
	var s []Sequence
	for _, seq := range sequences {
		s = append(s, Sequence{
			Id:    seq.ID,
			Start: navigation.Coordinate{},
			End:   navigation.Coordinate{},
		})
	}
	return s
}

func (r *RoomEntityService) doorArrayMapper(doors []*ent.Door) ([]Door) {
	var d []Door
	for _, door := range doors {
		d = append(d, *r.doorMapper(door))
	}
	return d
}

func (r *RoomEntityService)  doorMapper(door *ent.Door) (*Door) {
	return &Door{
		Id:       door.ID,
		Sequence: *r.sequenceMapper(door.Edges.Sequence),
		PathNode: navigation.PathNode{
			Id:             door.Edges.PathNode.ID,
			Coordinate:     navigation.Coordinate{X: door.Edges.PathNode.XCoordinate, Y: door.Edges.PathNode.YCoordinate, Z: door.Edges.PathNode.ZCoordinate},
			Group:          nil,
			ConnectedNodes: r.pathNodeArrayMapper(door.Edges.PathNode.Edges.LinkedTo),
		},
		}
}

func (r *RoomEntityService)  pathNodeMapper(pathNode *ent.PathNode) (*navigation.PathNode) {
	return &navigation.PathNode{
		Id:             pathNode.ID,
		Coordinate:     navigation.Coordinate{X: pathNode.XCoordinate, Y: pathNode.YCoordinate, Z: pathNode.ZCoordinate},
		Group:          nil,
		ConnectedNodes: r.pathNodeArrayMapper(pathNode.Edges.LinkedTo),
	}
}

func (r *RoomEntityService)  pathNodeArrayMapper(pathNodePtr []*ent.PathNode) ([]navigation.PathNode) {
	var pathNodes []navigation.PathNode
	for _, node := range pathNodePtr {
		pathNodes = append(pathNodes, *r.pathNodeMapper(node))
	}
	return pathNodes
}