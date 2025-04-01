# Define the binary name
BINARY=nusantara-data-api

# Define the main Go file
MAIN=main.go

# Define output directory
BUILD_DIR=bin

# Define environment variables
GO_ENV=development
PORT_ENV?=3000

# Default target
.PHONY: all
all: build-run

# Build the application with optimizations
.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY) with optimizations..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) CUSTOM_ENV=$(CUSTOM_ENV) go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY) $(MAIN)
	@echo "Build complete!"

# Build and run the application
.PHONY: build-run
build-run: build
	@echo "Running $(BINARY)..."
	CUSTOM_ENV=$(CUSTOM_ENV) $(BUILD_DIR)/$(BINARY)

# Run the application in development mode
.PHONY: dev
dev:
	@echo "Running in development mode..."
	GO_ENV=$(GO_ENV) CUSTOM_ENV=$(CUSTOM_ENV) go run $(MAIN)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	@echo "Cleanup complete!"

# Format the code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint the code
.PHONY: lint
lint:
	go vet ./...

# Run tests
.PHONY: test
test:
	go test ./...
