# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

# Binary name
GRAPH_BINARY_NAME=graph

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

# Build directive for the graph application
.PHONY: graph
graph:
	go build -o $(GRAPH_BINARY_NAME) ./cmd/graph


