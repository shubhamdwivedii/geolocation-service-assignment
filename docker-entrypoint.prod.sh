#!/bin/sh

echo "Waiting for DB to start..."

cd geolocation

./wait-for database:8080 -- echo "Database Has Started..."
# https://github.com/eficode/wait-for

# Make sure all files are copied onto production container 

echo "Preparing Database..."
./populate 

echo "Running Server..."
./server 

# Run this by "docker-compose -f docker-compose.prod.yml up"
# "docker-compose -f docker-compose.prod.yml up -d" for detached mode. 