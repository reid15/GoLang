// Handle configuration file data

package main

import (
	"encoding/json"
	"fmt"
    "io/ioutil"
)

type DatabaseConfiguration struct {
	Host string `json:"db_host"`
	Port int `json:"db_port"`
	DB_Name string `json:"db_name"`
	UserName string `json:"db_userName"`
	Password string `json:"db_password"`
}

type DatabaseConfig struct {
	DatabaseConfiguration
}

type ServiceConfiguration struct {
	Port int `json:"port"`
}
	
type ServiceConfig struct {
	ServiceConfiguration
}

type Configuration struct {
	ServiceConfig ServiceConfig
	DatabaseConfig DatabaseConfig
}

func getConfiguration() Configuration {
	var config Configuration
	fmt.Println("Reading Configuration File")
	data, err := ioutil.ReadFile("roster_api.config")
	errorHandler(err)
	
	err = json.Unmarshal(data, &config)
	errorHandler(err)
	
	return config
}