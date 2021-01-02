.PHONY: build
build:
		go build -v

.PHONY: run
run:
		go run .


.DEFAULT_GOAL := build