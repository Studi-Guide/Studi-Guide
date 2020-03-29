package models

import (
	"context"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/color"
	"studi-guide/ent/door"
	"studi-guide/ent/pathnode"
	"studi-guide/ent/room"
	"studi-guide/ent/section"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"
)

type RoomEntityService struct {
	client  *ent.Client
	context context.Context
	table   string
}

func newRoomEntityService(env *env.Env) (*RoomEntityService, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx, err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func NewRoomEntityService(env *env.Env) (RoomServiceProvider, error) {
	return newRoomEntityService(env)
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

	_, err := r.mapRoom(&room)

	return err
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

func (r *RoomEntityService) GetAllPathNodes() ([]navigation.PathNode, error) {
	nodesPrt, err := r.client.PathNode.Query().WithLinkedFrom().WithLinkedTo().All(r.context)
	if err != nil {
		return nil, err
	}

	var nodes []navigation.PathNode

	for _, nodePtr := range nodesPrt {
		nodes = append(nodes, *r.pathNodeMapper(nodePtr))
	}

	return nodes, nil
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
		s = append(s, *r.sectionMapper(seq))
	}
	return s
}

func (r *RoomEntityService) sectionMapper(s *ent.Section) *Section {

	if s == nil {
		return nil
	}

	sec, err := r.client.Section.Query().Where(section.ID(s.ID)).First(r.context)
	if err != nil {
		return nil
	}

	return &Section{
		Id:    s.ID,
		Start: navigation.Coordinate{X: sec.XStart, Y: sec.YStart, Z: sec.ZStart},
		End:   navigation.Coordinate{X: sec.XEnd, Y: sec.YEnd, Z: sec.ZEnd},
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

func (r *RoomEntityService) mapRoom(rm *Room) (*ent.Room, error) {

	if rm.Id != 0 {
		return r.client.Room.Query().Where(room.ID(rm.Id)).First(r.context)
	}

	entSections, err := r.mapSectionArray(rm.Sections)
	if err != nil {
		return nil, err
	}

	entDoors, err := r.mapDoorArray(rm.Doors)
	if err != nil {
		return nil, err
	}

	entPathNode, err := r.mapPathNode(&rm.PathNode)
	if err != nil {
		return nil, err
	}

	entColor, err := r.mapColor(rm.Color)
	if err != nil {
		return nil, err
	}

	return r.client.Room.Create().
		SetName(rm.Name).
		SetDescription(rm.Description).
		AddDoors(entDoors...).
		SetColor(entColor).
		AddSections(entSections...).
		SetFloor(rm.Floor).
		SetPathNode(entPathNode).
		Save(r.context)
}

func (r *RoomEntityService) mapSectionArray(sections []Section) ([]*ent.Section, error) {

	var entSections []*ent.Section

	for _, s := range sections {
		entS, err := r.mapSection(&s)
		if err != nil {
			return nil, err
		}
		entSections = append(entSections, entS)
	}

	return entSections, nil

}

func (r *RoomEntityService) mapSection(s *Section) (*ent.Section, error) {

	if s.Id != 0 {
		return r.client.Section.Query().Where(section.ID(s.Id)).First(r.context)
	}

	return r.client.Section.Create().
		SetXStart(s.Start.X).SetXEnd(s.End.X).
		SetYStart(s.Start.Y).SetYEnd(s.End.Y).
		SetZStart(s.Start.Z).SetZEnd(s.End.Z).
		Save(r.context)
}

func (r *RoomEntityService) mapDoorArray(doors []Door) ([]*ent.Door, error) {
	var entDoors []*ent.Door

	for _, d := range doors {
		entD, err := r.mapDoor(&d)
		if err != nil {
			return nil, err
		}
		entDoors = append(entDoors, entD)
	}

	return entDoors, nil
}

func (r *RoomEntityService) mapDoor(d *Door) (*ent.Door, error) {

	if d.Id != 0 {
		return r.client.Door.Query().Where(door.ID(d.Id)).First(r.context)
	}

	sec, err := r.mapSection(&d.Section)
	if err != nil {
		return nil, err
	}

	pNode, err := r.mapPathNode(&d.PathNode)
	if err != nil {
		return nil, err
	}

	return r.client.Door.Create().
		SetPathNodeID(pNode.ID).
		SetSectionID(sec.ID).
		Save(r.context)
}

func (r *RoomEntityService) mapPathNodeArray(pathNodePtr []*navigation.PathNode) ([]*ent.PathNode, error) {

	var entPathNodes []*ent.PathNode

	var errorStrs []string

	for _, ptr := range pathNodePtr {
		p, err := r.mapPathNode(ptr)
		if err != nil {
			errorStrs = append(errorStrs, err.Error())
			continue
		}
		entPathNodes = append(entPathNodes, p)
	}

	if len(errorStrs) != 0 {
		return entPathNodes, errors.New(strings.Join(errorStrs, ","))
	}

	return entPathNodes, nil
}

func (r *RoomEntityService) mapPathNode(p *navigation.PathNode) (*ent.PathNode, error) {

	if p.Id != 0 {
		return r.client.PathNode.Query().Where(pathnode.ID(p.Id)).First(r.context)
	}

	return r.client.PathNode.Create().
		SetXCoordinate(p.Coordinate.X).
		SetYCoordinate(p.Coordinate.Y).
		SetZCoordinate(p.Coordinate.Z).
		Save(r.context)
}

func (r *RoomEntityService) mapColor(c string) (*ent.Color, error) {

	format := "#[0-9a-fA-F]{3}$|#[0-9a-fA-F]{6}$"
	reg := regexp.MustCompile(format)

	if !reg.MatchString(c) {
		return nil, errors.New("color " + c + " does not match the required format: " + format)
	}

	col, err := r.client.Color.Query().Where(color.Color(c)).First(r.context)

	if err != nil && col == nil {
		col, err = r.client.Color.Create().SetName(c).SetColor(c).Save(r.context)
		if err != nil {
			return nil, err
		}
	}

	return col, nil
}
