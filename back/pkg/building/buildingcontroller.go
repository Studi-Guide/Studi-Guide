package building

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuildingController struct {
	router   *gin.RouterGroup
	provider BuildingProvider
}

func NewBuildingController(router *gin.RouterGroup, provider BuildingProvider) error {
	b:= BuildingController{
		router:   router,
		provider: provider,
	}

	b.router.GET("", b.GetBuildings)
	b.router.GET("/:name", b.GetBuildingByName)

	return nil
}


// GetLocations godoc
// @Summary Get All Buildings
// @Description Gets buildings by possible filters
// @Accept  json
// @Produce  json
// @Tags BuildingController
// @Param name query string false "name of the building"
// @Success 200 {array} building.Building
// @Router /buildings [get]
func (b BuildingController) GetBuildings(c *gin.Context) {
	name := c.Query("name")

	var buildings []Building
	var err error

	var useFilterApi bool

	if len(name) == 0 {
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		buildings, err = b.provider.FilterBuildings(name)
	} else {
		buildings, err = b.provider.GetAllBuildings()
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
// @Param name path string false "name of the building"
// @Success 200 {array} building.Building
// @Router /buildings/{name} [get]
func (b BuildingController) GetBuildingByName(c *gin.Context) {
	name := c.Param("name")

	building, err := b.provider.GetBuilding(name)
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
