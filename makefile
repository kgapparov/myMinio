.PHONY: build
build: 
	go build -v app/uploader.go
.DEFAULT_GOAL := build