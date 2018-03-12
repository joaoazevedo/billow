# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY:=$(notdir $(shell pwd))

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o bin/$(BINARY)

docker: build
	docker build -t jazevedo/$(BINARY) .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
