# Go parameters
GOCMD=go
GOS=gos
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=IDE
GOSOUTPUT=server_out.go


all: test build
build: 
		$(GOS) --export
		$(GOBUILD) -o $(BINARY_NAME)        
test: 
		$(GOTEST) -bench=. -benchmem
clean: 
		$(GOCLEAN)
		rm -f bindata.*
		rm -f $(GOSOUTPUT)
		rm -f $(BINARY_NAME)
run:
		$(GOS) --export
		$(GOBUILD) -o $(BINARY_NAME)
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/cheikhshift/gos
		$(GOS) deps
