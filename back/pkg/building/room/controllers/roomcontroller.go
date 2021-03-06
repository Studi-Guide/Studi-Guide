package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"studi-guide/pkg/building/db/entitymapper"
	"studi-guide/pkg/building/room/models"
)

type RoomController struct {
	router   *gin.RouterGroup
	provider models.RoomServiceProvider
}

func NewRoomController(router *gin.RouterGroup, provider models.RoomServiceProvider) error {
	r := RoomController{router: router, provider: provider}
	r.router.GET("", r.GetRoomList)
	r.router.GET("/:room", r.GetRoom)
	//r.router.GET("/building/:building/floor/:floor", r.GetRoomListFromFloor)
	return nil
}

// GetRoomList godoc
// @Summary Get Room List
// @Description Gets all available rooms
// @ID get-room-list
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Param name query string false "room name"
// @Param building query string false "building name"
// @Param campus query string false "campus name"
// @Param floor query int false "floor of the room"
// @Param alias query string false "potential alias of the room"
// @Param room query string false "rooms that contain the query string in name, alias or description"
// @Success 200 {array} entitymapper.Room
// @Router /rooms/ [get]
func (l *RoomController) GetRoomList(c *gin.Context) {

	buildingFilter := c.Param("building")

	nameFilter := c.Query("name")
	floorFilter := c.Query("floor")
	aliasFilter := c.Query("alias")
	roomFilter := c.Query("room")
	if len(buildingFilter) == 0 {
		buildingFilter = c.Query("building")
	}
	campusFilter := c.Query("campus")

	var rooms []entitymapper.Room
	var err error

	if len(nameFilter) == 0 && len(floorFilter) == 0 && len(aliasFilter) == 0 && len(roomFilter) == 0 && len(buildingFilter) == 0 && len(campusFilter) == 0 {
		rooms, err = l.provider.GetAllRooms()
	} else {
		rooms, err = l.provider.FilterRooms(floorFilter, nameFilter, aliasFilter, roomFilter, buildingFilter, campusFilter)
	}

	if err != nil {
		fmt.Println("GetAllRomms() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		//fmt.Println(rooms)
		c.JSON(http.StatusOK, rooms)
	}
}

// GetRoom godoc
// @Summary Get Room by Name
// @Description Gets a specify room by its unique name
// @ID get-room
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Param name path string true "get room by name"
// @Success 200 {object} entitymapper.Room
// @Router /rooms/{name} [get]
func (l *RoomController) GetRoom(c *gin.Context) {
	//name := c.Query("name") //mux.Vars(r)["name"]
	room := c.Param("room")
	building := c.Param("building")
	campus := c.Param("campus")

	r, err := l.provider.GetRoom(room, building, campus)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, r)
	}
}
