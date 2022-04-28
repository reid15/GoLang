package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func getDatabase(
	host string,
	port string,
	databaseName string,
	userName string,
	pwd string) *sql.DB {

	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userName, pwd, databaseName)

	database, err := sql.Open("postgres", connectionString)
	errorHandler(err)

	return database
}

func getTables(database *sql.DB,
	databaseName string,
	schemaName string) []Table {

	var query string = `SELECT table_name, 
		COALESCE(OBJ_DESCRIPTION(CONCAT(table_schema, '.',  table_name)::regclass), '') as description 
	FROM information_schema.tables
	WHERE table_catalog = $1 AND table_schema = $2 AND table_type = 'BASE TABLE' 
	ORDER BY table_name;`

	result, err := database.Query(query, databaseName, schemaName)
	errorHandler(err)

	var outputArray []Table
	for result.Next() {
		var tableName string
		var description string

		err := result.Scan(&tableName, &description)
		errorHandler(err)
		
		table := Table{
			TableName:   tableName,
			Description: description,
		}

		outputArray = append(outputArray, table)

	}
	return outputArray
}

func getColumns(database *sql.DB,
	databaseName string,
	schemaName string,
	tableName string) []Column {

	var query string = `WITH cte_pk as (
		SELECT tc.table_catalog, tc.table_schema, tc.table_name, kc.column_name
		FROM information_schema.table_constraints as tc
		JOIN information_schema.key_column_usage as kc
			ON kc.table_catalog = tc.table_catalog
			AND kc.table_schema = tc.table_schema
			AND kc.table_name = tc.table_name
		WHERE tc.constraint_type = 'PRIMARY KEY'
	)
	SELECT c.column_name, 
		COALESCE(COL_DESCRIPTION(CONCAT(c.table_schema, '.', c.table_name)::regclass, ordinal_position), '') as description,
		CASE WHEN c.is_nullable = 'YES' THEN 'TRUE' ELSE 'FALSE' END as is_nullable,
		CASE WHEN c.character_maximum_length IS NOT NULL THEN CONCAT(c.udt_name, '(', c.character_maximum_length, ')')
		WHEN c.udt_name IN ('decimal', 'numeric') THEN CONCAT(c.udt_name, '(', c.numeric_precision, ',', c.numeric_scale, ')')
		ELSE c.udt_name END as data_type,
		CASE WHEN p.column_name IS NOT NULL THEN 'TRUE' ELSE 'FALSE' END as is_pk
	FROM information_schema.columns as c  
	JOIN information_schema.tables as t 
		ON t.table_catalog = c.table_catalog 
		AND t.table_schema = c.table_schema 
		AND t.table_name = c.table_name 
	LEFT JOIN cte_pk as p
		ON p.table_catalog = c.table_catalog 
		AND p.table_schema = c.table_schema 
		AND p.table_name = c.table_name
		AND p.column_name = c.column_name
	WHERE c.table_catalog = $1
		AND c.table_schema = $2
		AND c.table_name = $3
		AND t.table_type = 'BASE TABLE'
	ORDER BY c.ordinal_position;`
		
	result, err := database.Query(query, databaseName, schemaName, tableName)
	errorHandler(err)
			
	var outputArray []Column
	for result.Next() {
		var columnName string
		var description string
		var isNullable string
		var dataType string
		var isPrimaryKey string

		err := result.Scan(&columnName, &description, &isNullable, &dataType, &isPrimaryKey)
		errorHandler(err)
		
		column := Column{
			ColumnName:   columnName,
			Description:  description,
			IsNullable:   isNullable,
			IsPrimaryKey: isPrimaryKey,
			DataType:     dataType,
		}

		outputArray = append(outputArray, column)

	}
	return outputArray
}
