FROM golang:latest

workdir /app

copy go.mod go.sum ./

run go mod download

copy . . 

run go build -o main ./cmd/

expose 8080

CMD ["./main"]