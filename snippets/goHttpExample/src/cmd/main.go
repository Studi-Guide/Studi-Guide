package cmd

import (
	"github.com/gorilla/mux"
	"httpExample/pkg/roomcontroller"
	"httpExample/pkg/shoppinglist"
	"net/http"
)

func Main() (error) {

	router := mux.NewRouter()
	a := shoppinglist.ShoppingList{}
	a.Initialize(router)
	roomController := roomcontroller.RoomController{}
	roomController.Initialize(router)
	//a.Run(":8080")

	http.ListenAndServe(":8080", router)

	return nil

}