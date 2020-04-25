package entityservice

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"studi-guide/ent"
	"studi-guide/ent/building"
	"studi-guide/ent/location"
	"studi-guide/ent/mapitem"
	"studi-guide/ent/room"
	"studi-guide/ent/tag"
	"studi-guide/pkg/env"
	"studi-guide/pkg/navigation"

	_ "github.com/mattn/go-sqlite3"
)

type EntityService struct {
	client  *ent.Client
	context context.Context
	table   string
}

func newEntityService(env *env.Env) (*EntityService, error) {
	driverName := env.DbDriverName()
	dataSourceName := env.DbDataSource()
	table := "rooms"
	client, ctx, err := openDB(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	roomCount, _ := client.Room.Query().Count(ctx)
	log.Println("Found number of rooms:", roomCount)
	return &EntityService{client: client, table: table, context: ctx}, nil
}

func NewEntityService(env *env.Env) (*EntityService, error) {
	return newEntityService(env)
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

func (r *EntityService) GetRoom(name string) (Room, error) {

	entRoom, err := r.client.Room.Query().Where(room.HasLocationWith(location.Name(name))).
		First(r.context)

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

func (r *EntityService) GetAllPathNodes() ([]navigation.PathNode, error) {
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

func (r *EntityService) GetAllLocations() ([]Location, error) {
	entLoactions, err := r.client.Location.Query().
		WithTags().
		WithPathnode().
		All(r.context)
	if err != nil {
		return nil, err
	}

	return r.locationArrayMapper(entLoactions), nil
}

func (r *EntityService) GetLocation(name string) (Location, error) {
	entLocation, err := r.client.Location.Query().WithPathnode().WithTags().Where(location.NameEQ(name)).First(r.context)
	if err != nil {
		return Location{}, err
	}
	return *r.locationMapper(entLocation), nil
}

func (r *EntityService) FilterLocations(name, tagStr, floor, building, campus string) ([]Location, error) {

	query := r.client.Location.Query().
		WithPathnode().WithTags()

	if len(name) > 0 {
		query = query.Where(location.NameContains(name))
	}

	if len(tagStr) > 0 {
		query = query.Where(location.HasTagsWith(tag.NameContains(tagStr)))
	}

	if len(floor) > 0 {
		iFloor, err := strconv.Atoi(floor)
		if err != nil {
			return nil, err
		}
		query = query.Where(location.FloorEQ(iFloor))
	}

	if len(building) > 0 {
		// Todo query building
	}

	if len(campus) > 0 {
		// Todo query campus
	}

	entLocations, err := query.All(r.context)
	if err != nil {
		return nil, err
	}
	return r.locationArrayMapper(entLocations), nil
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

func (r *EntityService) GetAllMapItems() ([]MapItem, error) {
	entMapItems, err := r.client.MapItem.Query().
		WithPathNodes().
		WithColor().
		WithBuilding().
		WithDoors().
		WithSections().
		All(r.context)
	if err != nil {
		return nil, err
	}

	return r.mapItemArrayMapper(entMapItems), nil
}

func (r *EntityService) FilterMapItems(floor, buildingFilter, campus string) ([]MapItem, error) {
	iFloor, err := strconv.Atoi(floor)
	if err != nil {
		return nil, err
	}

	mapQuery := r.client.MapItem.Query()

	if len(buildingFilter) > 0 {
		mapQuery.Where(mapitem.HasBuildingWith(building.Name(buildingFilter)))
	}

	// TODO Missing items: campus
	entMapItems, err := mapQuery.Where(mapitem.Floor(iFloor)).
		WithBuilding().
		WithDoors().
		WithSections().
		WithPathNodes().
		All(r.context)
	if err != nil {
		return nil, err
	}
	return r.mapItemArrayMapper(entMapItems), nil
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

func (r *EntityService) linkPathNode(pathNode *navigation.PathNode) error {

	var connectedIDs []int

	//Get database IDs
	for _, connectedNode := range pathNode.ConnectedNodes {

		entityConnectedNode, err := r.client.PathNode.Get(r.context, connectedNode.Id)
		if err != nil {
			return err
		}

		connectedIDs = append(connectedIDs, entityConnectedNode.ID)
	}

	entityNode, _ := r.client.PathNode.Get(r.context, pathNode.Id)

	update := entityNode.Update()
	update.AddLinkedToIDs(connectedIDs...)
	entityNode, err := update.Save(r.context)
	return err
}

