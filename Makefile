_build:
	go build -o wip main.go
	chmod +x wip

lint:
	clear
	golanci-lint run

test:
	go test -v ./...
