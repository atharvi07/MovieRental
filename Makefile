# Makefile for a Go project

# Variables
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOCLEAN := $(GOCMD) clean

build:
	$(GOBUILD) -o bin/main ./cmd

run:
	$(GORUN) cmd/main.go

clean:
	$(GOCLEAN)
	rm -rf bin/

test:
	go test ./...

db:
	docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=root -d postgres

populateDatabase:
