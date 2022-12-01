BINARY=bin/engine

DB_HOST:=localhost
DB_PORT:=5432
DB_USER:=root
DB_PASSWORD:=password
DB_NAME:=postgres

run:
	DB_HOST=localhost DB_PORT=5432 DB_USER=root DB_PASSWORD=password DB_NAME=postgres go run main.go

.PHONY: run