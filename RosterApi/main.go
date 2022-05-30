package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Data Structures

// Player : Player attributes
type Player struct {
	JerseyNumber int    `json:"jerseyNumber"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Position     string `json:"position"`
}

// ReturnMessage : Message to return to API user - Error or informational
type ReturnMessage struct {
	Message string `json:"message"`
}

// Service Entry Point

func main() {
	router := mux.NewRouter()
	config := getConfiguration()
	var servicePort string = strconv.FormatInt(int64(config.ServiceConfig.Port), 10)
	db := getDatabase(config.DatabaseConfig)

	// Map API calls to a request handler function

	router.HandleFunc("/players/connectionTest", func(w http.ResponseWriter, r *http.Request) { connectionTestRequest(w, r, db) }).Methods("GET")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { getAllPlayersRequest(w, r, db) }).Methods("GET")
	router.HandleFunc("/players/{jerseyNumber}", func(w http.ResponseWriter, r *http.Request) { getPlayerRequest(w, r, db) }).Methods("GET")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { addPlayerRequest(w, r, db) }).Methods("POST")
	router.HandleFunc("/players/{jerseyNumber}", func(w http.ResponseWriter, r *http.Request) { deletePlayerRequest(w, r, db) }).Methods("DELETE")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { updatePlayerRequest(w, r, db) }).Methods("PATCH")

	fmt.Println("Starting Service")
	fmt.Println("Listening On Port " + servicePort)

	err := http.ListenAndServe((":" + servicePort), router)
	errorHandler(err)
}

// Utility Functions

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
