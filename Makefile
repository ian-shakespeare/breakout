MAIN_FILE := main.go

all: run

build:
	go build .

run:
	go run $(MAIN_FILE)

.PHONY: all build run
