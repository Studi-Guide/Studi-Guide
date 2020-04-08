package maps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/roomcontroller/models"
)

type MapController struct {
	router   *gin.RouterGroup
	provider models.RoomServiceProvider
}

func NewMapController(router *gin.RouterGroup, provider models.RoomServiceProvider) error {
	r := MapController{router: router, provider: provider}
	r.router.GET("/", r.GetMapItems)
	r.router.GET("/floor/:floor", r.GetMapItemsFromFloor)
	return nil
}

// GetMapItems godoc
// @Summary Get All Map Items
// @Description Gets map items of available rooms and connector spaces (corridor, stairs, etc..) with optional filter parameters
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param name query string false "name of the map items"
// @Param floor query int false "floor of the map items"
// @Param alias query string false "potential alias of the room"
// @Param campus query string false "map item is linked to a certain campus"
// @Param building query string false "map item is linked to a building"
// @Success 200 {array} models.MapItem
// @Router /map [get]
func (l MapController) GetMapItems(c *gin.Context) {
	floor := c.Query("floor")
	name := c.Query("name") //mux.Vars(r)["name"]
	alias := c.Query("alias")
	campus := c.Query("campus")
	building := c.Query("building")

	//TODO include these filters
	coordinate := c.Query("coordinate")
	coordinateDelta := c.Query("coordinate-delta")
	//-----------------------------

	var rooms []models.Room
	var connectors []models.ConnectorSpace
	var err error

	var useFilterApi bool

	if len(coordinate) == 0 && len(floor) == 0 && len(coordinateDelta) == 0 && len(building) == 0 && len(campus) == 0  {
		//rooms, err = l.provider.GetAllRooms()
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		rooms, err = l.provider.FilterRooms(floor, name, alias, "")
	} else {
		rooms, err = l.provider.GetAllRooms()
	}

	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	if useFilterApi {
		connectors, err = l.provider.FilterConnectorSpaces(floor, name, alias, building, campus, nil, nil)
	} else {
		connectors, err = l.provider.GetAllConnectorSpaces()
	}

	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	l.CreateAndSendMapList(rooms, connectors, c)
}

// GetMapItems godoc
// @Summary Get All Map Items from a certain floor
// @Description Gets all map items of available rooms and connector spaces (corridor, stairs, etc..) of a certain floor
// @ID get-room-list-floor
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param floor path int true "filter map items by floor"
// @Success 200 {array} models.MapItem
// @Router /map/floor/{floor} [get]
func (l MapController) GetMapItemsFromFloor(c *gin.Context) {
	floor := c.Param("floor") //mux.Vars(r)["name"]

	rooms, err := l.provider.FilterRooms(floor, "", "", "")
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	connectors, err := l.provider.FilterConnectorSpaces(floor, "", "", "", "", nil, nil )
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	l.CreateAndSendMapList(rooms, connectors, c)
}

// Helper method
func (l MapController) CreateAndSendMapList(rooms []models.Room, connectors []models.ConnectorSpace, c *gin.Context) {
	var mapItems []models.MapItem
	for _, room := range rooms {
		mapItems = append(mapItems, room.MapItem)
	}

	for _, connector := range connectors {
		mapItems = append(mapItems, connector.MapItem)
	}

	c.JSON(http.StatusOK, mapItems)
}