package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	
	// Map API calls to a request handler function
	router.HandleFunc("/players/", getAllPlayersRequest).Methods("GET")
	router.HandleFunc("/players/{jerseyNumber}", getPlayerRequest).Methods("GET")
	
	fmt.Println("Starting Server")
	
	errorCode := http.ListenAndServe(":2222", router)
	if errorCode != nil {
		panic(errorCode)
	}
	
}

// Request Handlers

func getAllPlayersRequest(rw http.ResponseWriter, r *http.Request) {
	players := getAllPlayers()
	fmt.Fprintf(rw, players)
}

func getPlayerRequest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    player := getPlayer(vars["jerseyNumber"])
	fmt.Fprintf(rw, player)
}

// Database Functions

func getPlayer(jerseyNumber string) string {
	var firstName string
	var lastName string
	var position string
	
	database := getDatabase()
	
	var query string = "SELECT first_name, last_name, position FROM public.roster WHERE jersey_number = $1"
	
	row := database.QueryRow(query, jerseyNumber)
	errorCode := row.Scan(&firstName, &lastName, &position)
	if errorCode != nil {
		if errorCode == sql.ErrNoRows{
			panic("Player Not Found")
		} else {
			panic(errorCode)
		}
	}
	
	return fmt.Sprintf("Player %s %s Is A %s", firstName, lastName, position)
}

func getAllPlayers () string {
	database := getDatabase()
	var query string = "SELECT jersey_number, first_name, last_name, position FROM public.roster ORDER BY jersey_number"
	
	result, errorCode := database.Query(query)
	if errorCode != nil {
		panic(errorCode)
	}
	
	var outputString string
	for result.Next() {
		var jerseyNumber string
		var firstName string
		var lastName string
		var position string
	
		result.Scan(&jerseyNumber, &firstName, &lastName, &position)
		outputString += fmt.Sprintf("Player #%s %s %s Is A %s \n", jerseyNumber, firstName, lastName, position)
	}
	return outputString
}

func getDatabase() *sql.DB {
	var host string = "localhost"
	var port string = "5432"
	var databaseName string = "test_db"
	var userName string = "test_user2"
	var pwd string = "password1"
	
	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userName, pwd, databaseName)
	database, _ := sql.Open("postgres", connectionString)
	
	return database
}