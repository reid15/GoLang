package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Request Handlers

func connectionTestRequest(rw http.ResponseWriter, r *http.Request) {
	connection := connectionTest()
	fmt.Fprintf(rw, connection)
}

func getAllPlayersRequest(rw http.ResponseWriter, r *http.Request) {
	players := getAllPlayers()
	fmt.Fprintf(rw, players)
}

func getPlayerRequest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    player := getPlayer(vars["jerseyNumber"])
	fmt.Fprintf(rw, player)
}

func addPlayerRequest(rw http.ResponseWriter, r *http.Request) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	if err != nil {
		panic(err)
	}
	
	recordCount := addPlayer(player)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	fmt.Fprintf(rw, response)
}

func deletePlayerRequest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    recordCount := deletePlayer(vars["jerseyNumber"])
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	fmt.Fprintf(rw, response)
}

func updatePlayerRequest(rw http.ResponseWriter, r *http.Request) {
	var player Player
	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&player)
	if err != nil {
		panic(err)
	}
	
	recordCount := updatePlayer(player)
	response := strconv.FormatInt(recordCount, 10) + " Record(s) Affected"
	fmt.Fprintf(rw, response)
}
