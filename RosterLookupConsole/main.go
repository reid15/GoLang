package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting Roster Lookup")
	fmt.Println("Enter Player Number Or ALL: ")
	
	var output string
	var consoleInput string
	fmt.Scanln(&consoleInput)
	if consoleInput == "ALL" {
		output = getAllPlayers ()
	} else {
		fmt.Printf("Jersey Number = %s", consoleInput)
		fmt.Println()
		output = getPlayer(consoleInput)
	}
	
	fmt.Println(output)
}

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
	var userName string = "test_user"
	var pwd string = "password1"
	
	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userName, pwd, databaseName)
	database, _ := sql.Open("postgres", connectionString)
	
	return database
}