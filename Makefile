SHELL=/bin/bash
.PHONY:	build
build:
	go mod download
	go build -buildvcs=false -o app
	./app
