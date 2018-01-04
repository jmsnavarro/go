package main

//Create Go apps using SQL Server on Windows
//https://www.microsoft.com/en-us/sql-server/developer-get-started/go/windows/

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Replace with your own connection parameters
var server = "centos.localsandbox.com"
var port = 1433
var user = "sa"
var password = "pAssw0rd123$"
var database = "SampleDB"

var db *sql.DB

func main() {

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool:", err.Error())
	}
	fmt.Printf("Connected!\n\n")

	// Prepare test table
	PrepareTestTable()

	// Create employee
	createID, err := CreateEmployee("Jake", "United States")
	fmt.Printf("Inserted ID: %d successfully.\n\n", createID)

	// Read employees
	count, err := ReadEmployees()
	fmt.Printf("Read %d rows successfully.\n\n", count)

	// Update from database
	updateID, err := UpdateEmployee("Jake", "Poland")
	fmt.Printf("Updated row with ID: %d successfully.\n\n", updateID)

	// Delete from database
	rows, err := DeleteEmployee("Jake")
	fmt.Printf("Deleted %d rows successfully.\n\n", rows)
}

// PrepareTestTable function
func PrepareTestTable() (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		log.Fatal("What?")
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("DROP TABLE IF EXISTS dbo.Employees; " +
		"CREATE TABLE dbo.Employees ( " +
		"Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY, " +
		"Name NVARCHAR(50), " +
		"Location NVARCHAR(50)); " +
		"INSERT INTO dbo.Employees (Name, Location) VALUES " +
		"(N'Jared', N'Australia') ," +
		"(N'Nikita', N'India'), " +
		"(N'Tom', N'Germany');")

	// Execute non-query
	result, err := db.ExecContext(ctx, tsql)
	if err != nil {
		log.Fatal("Error preparing test table: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

// CreateEmployee function
func CreateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()
	var err error

	if db == nil {
		log.Fatal("What?")
	}

	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("INSERT INTO dbo.Employees (Name, Location) VALUES ('%s','%s');",
		name, location)

	// Execute non-query
	result, err := db.ExecContext(ctx, tsql)
	if err != nil {
		log.Fatal("Error inserting new row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

// ReadEmployees function
func ReadEmployees() (int, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM dbo.Employees;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal("Error reading rows: " + err.Error())
		return -1, err
	}

	defer rows.Close()

	var count int // = 0

	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		if err != nil {
			log.Fatal("Error reading rows: " + err.Error())
			return -1, err
		}

		fmt.Printf("\nID: %d, Name: %s, Location: %s", id, name, location)
		count++
	}

	return count, nil
}

// UpdateEmployee function
func UpdateEmployee(name string, location string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("UPDATE dbo.Employees SET Location = @Location WHERE Name= @Name")

	// Execute non-query with named parameters
	result, err := db.ExecContext(
		ctx,
		tsql,
		sql.Named("Location", location),
		sql.Named("Name", name))
	if err != nil {
		log.Fatal("Error updating row: " + err.Error())
		return -1, err
	}

	return result.LastInsertId()
}

// DeleteEmployee function
func DeleteEmployee(name string) (int64, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	tsql := fmt.Sprintf("DELETE FROM dbo.Employees WHERE Name=@Name;")

	// Execute non-query with named parameters
	result, err := db.ExecContext(ctx, tsql, sql.Named("Name", name))
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}

	return result.RowsAffected()
}
