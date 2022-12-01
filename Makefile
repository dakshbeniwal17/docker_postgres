BINARY=bin/engine

DB_HOST:=localhost
DB_PORT:=5432
DB_USER:=root
DB_PASSWORD:=password
DB_NAME:=test

run:
	DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} DB_USER=${DB_USER} DB_PASSWORD=${DB_PASSWORD} DB_NAME=${DB_NAME} go run main.go

.PHONY: run