package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/building/model"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
	"studi-guide/pkg/roomcontroller/models"
)

type BuildingController struct {
	router           *gin.RouterGroup
	buildingProvider BuildingProvider
	roomProvider     models.RoomServiceProvider
	locationProvider location.LocationProvider
	mapProvider      maps.MapServiceProvider
}

func NewBuildingController(router *gin.RouterGroup, buildingProvider BuildingProvider,
	roomProvider models.RoomServiceProvider, locationProvider location.LocationProvider,
	mapProvider maps.MapServiceProvider) error {
	b := BuildingController{
		router:           router,
		buildingProvider: buildingProvider,
		roomProvider:     roomProvider,
		locationProvider: locationProvider,
		mapProvider:      mapProvider,
	}

	b.router.GET("", b.GetBuildings)
	b.router.GET("/:building", b.GetBuildingByName)
	b.router.GET("/:building/floors", b.GetFloorsFromBuilding)
	b.router.GET("/:building/floors/:floor/rooms", b.GetRoomsFromBuildingFloor)
	b.router.GET("/:building/floors/:floor/maps", b.GetMapsFromBuildingFloor)
	b.router.GET("/:building/floors/:floor/locations", b.GetLocationFromBuildingFloor)

	return nil
}

// GetLocations godoc
// @Summary Get All Buildings
// @Description Gets buildings by possible filters
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param name query string false "name of the building"
// @Success 200 {array} model.Building
// @Router /buildings [get]
func (b BuildingController) GetBuildings(c *gin.Context) {
	name := c.Query("name")

	var buildings []model.Building
	var err error

	var useFilterApi bool

	if len(name) == 0 {
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		buildings, err = b.buildingProvider.FilterBuildings(name)
	} else {
		buildings, err = b.buildingProvider.GetAllBuildings()
	}

	if err != nil {
		fmt.Println("GetBuildings() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, buildings)
}

// GetBuildingByName godoc
// @Summary Get Building by a certain name
// @Description Get one building by name
// @ID get-building-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string true "name of the building"
// @Success 200 {array} model.Building
// @Router /buildings/{building} [get]
func (b BuildingController) GetBuildingByName(context *gin.Context) {
	name := context.Param("building")

	building, err := b.buildingProvider.GetBuilding(name)
	if err != nil {
		fmt.Println("GetBuilding() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, building)
}

func (b BuildingController) GetFloorsFromBuilding(context *gin.Context) {
	buildingName := context.Param("building")
	building, _ := b.buildingProvider.GetBuilding(buildingName)
	floors, err := b.buildingProvider.GetFloorsFromBuilding(building)

	if err != nil {
		fmt.Println("GetFloorsFromBuilding() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, floors)
}

func (b BuildingController) GetRoomsFromBuildingFloor(context *gin.Context) {

}

func (b BuildingController) GetMapsFromBuildingFloor(context *gin.Context) {

}

func (b BuildingController) GetLocationFromBuildingFloor(context *gin.Context) {

}

// from room controller

// @Summary Get Rooms of a Building by a certain name
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Success 200 {array} model.Building
// @Router /buildings/{building}/rooms [get]

// @Summary Get Rooms of a Building by a certain name
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string true "name of the building"
// @Param room path string true "name of the room"
// @Success 200 {array} entityservice.Room
// @Router /buildings/{building}/rooms/{room} [get]

/*
 /buildings
 /buildings/:name

 /buildings/:name/rooms
 /buildings/:name/locations
 /buildings/:name/mapitems
 /buildings/:name/floors

 /buildings/:name/rooms/:name
 /buildings/:name/locations/:name
 /buildings/:name/mapitems/:id
 /buildings/:name/floors/

 /buildings?filter=value
 /buildings/:name/rooms?filter=value
 /buildings/:name/locations?filter=value
 /buildings/:name/mapitems?filter=value

 /locations
 /rooms
 /mapitems

 /locations/:name
 /rooms/:name
 /mapitems/:name

 /locations?filter=value
 /rooms?filter?value
 /mapitems?filter=value

*/
