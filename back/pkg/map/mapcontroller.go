package maps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/entityservice"
)

type MapController struct {
	router   *gin.RouterGroup
	provider MapServiceProvider
}

func NewMapController(router *gin.RouterGroup, provider MapServiceProvider) error {
	r := MapController{router: router, provider: provider}
	r.router.GET("", r.GetMapItems)
	r.router.GET("/floor/:floor", r.GetMapItemsFromFloor)
	return nil
}

// GetMapItems godoc
// @Summary Get All Map Items
// @Description Gets map items of available rooms and connector spaces (corridor, stairs, etc..) with optional filter parameters
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param floor query int false "floor of the map items"
// @Param campus query string false "map item is linked to a certain campus"
// @Param building query string false "map item is linked to a building"
// @Success 200 {array} entityservice.MapItem
// @Router /maps [get]
func (l MapController) GetMapItems(c *gin.Context) {
	building := c.Param("building")

	floor := c.Query("floor")
	campus := c.Query("campus")
	if len(building) == 0 {
		building = c.Query("building")
	}

	// TODO implementation of correct building, campus and floor query
	//TODO include these filters
	coordinate := c.Query("coordinate")
	coordinateDelta := c.Query("coordinate-delta")
	//-----------------------------

	var mapItems []entityservice.MapItem
	var err error

	var useFilterApi bool

	if len(coordinate) == 0 && len(floor) == 0 && len(coordinateDelta) == 0 && len(building) == 0 && len(campus) == 0  {
		//rooms, err = l.provider.GetAllRooms()
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		mapItems, err = l.provider.FilterMapItems(floor, building, campus)
	} else {
		mapItems, err = l.provider.GetAllMapItems()
	}

	if err != nil {
		fmt.Println("GetMapItems() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapItems)
}

// GetMapItems godoc
// @Summary Get All Map Items from a certain floor
// @Description Gets all map items of available rooms and connector spaces (corridor, stairs, etc..) of a certain floor
// @ID get-room-list-floor
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param floor path int true "filter map items by floor"
// @Success 200 {array} entityservice.MapItem
// @Router /maps/{floor} [get]
func (l MapController) GetMapItemsFromFloor(c *gin.Context) {
	floor := c.Param("floor") //mux.Vars(r)["name"]

	mapItems, err := l.provider.FilterMapItems(floor, "", "")
	if err != nil {
		fmt.Println("GetMapItemsFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapItems)
}