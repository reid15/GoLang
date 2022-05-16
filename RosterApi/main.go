package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// Data Structures

type Player struct {
	JerseyNumber int `json:"jerseyNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Position string `json:"position"`
}

type ReturnMessage struct {
	Message string `json:"message"`
}

type DatabaseConfig struct {
	Host string `json:"db_host"`
	Port int `json:"db_port"`
	DB_Name string `json:"db_name"`
	UserName string `json:"db_userName"`
	Password string `json:"db_password"`
}

// Service Entry Point

func main() {
	router := mux.NewRouter()
	db := getDatabase()
	
	// Map API calls to a request handler function

	router.HandleFunc("/players/connectionTest", func(w http.ResponseWriter, r *http.Request) { connectionTestRequest(w, r, db)}).Methods("GET")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { getAllPlayersRequest(w, r, db)}).Methods("GET")
	router.HandleFunc("/players/{jerseyNumber}", func(w http.ResponseWriter, r *http.Request) { getPlayerRequest(w, r, db)}).Methods("GET")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { addPlayerRequest(w, r, db)}).Methods("POST")
	router.HandleFunc("/players/{jerseyNumber}", func(w http.ResponseWriter, r *http.Request) { deletePlayerRequest(w, r, db)}).Methods("DELETE")
	router.HandleFunc("/players/", func(w http.ResponseWriter, r *http.Request) { updatePlayerRequest(w, r, db)}).Methods("PATCH")
	
	fmt.Println("Starting Server")
	
	err := http.ListenAndServe(":2222", router)
	errorHandler(err)
}

// Utility Functions

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
