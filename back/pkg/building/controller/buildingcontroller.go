package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/building/model"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
	"studi-guide/pkg/roomcontroller/controllers"
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

	roomRouter := b.router.Group("/:building/rooms")
	if err := controllers.NewRoomController(roomRouter, b.roomProvider); err != nil {
		return err
	}

	locationRouter := b.router.Group("/:building/locations")
	if err := location.NewLocationController(locationRouter, b.locationProvider); err != nil {
		return err
	}

	mapsRouter := b.router.Group("/:building/maps")
	if err := maps.NewMapController(mapsRouter, b.mapProvider); err != nil {
		return err
	}

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
func (b BuildingController) GetBuildingByName(c *gin.Context) {
	name := c.Param("building")

	building, err := b.buildingProvider.GetBuilding(name)
	if err != nil {
		fmt.Println("GetBuilding() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, building)
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
