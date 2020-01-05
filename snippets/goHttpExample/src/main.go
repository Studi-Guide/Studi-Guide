package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

type HelloWorldResponse struct {
	Message       string `json:"msg"`
	ProvidedName  string `json:"name"`
}

func HelloWorldGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		name := r.FormValue("name")

		response := HelloWorldResponse{
			ProvidedName: name,
			Message: fmt.Sprintf("Hello, %s!", name),
		}

		JSONResponse(w, 200, response)
	})
}

func HelloWorldHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		response := HelloWorldResponse{
			ProvidedName: vars["name"],
			Message: fmt.Sprintf("Hello, %s!", vars["name"]),
		}

		JSONResponse(w, 200, response)
	})
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Convert our interface to JSON
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json")
	// Our response code
	w.WriteHeader(code)

	w.Write(response)
}

func (a *App) Initialize() {
	router := mux.NewRouter()
	router.Handle("/hello", HelloWorldGetHandler() ).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	a.Router = router
	
	router.Handle("/hello/{name}", HelloWorldHandler())

	router.HandleFunc("/app/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
}

func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}

func main() {
	a := ShoppingList{}
	a.Initialize()
	a.Run(":8080")
}