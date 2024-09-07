# Variables
BINARY_NAME=tasks_project
BINARY_DIR=bin

# Go commands
GOBUILD=go build
GOINSTALL=go install

# Targets and recipes
.PHONY: all build install clean

all: build install

build:
	@echo "Building..."
	@mkdir -p $(BINARY_DIR)
	@$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME)

install:
	@echo "Installing..."
	@$(GOINSTALL)

clean:
	@echo "Cleaning..."
	@rm -rf $(BINARY_DIR)