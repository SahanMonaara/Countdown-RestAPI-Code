package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Time Struct (Model)
type Time struct {
	ID   string `json:"id"`
	TIME string `json:"time"`
}

var times []Time

//Get all the times
func getTimes(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(times)
}

//Create Time
func postTime(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var time Time
	_ = json.NewDecoder(request.Body).Decode(&time)
	times = append(times, time)
	json.NewEncoder(w).Encode(time)
}

//Update the Time
func updateTime(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range times {
		if item.ID == params["id"] {
			times = append(times[:index], times[index+1:]...)
			var time Time
			_ = json.NewDecoder(request.Body).Decode(&time)
			time.ID = params["id"]
			times = append(times, time)
			json.NewEncoder(w).Encode(time)
			return
		}
	}
	json.NewEncoder(w).Encode(times)
}

func getTime(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range times {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Time{})
}

func main() {
	//initial router
	router := mux.NewRouter()

	times = append(times, Time{ID: "1", TIME: "15"})

	//Router Handlers
	router.HandleFunc("/api/times", getTimes).Methods("GET")
	router.HandleFunc("/api/times", postTime).Methods("POST")
	router.HandleFunc("/api/times/{id}", updateTime).Methods("PUT")
	router.HandleFunc("/api/times/{id}", getTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
