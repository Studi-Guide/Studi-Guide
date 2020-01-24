package cmd

import (
	"github.com/gorilla/mux"
	"httpExample/pkg/roomcontroller"
	"httpExample/pkg/shoppinglist"
	"log"
	"net/http"
)

func Main() error {
	log.Print("Starting initializing main controllers ...")
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Print("Creating shopping list")
	a := shoppinglist.ShoppingListApp{}
	a.Initialize(
		router.PathPrefix("/shoppinglist").Subrouter(),
		"/shoppinglist",
		"pkg/shoppinglist")

	log.Print("Creating room controller")
	roomController := roomcontroller.RoomController{}
	roomController.Initialize(router.PathPrefix("/roomlist").Subrouter())
	//a.Run(":8080")

	port := ":8080"
	log.Printf("Starting http listener on %s", port)
	log.Fatal(http.ListenAndServe(port, router))

	return nil

}
