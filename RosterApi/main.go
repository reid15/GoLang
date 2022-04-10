package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Player Data Structure

type Player struct {
	JerseyNumber int `json:"jerseyNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Position string `json:"position"`
}

// Service Entry Point

func main() {
	router := mux.NewRouter()
	
	// Map API calls to a request handler function
	
	router.HandleFunc("/players/connectionTest", connectionTestRequest).Methods("GET")
	router.HandleFunc("/players/", getAllPlayersRequest).Methods("GET")
	router.HandleFunc("/players/{jerseyNumber}", getPlayerRequest).Methods("GET")
	router.HandleFunc("/players/", addPlayerRequest).Methods("POST")
	router.HandleFunc("/players/{jerseyNumber}", deletePlayerRequest).Methods("DELETE")
	router.HandleFunc("/players/", updatePlayerRequest).Methods("PATCH")
	
	fmt.Println("Starting Server")
	
	errorCode := http.ListenAndServe(":2222", router)
	if errorCode != nil {
		panic(errorCode)
	}
	
}
