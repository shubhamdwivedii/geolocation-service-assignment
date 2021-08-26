hello:
	echo "Hello Makefile"

build: 
	go build -o bin/data ./cmd/data/main.go 
	go build -o bin/server ./cmd/server/main.go