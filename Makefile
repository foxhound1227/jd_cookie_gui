# Makefile for JD Cookie GUI (Go version)

.PHONY: build clean test run deps

# Default target
all: build

# Install dependencies
deps:
	go mod download
	go mod tidy

# Build for current platform
build: deps
	go build -ldflags "-s -w" -o jd-cookie-gui$(shell go env GOEXE) .

# Build for all platforms
build-all: deps
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/jd-cookie-gui-windows-amd64.exe .
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/jd-cookie-gui-linux-amd64 .
	# macOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/jd-cookie-gui-macos-amd64 .
	# macOS ARM64
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dist/jd-cookie-gui-macos-arm64 .

# Run the application
run: build
	./jd-cookie-gui$(shell go env GOEXE)

# Test the application
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -f jd-cookie-gui jd-cookie-gui.exe
	rm -rf dist/

# Create dist directory
dist:
	mkdir -p dist

# Help
help:
	@echo "Available targets:"
	@echo "  deps      - Download and install dependencies"
	@echo "  build     - Build for current platform"
	@echo "  build-all - Build for all platforms"
	@echo "  run       - Build and run the application"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build artifacts"
	@echo "  help      - Show this help message"