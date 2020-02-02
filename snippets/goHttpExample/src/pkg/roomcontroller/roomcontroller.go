package roomcontroller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"image"
	"io/ioutil"
	"log"
	"net/http"
)

type Room struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Coordinates image.Rectangle `json:"coordinates"`
}

type RoomController struct {
	router   *gin.RouterGroup
	roomList []Room
}

func (l *RoomController) Initialize(router *gin.RouterGroup) {

	l.roomList = append(l.roomList, Room{Name: "RoomN01", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN02", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN03", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN04", Description: "Dummy"})
	l.router = router

	l.router.GET("/", l.GetRoomList)
}

// GetRoomList godoc
// @Summary Get Room List
// @Description Gets all available rooms
// @ID get-room-list
// @Accept  json
// @Tags RoomController
// @Produce  json
// @Success 200 {array} Room
// @Router /roomlist/ [get]
func (l *RoomController) GetRoomList(c *gin.Context) {
	log.Print("[RoomController] Request RoomList received")
	c.JSON(http.StatusOK, gin.H{
		"code" : http.StatusOK,
		"message": l.roomList,// cast it to string before showing
	})
}

func (l *RoomController) GetItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(r)["name"]

	for _, item := range l.roomList {
		if item.Name == name {
			c.JSON(http.StatusOK ,item)
		}
	}
}

func (l *RoomController) AddItem(c *gin.Context) {
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	var item Room
	json.Unmarshal(reqBody, &item)

	l.roomList = append(l.roomList, item)
}

func (l *RoomController) RemoveItem(c *gin.Context) {
	name := c.Param("name") //mux.Vars(r)["name"]
	for index, item := range l.roomList {
		if item.Name == name {
			l.roomList = append(l.roomList[:index], l.roomList[index+1:]...)
		}
	}
}
