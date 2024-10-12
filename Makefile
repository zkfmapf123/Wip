_build:
	go build -o main main.go
	chmod +x main

lint:
	clear
	golanci-lint run

test:
	go test -v ./...
