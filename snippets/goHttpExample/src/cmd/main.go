package cmd

import (
	"github.com/gorilla/mux"
	"httpExample/pkg/roomcontroller"
	"httpExample/pkg/shoppinglist"
	"log"
	"net/http"
)

func Main() (error) {
	log.Print("Starting initializing main controllers ...")
	router := mux.NewRouter()

	log.Print("Creating shopping list")
	a := shoppinglist.ShoppingList{}
	a.Initialize(router)

	log.Print("Creating room controller")
	roomController := roomcontroller.RoomController{}
	roomController.Initialize(router)
	//a.Run(":8080")

	port := ":8080"
	log.Printf("Starting http listener on %s", port)
	http.ListenAndServe(port, router)

	return nil

}