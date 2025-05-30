# Name of the binary
BINARY_NAME=flightsql-client

# Go build flags
BUILD_FLAGS=-ldflags="-s -w"

# Default: build for host OS
all: build

# Build for host OS
build:
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) main.go

# Clean binary
clean:
	rm -f $(BINARY_NAME)

# Cross-compile for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-linux main.go

# Cross-compile for macOS
build-mac:
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME)-mac main.go

# Cross-compile for Windows
build-win:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BINARY_NAME).exe main.go
