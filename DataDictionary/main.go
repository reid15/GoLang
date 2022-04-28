package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Data Structures

type Table struct {
	TableName   string
	Description string
}

type Column struct {
	ColumnName   string
	Description  string
	IsNullable   string
	IsPrimaryKey string
	DataType     string
}

// Entry Point

func main() {
	var host string
	var port string
	var databaseName string
	var schemaName string
	var userName string
	var password string

	fmt.Println("Data Dictionary")
	fmt.Println("Enter host name: ")
	fmt.Scanln(&host)
	fmt.Println("Enter port number: ")
	fmt.Scanln(&port)
	fmt.Println("Enter database name: ")
	fmt.Scanln(&databaseName)
	fmt.Println("Enter schema name: ")
	fmt.Scanln(&schemaName)
	fmt.Println("Enter user name: ")
	fmt.Scanln(&userName)
	fmt.Println("Enter password: ")
	fmt.Scanln(&password)

	fmt.Println("Query For Tables")

	database := getDatabase(host, port, databaseName, userName, password)

	tables := getTables(database, databaseName, schemaName)

	builder := strings.Builder{}

	builder.WriteString(formatHeader(databaseName, schemaName))
	builder.WriteString(formatTableList(tables))

	for _, element := range tables {
		fmt.Println(element.TableName)
		builder.WriteString(formatTable(element))
		columns := getColumns(database, databaseName, schemaName, element.TableName)
		builder.WriteString(formatColumns(columns))
	}

	builder.WriteString(formatFooter())

	// Write to file
	message := []byte(builder.String())
	// 0666 = Read/Write for all users
	err := ioutil.WriteFile("DataDictionary.html", message, 0666)
	errorHandler(err)
	
	var a string
	fmt.Println("Completed")
	fmt.Scanln(&a)
}

// Utility Functions

func errorHandler(err error) {
	if err != nil {
		var a string
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Scanln(&a)
		panic(err)
	}
}
