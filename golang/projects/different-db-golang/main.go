package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"     // replace with your MySQL username
	password = "password" // replace with your MySQL password
	hostname = "127.0.0.1:3306"
	dbname   = "sampledb" // name of the database to create
)

// Function to create a DSN (Data Source Name)
func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	// Connect to MySQL without specifying a database
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
	defer db.Close()

	// Create the database if it doesn't exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Fatalf("Error creating database: %s", err)
	}

	// Connect to the newly created database
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	// Create a table
	createTable := `CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100)
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Error creating table: %s", err)
	}

	// Insert a record
	insertUser := `INSERT INTO users (name, email) VALUES (?, ?);`
	_, err = db.Exec(insertUser, "John Doe", "john@example.com")
	if err != nil {
		log.Fatalf("Error inserting user: %s", err)
	}

	fmt.Println("User inserted successfully!")

	// Query the table
	rows, err := db.Query("SELECT id, name, email FROM users;")
	if err != nil {
		log.Fatalf("Error querying users: %s", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s (%s)\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
