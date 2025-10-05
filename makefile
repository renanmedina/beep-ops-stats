APP_NAME := beep_ops_stats
MAIN_GO_FILE := cmd/main.go
.PHONY: all build run clean
all: build
build:
	go build  -ldflags="-s -w" -trimpath -o bin/$(APP_NAME) $(MAIN_GO_FILE)
run: build
	./bin/$(APP_NAME)
clean:
	go clean
	rm -f bin/$(APP_NAME)