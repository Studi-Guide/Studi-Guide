package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studi-guide/pkg/roomcontroller/models"
)

type RoomController struct {
	router   *gin.RouterGroup
	provider models.RoomServiceProvider
}

func NewRoomController(router *gin.RouterGroup, provider models.RoomServiceProvider) error {
	r := RoomController{router: router, provider: provider}
	r.router.GET("/", r.GetRoomList)
	r.router.GET("/room/:name", r.GetRoom)
	r.router.GET("/floor/:floor", r.GetRoomListFromFloor)
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
// @Param floor query int false "floor of the room"
// @Param alias query string false "potential alias of the room"
// @Param room query string false "rooms that contain the query string in name, alias or description"
// @Success 200 {array} models.Room
// @Router /roomlist/ [get]
func (l *RoomController) GetRoomList(c *gin.Context) {

	nameFilter := c.Query("name")
	floorFilter := c.Query("floor")
	aliasFilter := c.Query("alias")
	roomFilter := c.Query("room")

	var rooms []models.Room
	var err error

	if len(nameFilter) == 0 && len(floorFilter) == 0 && len(aliasFilter) == 0 && len(roomFilter) == 0 {
		rooms, err = l.provider.GetAllRooms()
	} else {
		rooms, err = l.provider.FilterRooms(floorFilter, nameFilter, aliasFilter, roomFilter)
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
	return

	fmt.Println(len(c.Request.URL.Query()))
	fmt.Println(c.Request.URL.Query())
	fmt.Println("c.Query(\"room\"): ", c.Query("room"))


}

// GetRoom godoc
// @Summary Get Room by Name
// @Description Gets a specify room by its unique name
// @ID get-room
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Param name path string true "get room by name"
// @Success 200 {object} models.Room
// @Router /roomlist/room/{name} [get]
func (l *RoomController) GetRoom(c *gin.Context) {
	//name := c.Query("name") //mux.Vars(r)["name"]
	name := c.Param("name")

	room, err := l.provider.GetRoom(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, room)
	}
}

// GetRoomListFromFloor godoc
// @Summary Get Room List From Floor
// @Description Gets all available rooms for a certain floor
// @ID get-room-list-floor
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Param floor path int true "filter rooms by floor"
// @Success 200 {array} models.Room
// @Router /roomlist/floor/{floor} [get]
func (l *RoomController) GetRoomListFromFloor(c *gin.Context) {
	floor := c.Param("floor")

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
		fmt.Println("GetRoomListFromFloor() failed with error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		fmt.Println(rooms)
		c.JSON(http.StatusOK, rooms)
	}
}