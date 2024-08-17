MAIN_FILE := main.go

all: run

build:
	go build .

run:
	go run $(MAIN_FILE)

clean:
	go clean

.PHONY: all build run clean
