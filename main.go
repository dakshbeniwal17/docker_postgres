package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	host, lookup := os.LookupEnv("DB_HOST")
	if !lookup {
		panic("DB_HOST is missing")
	}
	port, lookup := os.LookupEnv("DB_PORT")
	if !lookup {
		panic("DB_PORT is missing")
	}
	user, lookup := os.LookupEnv("DB_USER")
	if !lookup {
		panic("DB_USER is missing")
	}
	password, lookup := os.LookupEnv("DB_PASSWORD")
	if !lookup {
		panic("DB_PASSWORD is missing")
	}
	dbName, lookup := os.LookupEnv("DB_NAME")
	if !lookup {
		panic("DB_NAME is missing")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("[ERROR] opening postgres connection: %+v\n", err))
	} else {
		fmt.Printf("Connected to database: %s\n", dbName)
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("[ERROR] pinging db: %v\n", err))
	} else {
		fmt.Println("Database ping-pong successful")
	}

	if err := db.Close(); err != nil {
		panic(fmt.Sprintf("[ERROR] closing db connection: %v\n", err))
	} else {
		fmt.Println("Database connection closed successfully")
	}
}
