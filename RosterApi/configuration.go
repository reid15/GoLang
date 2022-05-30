// Handle configuration file data

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// DatabaseConfiguration stores data from configuration file for database access
type DatabaseConfiguration struct {
	Host     string `json:"db_host"`
	Port     int    `json:"db_port"`
	DBName  string `json:"db_name"`
	UserName string `json:"db_userName"`
	Password string `json:"db_password"`
}

// DatabaseConfig : Database section of configuration file
type DatabaseConfig struct {
	DatabaseConfiguration
}

// ServiceConfiguration stores data from configuration file for service
type ServiceConfiguration struct {
	Port int `json:"port"`
}

// ServiceConfig : Service section of configuration file
type ServiceConfig struct {
	ServiceConfiguration
}

// Configuration : Data from configuration file
type Configuration struct {
	ServiceConfig  ServiceConfig
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
