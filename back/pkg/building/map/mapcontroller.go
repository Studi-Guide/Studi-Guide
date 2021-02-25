package maps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studi-guide/pkg/building/db/entitymapper"
)

type MapController struct {
	router   *gin.RouterGroup
	provider MapServiceProvider
}

func NewMapController(router *gin.RouterGroup, provider MapServiceProvider) error {
	r := MapController{router: router, provider: provider}
	r.router.GET("", r.GetMapItems)
	r.router.GET("/buildings/:building/floors/:floor", r.GetMapsFromBuildingFloor)
	return nil
}

// GetMapsFromBuildingFloor godoc
// @Summary Get map items of a Building of a floor
// @Description Get map items of a building filtered by floor
// @ID get-mapitems-from-building
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param building path string false "name of the building"
// @Param floor path string false "name of the floor"
// @Success 200 {array} entitymapper.MapItem
// @Router /maps/buildings/{building}/floors/{floor} [get]
func (b MapController) GetMapsFromBuildingFloor(context *gin.Context) {
	building := context.Param("building")
	floor := context.Param("floor")
	maps, err := b.provider.FilterMapItems(floor, building, "")
	if err != nil {
		fmt.Println("GetMapsFromBuildingFloor() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, maps)
}

// GetMapItems godoc
// @Summary Query Map Items
// @Description Gets map items of available rooms and connector spaces (corridor, stairs, etc..) with optional filter parameters
// @Accept  json
// @Produce  json
// @Tags MapController
// @Param pathnodeid query int false "get the map item linked to a certain path node"
// @Param floor query int false "floor of the map items"
// @Param campus query string false "map item is linked to a certain campus"
// @Param building query string false "map item is linked to a building"
// @Success 200 {array} entitymapper.MapItem
// @Router /maps [get]
func (l MapController) GetMapItems(c *gin.Context) {

	pathNodeID := c.Query("pathnodeid")
	floor := c.Query("floor")
	campus := c.Query("campus")
	building := c.Query("building")

	// TODO implementation of correct building, campus and floor query
	//TODO include these filters
	coordinate := c.Query("coordinate")
	coordinateDelta := c.Query("coordinate-delta")
	//-----------------------------

	var mapItems []entitymapper.MapItem
	var err error

	if len(pathNodeID) > 0 {
		var intID int
		var item entitymapper.MapItem
		intID, err = strconv.Atoi(pathNodeID)
		if err == nil {
			item, err = l.provider.GetMapItemByPathNodeID(intID)
			mapItems = append(mapItems, item)
		}
	} else if len(coordinate) == 0 && len(floor) == 0 && len(coordinateDelta) == 0 && len(building) == 0 && len(campus) == 0 {
		mapItems, err = l.provider.GetAllMapItems()
	} else {
		mapItems, err = l.provider.FilterMapItems(floor, building, campus)
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
