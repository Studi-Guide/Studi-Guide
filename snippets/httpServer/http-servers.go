package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	jsonSerialized, _ := getJson()

	if jsonSerialized == "" {
		jsonSerialized = "an error occurred"
	}

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(1 * time.Second):
		fmt.Fprintf(w, jsonSerialized) // "<head><title>STGD</title></head><body><h1>MAP</h1></body>"
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func getJson() (string, error) {
	dummyData, err := ioutil.ReadFile("dummy-data.json")
	return string(dummyData), err
}

func main() {
	http.HandleFunc("/api/KA.3", hello)
	http.ListenAndServe(":8090", nil)
}
