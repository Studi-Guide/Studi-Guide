package info

import (
	"fmt"
	"net/http"
	"studi-guide/pkg/building/db/ent"
	"studi-guide/pkg/building/location"
	maps "studi-guide/pkg/building/map"
	"studi-guide/pkg/building/room/models"

	"github.com/gin-gonic/gin"
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

// GetBuildings godoc
// @Summary Get Buildings by name
// @ID get-buildings
// @Description Gets a list of buildings filtered by name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param name query string false "name of the building"
// @Success 200 {array} ent.Building
// @Router /buildings [get]
func (b BuildingController) GetBuildings(c *gin.Context) {
	name := c.Query("name")

	var buildings []*ent.Building
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
// @Success 200 {object} ent.Building
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

// @Summary Get Rooms of a Building of a floor
// @Description Get one building by name
// @ID get-building-room-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Param floor path string false "name of the floor"
// @Success 200 {array} entitymapper.Room
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
// @Success 200 {array} entitymapper.MapItem
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
// @Success 200 {array} entitymapper.Location
// @Router /buildings/{building}/floors/{floor}/locations [get]
func (b BuildingController) GetLocationsFromBuildingFloor(context *gin.Context) {
	building := context.Param("building")
	floor := context.Param("floor")
	location, err := b.locationProvider.FilterLocations("", "", floor, building, "")
	if err != nil {
		fmt.Println("GetMapsFromBuildingFloor() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, location)
}

// @Summary Get floors of a Building
// @Description Get floors of a building by name
// @ID get-building-floor-name
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param building path string false "name of the building"
// @Success 200 {array} string
// @Router /buildings/{building}/floors [get]
func (b BuildingController) GetFloorsFromBuilding(context *gin.Context) {
	buildingStr := context.Param("building")

	building, err := b.buildingProvider.GetBuilding(buildingStr)
	if err != nil {
		fmt.Println("GetFloorsFromBuilding() failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

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
