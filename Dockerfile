FROM golang:alpine 

COPY . /app

WORKDIR /app 

EXPOSE 8080

# RUN go run cmd/data/main.go 

# CMD go run cmd/server/main.go