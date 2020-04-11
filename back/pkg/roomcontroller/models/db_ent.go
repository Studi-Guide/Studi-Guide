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
	"studi-guide/ent/tag"
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

	roomCount, err2 :=	client.Room.Query().Count(ctx)
	log.Println("Found number of rooms:", roomCount)

	if err2 != nil {
		return nil, err2
	}


	return &RoomEntityService{client: client, table: table, context: ctx}, nil
}

func NewRoomEntityService(env *env.Env) (RoomServiceProvider, error) {
	return newRoomEntityService(env)
}

func (r *RoomEntityService) GetAllRooms() ([]Room, error) {

	roomsPtr, err := r.client.Room.Query().WithSections().WithDoors().WithColor().WithPathNodes().WithTags().All(r.context)
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

	room, err := r.client.Room.Query().Where(room.Name(name)).WithSections().WithDoors().WithColor().WithPathNodes().WithTags().First(r.context)
	if err != nil {
		return Room{}, err
	}

	return *r.roomMapper(room), nil
}

func (r *RoomEntityService) AddRoom(room Room) error {

	_, err := r.storeRooms([]Room {room})
	return err
}

func (r *RoomEntityService) AddRooms(rooms []Room) error {
	_, err := r.storeRooms(rooms)
	return err
}

func (r *RoomEntityService) GetAllPathNodes() ([]navigation.PathNode, error) {
	nodesPrt, err := r.client.PathNode.Query().WithLinkedFrom().WithLinkedTo().All(r.context)
	if err != nil {
		return nil, err
	}

	var nodes []navigation.PathNode
	var nodesCache []*navigation.PathNode
	for _, nodePtr := range nodesPrt {

		node := *r.pathNodeMapper(nodePtr, nodesCache, true)
		nodes = append(nodes, node)
		nodesCache = append(nodesCache, &node)
	}

	return nodes, nil
}

func (r *RoomEntityService) FilterRooms(floorFilter, nameFilter, aliasFilter, roomFilter string) ([]Room, error) {

	var entRooms []*ent.Room
	var err error = nil

	if len(roomFilter) > 0 {
		q := r.client.Room.Query().Where(room.Or(room.NameContains(roomFilter), room.DescriptionContains(roomFilter)))
		if floor, err := strconv.Atoi(floorFilter); len(floorFilter) > 0 && err != nil {
			return nil, err
		} else {
			// Just use query when its available
			if len(floorFilter) > 0 {
				q = q.Where(room.FloorEQ(floor))
			}
		}

		entRooms, err = q.WithSections().WithDoors().WithColor().WithPathNodes().All(r.context)
		if err != nil {
			return nil, err
		}


	} else {
		q:= r.client.Room.Query().Where(room.NameContains(nameFilter))
		if floor, err := strconv.Atoi(floorFilter); len(floorFilter) > 0 && err != nil {
			return nil, err
		} else {
			// Just use query when its available
			if len(floorFilter) > 0 {
				q = q.Where(room.FloorEQ(floor))
			}
		}

		// alias is missing here ...
		entRooms, err = q.WithSections().WithDoors().WithColor().WithPathNodes().All(r.context)
		if err != nil {
			return nil, err
		}

	}
	return r.roomArrayMapper(entRooms), nil
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

func (r *RoomEntityService) roomArrayMapper(entRooms []*ent.Room) []Room {
	var rooms []Room

	for _, roomPtr := range entRooms {
		rooms = append(rooms, *r.roomMapper(roomPtr))
	}

	return rooms
}

func (r *RoomEntityService) roomMapper(entRoom *ent.Room) *Room {

	entRoom, err := r.client.Room.Query().Where(room.ID(entRoom.ID)).WithPathNodes().WithColor().WithDoors().WithSections().WithTags().First(r.context)

	rm := Room{
		Id:          entRoom.ID,
		MapItem:MapItem{
			Name:        entRoom.Name,
			Description: entRoom.Description,
			Tags:       nil,
			Doors:       nil,
			Color:       "",
			Sections:    nil,
			Floor:       entRoom.Floor,
		},
		PathNodes:  []*navigation.PathNode{} ,
	}

	c, err := entRoom.Edges.ColorOrErr()
	if err == nil {
		rm.MapItem.Color = c.Color
	}

	d, err := entRoom.Edges.DoorsOrErr()
	if err == nil {
		rm.MapItem.Doors = r.doorArrayMapper(d)
	}

	s, err := entRoom.Edges.SectionsOrErr()
	if err == nil {
		rm.MapItem.Sections = r.sectionArrayMapper(s)
	}

	p, err := entRoom.Edges.PathNodesOrErr()
	if err == nil {
		rm.PathNodes = r.pathNodeArrayMapper(p, []*navigation.PathNode{})
	}

	t, err := entRoom.Edges.TagsOrErr()
	if err == nil {
		rm.Tags = r.tagsArrayMapper(t)
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
		d.PathNode = *r.pathNodeMapper(p, []*navigation.PathNode{}, true)
	}

	return &d
}

func (r *RoomEntityService) pathNodeArrayMapper(pathNodePtr []*ent.PathNode, availableNodes []*navigation.PathNode) []*navigation.PathNode {
	var pathNodes []*navigation.PathNode
	for _, node := range pathNodePtr {
		pathNodes = append(pathNodes, r.pathNodeMapper(node, availableNodes, false))
	}
	return pathNodes
}

func (r *RoomEntityService) pathNodeMapper(entPathNode *ent.PathNode, availableNodes []*navigation.PathNode, resolveConnectedNodes bool) *navigation.PathNode {

	entPathNode, err := r.client.PathNode.Query().Where(pathnode.PathId(entPathNode.PathId)).WithLinkedTo().First(r.context)
	if err != nil {
		return &navigation.PathNode{}
	}

	for _, node := range availableNodes {
		if node.Id == entPathNode.PathId {
			return node
		}
	}

	p := navigation.PathNode{
		Id:             entPathNode.PathId,
		Coordinate:     navigation.Coordinate{X: entPathNode.XCoordinate, Y: entPathNode.YCoordinate, Z: entPathNode.ZCoordinate},
		Group:          nil,
		ConnectedNodes: nil,
	}

	availableNodes = append(availableNodes, &p)
	if resolveConnectedNodes {
		p.ConnectedNodes = r.pathNodeArrayMapper(entPathNode.Edges.LinkedTo, availableNodes)
	}

	return &p
}

func (r *RoomEntityService) tagMapper(entTag *ent.Tag) string {
	return entTag.Name
}

func (r *RoomEntityService) tagsArrayMapper(entTags []*ent.Tag) []string {
	var tags []string
	for _, t := range entTags {
		tags = append(tags, r.tagMapper(t))
	}
	return tags
}


func (r *RoomEntityService) storeRooms(rooms []Room) ([]*ent.Room, error) {
	var entRooms []*ent.Room
	var errorStr []string

	for _, rm := range rooms {
		if rm.Id != 0 {
			entRoom, err:= r.client.Room.Get(r.context, rm.Id)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}
			entRooms = append(entRooms, entRoom)
			continue
		}

		var entNodes []*ent.PathNode
		for _, node := range rm.PathNodes {

			entPathNode, err := r.mapPathNode(node)
			if err != nil {
				errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
			}
			entNodes = append(entNodes, entPathNode)
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

		entRoom, err := r.client.Room.Create().
			SetName(rm.MapItem.Name).
			SetDescription(rm.MapItem.Description).
			AddDoors(entDoors...).
			SetColor(entColor).
			AddSections(entSections...).
			AddPathNodes(entNodes...).
			SetFloor(rm.MapItem.Floor).
			Save(r.context)

		if err != nil {
			errorStr = append(errorStr, err.Error()+" "+strconv.Itoa(rm.Id))
		} else
		{
			log.Println("add room:", rm)
		}

		if len(rm.Tags) > 0 {
			entTags, err := r.mapTagArray(rm.Tags, entRoom)
			if err == nil && entTags != nil {
				_, err = entRoom.Update().AddTags(entTags...).Save(r.context)
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
		for _, door := range rm.MapItem.Doors{
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

	return entRooms, err
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
		node, err := r.client.PathNode.Query().Where(pathnode.PathId(p.Id)).First(r.context)
		if node != nil {
			return  node, nil
		}

		switch t := err.(type) {
		default:
			log.Fatal(t)
			return nil, err
		case *ent.NotFoundError:
			// do nothing
		}
	}


	log.Println("add path node:", p)
	return r.client.PathNode.Create().
		SetPathId(p.Id).
		SetXCoordinate(p.Coordinate.X).
		SetYCoordinate(p.Coordinate.Y).
		SetZCoordinate(p.Coordinate.Z).
		Save(r.context)
}

func (r *RoomEntityService) linkPathNode(pathNode *navigation.PathNode) error {

	var connectedIDs []int

	//Get database IDs
	for _, connectedNode := range pathNode.ConnectedNodes {

		entityConnectedNode, err := r.client.PathNode.Query().Where(pathnode.PathIdEQ(connectedNode.Id)).First(r.context)
		if err != nil {
			return err
		}

		connectedIDs = append(connectedIDs, entityConnectedNode.ID)
	}

	entityNode, err := r.client.PathNode.Query().Where(pathnode.PathIdEQ(pathNode.Id)).First(r.context)
	if err != nil {
		return err
	}

	update := entityNode.Update()
	update.AddLinkedToIDs(connectedIDs...)
	entityNode, err = update.Save(r.context)

	if err != nil {

	}
	return err
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

func (r *RoomEntityService) mapTag(t string, entRoom *ent.Room) (*ent.Tag, error) {
	entTag, err := r.client.Tag.Query().Where(tag.Name(t)).First(r.context)
	if err != nil && entTag == nil {
		entTag, err = r.client.Tag.Create().SetName(t).AddRooms(entRoom).Save(r.context)
		if err != nil {
			return nil, err
		}
	} else {
		entTag, err = entTag.Update().AddRooms(entRoom).Save(r.context)
	}
	return entTag, err
}

func (r *RoomEntityService) mapTagArray(ts []string, entRoom *ent.Room) ([]*ent.Tag, error) {
	var entTags []*ent.Tag
	for _, t := range ts {
		entTag, err := r.mapTag(t, entRoom)
		if err != nil {
			return nil, err
		}
		entTags = append(entTags, entTag)
	}
	return entTags, nil
}