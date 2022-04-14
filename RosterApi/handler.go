package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Request Handlers

func connectionTestRequest(rw http.ResponseWriter, r *http.Request) {
	connectionMessage := connectionTest()
	returnMessage(rw, connectionMessage)
}

func returnMessage(rw http.ResponseWriter, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	returnMessage := ReturnMessage { message }
	returnMessageJson, err := json.Marshal(returnMessage)
	errorHandler(err)
	rw.Write(returnMessageJson)
}

func getAllPlayersRequest(rw http.ResponseWriter, r *http.Request) {
	players := getAllPlayers()
	playerJson, err := json.Marshal(players)
	errorHandler(err)
		
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(playerJson)
}

func getPlayerRequest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    player, err := getPlayer(vars["jerseyNumber"])
	if err != nil {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusInternalServerError)
		returnMessage := ReturnMessage { err.Error() }
		returnMessageJson, err := json.Marshal(returnMessage)
		errorHandler(err)
		rw.Write(returnMessageJson)
		return
	}
	
	playerJson, err := json.Marshal(player)
	errorHandler(err)
	
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(playerJson)
}

func addPlayerRequest(rw http.ResponseWriter, r *http.Request) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	errorHandler(err)
	
	recordCount := addPlayer(player)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}

func deletePlayerRequest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    recordCount := deletePlayer(vars["jerseyNumber"])
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}

func updatePlayerRequest(rw http.ResponseWriter, r *http.Request) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	errorHandler(err)
	
	recordCount := updatePlayer(player)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	returnMessage(rw, response)
}
