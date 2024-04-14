package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // Import, but don't use directly (blank import)
)

func main() {
    var (
        host     = "localhost"
        port     = "5432"
        user     = "postgres"
        password = ""
        dbname   = ""
    )

    // Connect to the database
    connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", connString)
    if err != nil {
        panic(err)
    }

    defer func() {
        // Close the database connection when the function exits
        if err := db.Close(); err != nil {
            panic(err)
        }
    }()

    // Test the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to database!")

		rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err)
    }

    defer rows.Close()

    // interate through rows
		for rows.Next() {
			var id int
			var username, email string
			err := rows.Scan(&id, &username, &email)
			if err != nil {
					panic(err)
			}
			fmt.Println("ID:", id, "Username:", username, "Email:", email)
	    }
}
