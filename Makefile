
SHELL=/bin/bash
GOCMD=go

GOBUILD=$(GOCMD) build -v
BUILDFLAGS=-ldflags="-s -w"

GOGET=$(GOCMD) get
GETFLAGS=-v

GOTEST=$(GOCMD) test
TESTFLAGS=-cover


all: bin

# makes a minimal installation for use within docker
upx: bin
	upx --ultra-brute bin/place bin/place-server bin/place-git-update

setup:
	mkdir -p bin

deps:
	$(GOGET) $(GETFLAGS) gopkg.in/src-d/go-git.v4
	$(GOGET) $(GETFLAGS) golang.org/x/crypto/ssh

test: deps
	$(GOTEST) $(TESTFLAGS) ./...


bin: bin/place bin/place-server bin/place-git-update

bin/place: setup deps
	$(GOBUILD) $(BUILDFLAGS) -o bin/place cmd/place/main.go

bin/place-server: setup deps
	$(GOBUILD) $(BUILDFLAGS) -o bin/place-server cmd/place-server/main.go


bin/place-git-update: setup deps
	$(GOBUILD) $(BUILDFLAGS) -o bin/place-git-update cmd/place-git-update/main.go


clean:
	rm -f bin/place
	rm -f bin/place-server
	rm -f bin/place-git-update
