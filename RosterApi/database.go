package main

import (
	"errors"
	"fmt"
	"database/sql"
	"strconv"
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

func getPlayer(jerseyNumber string) (Player, error) {
	var firstName string
	var lastName string
	var position string
	
	database := getDatabase()
	
	var query string = "SELECT first_name, last_name, position FROM public.roster WHERE jersey_number = $1"
	
	row := database.QueryRow(query, jerseyNumber)
	err := row.Scan(&firstName, &lastName, &position)
	if err != nil {
		if err == sql.ErrNoRows{
			return Player {}, errors.New("Player Not Found")
		} else {
			panic(err)
		}
	}
	
	jerseyNumberInt, err := strconv.Atoi(jerseyNumber)
	errorHandler(err)
	
	player := Player {
		JerseyNumber: jerseyNumberInt,
		FirstName: firstName,
		LastName: lastName,
		Position: position,
	}
	
	return player, nil
}

func getAllPlayers () []Player {
	database := getDatabase()
	var query string = "SELECT jersey_number, first_name, last_name, position FROM public.roster ORDER BY jersey_number"
	
	result, err := database.Query(query)
	errorHandler(err)
	
	var outputArray []Player
	for result.Next() {
		var jerseyNumber int
		var firstName string
		var lastName string
		var position string
	
		result.Scan(&jerseyNumber, &firstName, &lastName, &position)
		
		player := Player {
			JerseyNumber: jerseyNumber,
			FirstName: firstName,
			LastName: lastName,
			Position: position,
		}
		
		outputArray = append(outputArray, player)
		
	}
	return outputArray
}

func addPlayer (player Player) int64 {
	database := getDatabase()
	sql, err := database.Prepare("INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES ($1, $2, $3, $4)")
	errorHandler(err)
	result, err := sql.Exec(player.JerseyNumber, player.FirstName, player.LastName, player.Position)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}

func connectionTest() string {
	var version string
	
	database := getDatabase()
	
	var query string = "SELECT version()"
	
	row := database.QueryRow(query)
	err := row.Scan(&version)
	errorHandler(err)
	
	return fmt.Sprintf("Connected: %s", version)
}

func deletePlayer(jerseyNumber string) int64 {
	database := getDatabase()
	sql, err := database.Prepare("DELETE FROM public.roster WHERE jersey_number = $1")
	errorHandler(err)
	result, err := sql.Exec(jerseyNumber)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}

// Allow the player's position to be updated

func updatePlayer(player Player) int64 {
	database := getDatabase()
	sql, err := database.Prepare("UPDATE public.roster SET position = $1 WHERE jersey_number = $2")
	errorHandler(err)
	result, err := sql.Exec(player.Position, player.JerseyNumber)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}