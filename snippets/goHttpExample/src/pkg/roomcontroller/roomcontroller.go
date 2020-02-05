package roomcontroller

import (
	"github.com/gin-gonic/gin"
	"httpExample/pkg/roomcontroller/conntrollers"
	"httpExample/pkg/roomcontroller/models"
)

type RoomControllerApp struct {
	roomController *controllers.RoomController
}

func (l *RoomControllerApp) Initialize(router *gin.RouterGroup) {

	var roomList []models.Room
	roomList = append(roomList, models.Room{Name: "RoomN01", Description: "Dummy"})
	roomList = append(roomList, models.Room{Name: "RoomN02", Description: "Dummy"})
	roomList = append(roomList, models.Room{Name: "RoomN03", Description: "Dummy"})
	roomList = append(roomList, models.Room{Name: "RoomN04", Description: "Dummy"})

	l.roomController = controllers.NewRoomController(router, roomList)
}
