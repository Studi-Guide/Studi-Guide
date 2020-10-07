package campus

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/building/db/ent"
)

type CampusController struct {
	router         *gin.RouterGroup
	campusProvider CampusProvider
}

func NewCampusController(router *gin.RouterGroup, campusProvider CampusProvider) error {
	b := CampusController{
		router:         router,
		campusProvider: campusProvider,
	}

	b.router.GET("", b.GetCampus)
	b.router.GET("/:campus", b.GetCampusByName)
	return nil
}

// GetCampus godoc
// @Summary Get Campus by name
// @ID get-campus
// @Description Gets a list of campus filtered by name
// @Accept  json
// @Produce  json
// @Tags CampusController
// @Param name query string false "name of the campus"
// @Success 200 {array} ent.Campus
// @Router /campus [get]
func (c CampusController) GetCampus(context *gin.Context) {
	name := context.Query("name")

	var campusArray []ent.Campus
	var err error

	var useFilterApi bool

	if len(name) == 0 {
		useFilterApi = false
	} else {
		useFilterApi = true
	}

	if useFilterApi {
		campusArray, err = c.campusProvider.FilterCampus(name)
	} else {
		campusArray, err = c.campusProvider.GetAllCampus()
	}

	if err != nil {
		fmt.Println("GetCampus failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})

		return
	}

	context.JSON(http.StatusOK, campusArray)
}

// GetCampusByName godoc
// @Summary Get Campus by a certain name
// @Description Get one campus by name
// @ID get-campus-name
// @Accept  json
// @Produce  json
// @Tags CampusController
// @Param campus path string true "name of the campus"
// @Success 200 {object} ent.Campus
// @Router /campus/{campus} [get]
func (c CampusController) GetCampusByName(context *gin.Context) {
	name := context.Param("campus")

	campus, err := c.campusProvider.GetCampus(name)
	if err != nil {
		fmt.Println("GetCampusbyName failed with error", err)
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, campus)
}
