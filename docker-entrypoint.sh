#!/bin/sh

echo "Waiting for DB to start..."
./wait-for database:8080 -- echo "Database Has Started..."
# https://github.com/eficode/wait-for

echo "Preparing Database..."
go run cmd/data/main.go

echo "Running Server..."
go run cmd/server/main.go