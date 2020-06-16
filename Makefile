GOCMD ?= go
GOPATH := $(CURDIR)/vendor:$(GOPATH)

BINARY_GC_SEARCH=out

all: build # test 

build:
	GOOS=linux GOARCH=amd64 $(GOCMD) build -o ./bin/$(BINARY_GC_SEARCH) src/*.go

run:
	$(GOCMD) run src/*.go