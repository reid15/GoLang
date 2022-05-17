package main

import (
	"errors"
	"fmt"
	"database/sql"
	"strconv"
	_ "github.com/lib/pq"
)

// Database Functions

func getDatabase(config DatabaseConfig) *sql.DB {
	port := strconv.FormatInt(int64(config.Port), 10)
	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, port, config.UserName, config.Password, config.DB_Name)
		
	database, err := sql.Open("postgres", connectionString)
	errorHandler(err)
	
	return database
}

func getPlayer(jerseyNumber string, database *sql.DB) (Player, error) {
	var firstName string
	var lastName string
	var position string
	
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

func getAllPlayers (database *sql.DB) []Player {
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

func addPlayer (player Player, database *sql.DB) int64 {
	sql, err := database.Prepare("INSERT INTO public.roster (jersey_number, first_name, last_name, position) VALUES ($1, $2, $3, $4)")
	errorHandler(err)
	result, err := sql.Exec(player.JerseyNumber, player.FirstName, player.LastName, player.Position)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}

func connectionTest(database *sql.DB) string {
	err := database.Ping()
	errorHandler(err)
	
	return fmt.Sprintf("Database Ping Successful")
}

func deletePlayer(jerseyNumber string, database *sql.DB) int64 {
	sql, err := database.Prepare("DELETE FROM public.roster WHERE jersey_number = $1")
	errorHandler(err)
	result, err := sql.Exec(jerseyNumber)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}

// Allow the player's position to be updated

func updatePlayer(player Player, database *sql.DB) int64 {
	sql, err := database.Prepare("UPDATE public.roster SET position = $1 WHERE jersey_number = $2")
	errorHandler(err)
	result, err := sql.Exec(player.Position, player.JerseyNumber)
	errorHandler(err)
	rowCount, err := result.RowsAffected()
	errorHandler(err)
	
	return rowCount
}