run:
	go run cmd/main.go

build:
	go build -o main .

swag:
	swag init -g cmd/main.go

lint: 
	golangci-lint run
format:
	go fmt ./...