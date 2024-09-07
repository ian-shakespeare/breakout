MAIN_FILE := main.go
OUTPUT_FILE := breakout

all: run

build:
	go build -o $(OUTPUT_FILE)

run:
	go run $(MAIN_FILE)

clean:
	go clean

.PHONY: all build run clean
