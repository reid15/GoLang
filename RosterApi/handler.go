// Handle services calls - format data for output

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Request Handlers

func connectionTestRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	connectionMessage := connectionTest(db)
	returnMessage(rw, connectionMessage)
}

func getAllPlayersRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	players := getAllPlayers(db)
	playerJson, err := json.Marshal(players)
	errorHandler(err)
		
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(playerJson)
}

func getPlayerRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	err := isValidJerseyNumberString(vars["jerseyNumber"])
	if err != nil {
		writeErrorMessage(rw, err, 400)
		return
	}
	
    player, err := getPlayer(vars["jerseyNumber"], db)
	if err != nil {
		writeErrorMessage(rw, err, 500)
		return
	}
	
	playerJson, err := json.Marshal(player)
	errorHandler(err)
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(playerJson)
}

func addPlayerRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	errorHandler(err)
	
	err = isValidPlayer(player)
	if err != nil {
		writeErrorMessage(rw, err, 400)
		return
	}
	
	recordCount := addPlayer(player, db)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}

func deletePlayerRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	err := isValidJerseyNumberString(vars["jerseyNumber"])
	if err != nil {
		writeErrorMessage(rw, err, 400)
		return
	}
    recordCount := deletePlayer(vars["jerseyNumber"], db)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}

func updatePlayerRequest(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	errorHandler(err)
	
	err = isValidPlayerForUpdate(player)
	if err != nil {
		writeErrorMessage(rw, err, 400)
		return
	}
	
	recordCount := updatePlayer(player, db)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}

// Format Return Messages

func returnMessage(rw http.ResponseWriter, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	returnMessage := ReturnMessage { message }
	returnMessageJson, err := json.Marshal(returnMessage)
	errorHandler(err)
	rw.Write(returnMessageJson)
}

func writeErrorMessage(rw http.ResponseWriter, err error, httpStatusCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatusCode)
	returnMessage := ReturnMessage { err.Error() }
	returnMessageJson, err := json.Marshal(returnMessage)
	errorHandler(err)
	rw.Write(returnMessageJson)
}
