package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

// Database Functions

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

func addPlayer (player Player) int64 {
	database := getDatabase()
	sql, err := database.Prepare("INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	result, err := sql.Exec(player.JerseyNumber, player.FirstName, player.LastName, player.Position)
	if err != nil {
		panic(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	
	return rowCount
}

func connectionTest() string {
	var version string
	
	database := getDatabase()
	
	var query string = "SELECT version()"
	
	row := database.QueryRow(query)
	errorCode := row.Scan(&version)
	if errorCode != nil {
		panic(errorCode)
	}
	
	return fmt.Sprintf("Connected: %s", version)
}

func deletePlayer(jerseyNumber string) int64 {
	database := getDatabase()
	sql, err := database.Prepare("DELETE FROM public.roster WHERE jersey_number = $1")
	if err != nil {
		panic(err)
	}
	result, err := sql.Exec(jerseyNumber)
	if err != nil {
		panic(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	
	return rowCount
}

// Allow the player's position to be updated

func updatePlayer(player Player) int64 {
	database := getDatabase()
	sql, err := database.Prepare("UPDATE public.roster SET position = $1 WHERE jersey_number = $2")
	if err != nil {
		panic(err)
	}
	result, err := sql.Exec(player.Position, player.JerseyNumber)
	if err != nil {
		panic(err)
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	
	return rowCount
}