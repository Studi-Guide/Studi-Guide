package maps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studi-guide/pkg/roomcontroller/models"
)

type MapController struct {
	router   *gin.RouterGroup
	provider models.RoomServiceProvider
}

func NewMapController(router *gin.RouterGroup, provider models.RoomServiceProvider) error {
	r := MapController{router: router, provider: provider}
	r.router.GET("/", r.GetMapItems)
	r.router.GET("/:floor", r.GetMapItemsFromFloor)
	return nil
}

func (l MapController) GetMapItems(c *gin.Context) {
	rooms, err := l.provider.GetAllRooms()
	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	connectors, err := l.provider.GetAllConnectorSpaces()
	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	l.CreateAndSendMapList(rooms, connectors, c)
}


func (l MapController) GetMapItemsFromFloor(c *gin.Context) {
	floor := c.Query("floor") //mux.Vars(r)["name"]

	floorInt, err := strconv.Atoi(floor)
	if err != nil {
		// handle error
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	rooms, err := l.provider.GetRoomsFromFloor(floorInt)
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	connectors, err := l.provider.GetConnectorsFromFloor(floorInt)
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
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

	c.JSON(http.StatusOK, rooms)
}