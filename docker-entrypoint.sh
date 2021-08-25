echo "Waiting for DB to start..."
./wait-for-it.sh database:8080 

echo "Starting the Service..."
go run main.go