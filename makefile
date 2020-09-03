.PHONY: build
build: 
	go build -v app/main.go
.DEFAULT_GOAL := build