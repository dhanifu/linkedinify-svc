BINARY_NAME := linkedinify-svc
BINARY_PATH := ./tmp/bin
MAIN_FILE := cmd/api/main.go

.PHONY: all build run clean

all: build run

build:
	@mkdir -p $(BINARY_PATH)
	CGO_ENABLED=0 go build -mod=readonly -o $(BINARY_PATH)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete"

run:
	@echo "Running..."
	$(BINARY_PATH)/$(BINARY_NAME)

clean:
	rm -rf tmp/bin
	@echo "Removed bin directory"