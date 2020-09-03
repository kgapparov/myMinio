.PHONY: build
build: 
	go build -v app/minioUploader
.DEFAULT_GOAL := build