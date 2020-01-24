package roomcontroller

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	router   *mux.Router
	roomList []Room
}

func (l *RoomController) Initialize(router *mux.Router) {

	l.roomList = append(l.roomList, Room{Name: "RoomN01", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN02", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN03", Description: "Dummy"})
	l.roomList = append(l.roomList, Room{Name: "RoomN04", Description: "Dummy"})
	l.router = router

	l.router.HandleFunc("/", l.GetRoomList).Methods("GET")
}

func (l *RoomController) Run(addr string) {
	http.ListenAndServe(addr, l.router)
}

func (l *RoomController) GetRoomList(w http.ResponseWriter, r *http.Request) {
	log.Print("[RoomController] Request RoomList received")
	json.NewEncoder(w).Encode(l.roomList)
}

func (l *RoomController) GetItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	for _, item := range l.roomList {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func (l *RoomController) AddItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Room
	json.Unmarshal(reqBody, &item)

	l.roomList = append(l.roomList, item)
}

func (l *RoomController) RemoveItem(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	for index, item := range l.roomList {
		if item.Name == name {
			l.roomList = append(l.roomList[:index], l.roomList[index+1:]...)
		}
	}
}
