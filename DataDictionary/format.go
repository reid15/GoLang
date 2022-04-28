// Build HTML for output

package main

import (
	"strings"
	"time"
)

func formatHeader(databaseName string, schemaName string) string {
	builder := strings.Builder{}

	builder.WriteString("<html>")
	builder.WriteString("<head>")
	builder.WriteString("<title>")
	builder.WriteString("Data Dictionary")
	builder.WriteString("</title>")
	builder.WriteString("</head>")
	builder.WriteString("<body>")
	builder.WriteString("Data Dictionary<br>")
	builder.WriteString("Database: " + databaseName + "<br>")
	builder.WriteString("Schema: " + schemaName + "<br>")
	builder.WriteString(time.Now().Format("January 2, 2006") + "<br>")
	builder.WriteString("<br>")

	return builder.String()
}

func formatTableList(tables []Table) string {
	builder := strings.Builder{}

	for _, element := range tables {
		builder.WriteString("<a href='#" + element.TableName + "'>" + element.TableName + "</a><br>")
	}

	return builder.String()
}

func formatTable(table Table) string {
	builder := strings.Builder{}

	builder.WriteString("<h3 id='" + table.TableName + "'>Table:" + table.TableName + "</h3>")
	builder.WriteString("Table Description:" + table.Description + "<br><br>")

	return builder.String()
}

func formatColumns(columns []Column) string {
	builder := strings.Builder{}

	builder.WriteString("<table border='1'>")
	builder.WriteString("<tr>")
	builder.WriteString("<th>Name</th>")
	builder.WriteString("<th>Is PK</th>")
	builder.WriteString("<th>Is Nullable</th>")
	builder.WriteString("<th>Data Type</th>")
	builder.WriteString("<th>Description</th>")
	builder.WriteString("</tr>")

	for _, element := range columns {
		builder.WriteString(formatColumn(element))
	}

	builder.WriteString("</table>")
	builder.WriteString("<br>")

	return builder.String()
}

func formatColumn(column Column) string {
	builder := strings.Builder{}

	builder.WriteString("<tr>")
	builder.WriteString("<td>" + column.ColumnName + "</td>")
	builder.WriteString("<td>" + column.IsNullable + "</td>")
	builder.WriteString("<td>" + column.IsPrimaryKey + "</td>")
	builder.WriteString("<td>" + column.DataType + "</td>")
	builder.WriteString("<td>" + column.Description + "</td>")
	builder.WriteString("</tr>")

	return builder.String()
}

func formatFooter() string {
	return "</body>" + "</html>"
}
