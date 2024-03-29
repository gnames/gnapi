PROJ_NAME = gnapi

VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell date -u '+%Y-%m-%d_%H:%M:%S%Z')

OUT_DIR = out/bin
OUTPUT = $(OUT_DIR)/$(PROJ_NAME)

FLAGS_SHARED = CGO_ENABLED=0 GOARCH=amd64
FLAGS_LD=-ldflags "-w -s -X github.com/gnames/$(PROJ_NAME)/pkg.Build=$(DATE) \
                  -X github.com/gnames/$(PROJ_NAME)/pkg.Version=$(VERSION)"
FLAGS_REL = -trimpath -ldflags "-w -s -X \
						github.com/gnames/$(PROJ_NAME)/pkg.Build=$(DATE)"

GOCMD=go
GOINSTALL=$(GOCMD) install $(FLAGS_LD)
GOBUILD=$(GOCMD) build $(FLAGS_LD)
GORELEASE = $(GOCMD) build $(FLAGS_REL)
GOCLEAN=$(GOCMD) clean
GOGET = $(GOCMD) get

all: install

test: deps install
	go test -shuffle=on -race -coverprofile=coverage.txt -covermode=atomic ./...

tools: deps
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %


deps:
	@echo Download go.mod dependencies
	$(GOCMD) mod download;

build:
	mkdir -p $(OUT_DIR)
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(GOBUILD) -o $(OUTPUT);

buildrel:
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(GORELEASE) -o $(OUTPUT);

release: dockerhub
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=linux $(GORELEASE); \
	tar zcvf /tmp/$(PROJ_NAME)-$(VER)-linux.tar.gz $(OUTPUT);

install:
	$(FLAGS_SHARED) $(GOINSTALL);

docker: buildrel
	mkdir -p $(OUT_DIR)
	docker buildx build -t gnames/$(PROJ_NAME):latest -t gnames/$(PROJ_NAME):$(VERSION) .; \
	$(GOCLEAN);

dockerhub: docker
	docker push gnames/gnapi; \
	docker push gnames/gnapi:$(VERSION)
