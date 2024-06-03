api-service-generator:
	go build -o api-service-generator .

clean: 
	rm api-service-generator

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

install:
	go install

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

.PHONY: api-service-generator, clean, test, install, build