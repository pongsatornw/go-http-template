APP_NAME = myapp
COMPOSE_FILE = docker-compose.yaml
BINARY_NAME = cmd

.PHONY: build test clean coverage run

## Run tests
test:
	go test -v ./...

## Run tests with coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out


build:
	go build -o bin ./...
	docker-compose -f $(COMPOSE_FILE) build


run:
	docker-compose -f $(COMPOSE_FILE) up -d
	./bin/$(BINARY_NAME)

clean:
	rm -rf bin coverage.out


lint:
	golangci-lint run ./...

swagger:
	swag init -g cmd/main.go