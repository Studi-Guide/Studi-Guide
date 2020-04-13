package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func handleFloorReq(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	jsonSerialized, _ := getJson("dummy-data.json")

	if jsonSerialized == "" {
		jsonSerialized = "an error occurred"
	}

	ctx := req.Context()
	fmt.Println("server: handleFloorReq started")
	defer fmt.Println("server: handleFloorReq ended")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintf(w, jsonSerialized) // "<head><title>STGD</title></head><body><h1>MAP</h1></body>"
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func handleRouteReq(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data, _ := getJson("dummy-route.json")

	if data == "" {
		data = "an error occurred"
	}

	ctx := req.Context()
	fmt.Println("server: handleRouteReq started")
	defer fmt.Println("server: handleRouteReq ended")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintf(w, data) // "<head><title>STGD</title></head><body><h1>MAP</h1></body>"
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func getJson(fileName string) (string, error) {
	dummyData, err := ioutil.ReadFile(fileName)
	return string(dummyData), err
}

func main() {
	http.HandleFunc("/roomlist/floor/0", handleFloorReq)
	http.HandleFunc("/navigation/dir/startroom/KA.308/endroom/KA.313", handleRouteReq)
	http.ListenAndServe(":8090", nil)
}
