package main

import (
	"database/sql"
	"fmt"
	"log"
	"context"
	"time"
	"errors"
	_ "github.com/lib/pq"	// blank identifier is used to import the package solely for its side-effects (driver initialization)
)

func main() {
	// PostgreSQL connection parameters
	dbuser := "appuser"
	dbpass := "learningpass"
	dbhost := "localhost"
	dbport := 5432
	dbname := "learningdb"

	// PostgreSQL connection string
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

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(
		context.Background(), 
		5*time.Second,
	)
	// context.Background() creates a starting context with no deadline.
	// context.WithTimeout() creates a child context with a five-second deadline.
	// cancel() releases the context's resources when you're finished.

	defer cancel() // Ensure the context is canceled to free resources

	// query the table using the context we created. 
	// This allows us to set a timeout for the query operation.
	// QueryContext() works like Query(), but it also accepts a context that can signal cancellation or a deadline to the database driver.
	rows, err := db.QueryContext(ctx, "SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("Failed to query users table:", err)
	}

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("Query timed out!")
		} else if errors.Is(err, context.Canceled) {
			fmt.Println("Query was canceled.")
		} else {
			fmt.Println("Query failed:", err)
		}
		return
	}
	defer rows.Close()

	fmt.Println("Query completed successfully.")
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
