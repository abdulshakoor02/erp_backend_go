.PHONY: build run

build:
	go mod download
	go build -o app

run: build
	        ./app
