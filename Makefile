# Makefile for gin-demo project

# Variables
PROJECT_NAME=gin-demo
BINARY_NAME=gin-demo
DOCKER_IMAGE_NAME=gin-demo-image
VERSION=latest

# Go build flags
GO_BUILD_FLAGS=-ldflags "-s -w"

# Directories
SRC_DIR=./cmd
DOCKER_DIR=./

# Default target
.DEFAULT_GOAL:=build

# Build the binary
build:
	@echo "Building ${BINARY_NAME}..."
	GO111MODULE=on GOARCH=amd64 go build ${GO_BUILD_FLAGS} -o ${BINARY_NAME} ${SRC_DIR}/main.go

# Run the binary locally
run:
	@echo "Running ${BINARY_NAME} locally..."
	@./${BINARY_NAME}

# Test the project
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean the build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f ${BINARY_NAME} ${BINARY_NAME1}

# Run all tests and build the project
all: test build

# Help target to display available commands
help:
	@echo "Available commands:"
	@echo "  make build         - Build the binary"
	@echo "  make run           - Run the binary locally"
	@echo "  make test          - Run tests"
	@echo "  make clean         - Clean build artifacts and Docker containers"
	@echo "  make all           - Run tests and build the binary"
	@echo "  make help          - Display this help message"