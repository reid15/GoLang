// Handle configuration file data

package main

import (
	"encoding/json"
	"fmt"
    "io/ioutil"
)

var GlobalDatabaseConfig DatabaseConfig

func setDatabaseConfig() {
	if (DatabaseConfig{}) == GlobalDatabaseConfig {
		fmt.Println("Reading Config File")
		data, err := ioutil.ReadFile("roster_api.config")
		errorHandler(err)
		err = json.Unmarshal(data, &GlobalDatabaseConfig)
		errorHandler(err)
	}
}