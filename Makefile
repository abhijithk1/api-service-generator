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
	go build -o api-service-generator

.PHONY: api-service-generator, clean, test, install, build