package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"	// blank identifier is used to import the package solely for its side-effects (driver initialization)
)

func main() {
	// PostgreSQL connection string
	dbuser := "appuser"
	dbpass := "learningpass"
	dbhost := "localhost"
	dbport := 5432
	dbname := "learningdb"
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbuser, dbpass, dbhost, dbport, dbname)

	// Open a database handle and a connection pool (connection to the PostgreSQL database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when main() exits

	// Verify the connection to the database (ping the database)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the PostgreSQL database!")
	// You can now perform database operations using the 'db' object.
	// query the table
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("Failed to query users table:", err)
	}
	defer rows.Close() // Close the rows iterator when main() exits

	// iterate through the result set
	for rows.Next() { 
		// next() returns true if there is a next row in the result set, 
		// and false if there are no more rows or if an error occurred
		var id int
		var name string
		var age int
		// scan the row into variables
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Printf("User ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}

}
