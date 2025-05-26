# Makefile for Go Study Environment

# Run a specific topic file (example: make run topic=01_basics)
run:
	go run $(topic)/*.go

# Build current topic
build:
	go build -o bin/$(topic) $(topic)/*.go

# Run all tests in the project
test:
	go test ./...

# Lint all Go code using golangci-lint
lint:
	golangci-lint run ./...

# Clean build artifacts
clean:
	rm -rf bin/