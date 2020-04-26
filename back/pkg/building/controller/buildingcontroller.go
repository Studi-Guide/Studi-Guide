package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/building/model"
	"studi-guide/pkg/location"
	maps "studi-guide/pkg/map"
	"studi-guide/pkg/room/models"
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
	b.router.GET("/:building/floors/:floor/locations", b.GetLocationsFromBuildingFloor)

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

// @Summary Get Rooms of a Building of a floor
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Param floor path string false "name of the floor"
// @Success 200 {array} entityservice.Room
// @Router /buildings/{building}/floors/{floor}/rooms [get]
func (b BuildingController) GetRoomsFromBuildingFloor(context *gin.Context) {
	building := context.Param("building")
	floor := context.Param("floor")
	rooms, err := b.roomProvider.FilterRooms(floor, "", "", "", building, "")
	if err != nil {
		fmt.Println("GetRoomsFromBuildingFloor() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, rooms)
}

// @Summary Get Rooms of a Building of a floor
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Param floor path string false "name of the floor"
// @Success 200 {array} entityservice.MapItem
// @Router /buildings/{building}/floors/{floor}/maps [get]
func (b BuildingController) GetMapsFromBuildingFloor(context *gin.Context) {
	building := context.Param("building")
	floor := context.Param("floor")
	maps, err := b.mapProvider.FilterMapItems(floor, building, "")
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

// @Summary Get Rooms of a Building of a floor
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Param floor path string false "name of the floor"
// @Success 200 {array} entityservice.Location
// @Router /buildings/{building}/floors/{floor}/locations [get]
func (b BuildingController) GetLocationsFromBuildingFloor(context *gin.Context) {
	building := context.Param("building")
	floor := context.Param("floor")
	maps, err := b.mapProvider.FilterMapItems(floor, building, "")
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