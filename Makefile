# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

RW_BINARY_NAME="rw"

# Build directives
all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/...

# Test directives
test:
	$(GOTEST) ./...

# Clean up
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f $(GRAPH_BINARY_NAME)


.PHONY: rw
rw:
	$(GOBUILD) -o $(RW_BINARY_NAME) ./main.go

