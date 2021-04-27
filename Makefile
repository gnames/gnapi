VERSION = $(shell git describe --tags)
VER = $(shell git describe --tags --abbrev=0)
DATE = $(shell date -u '+%Y-%m-%d_%H:%M:%S%Z')
FLAG_MODULE = GO111MODULE=on
FLAGS_SHARED = $(FLAG_MODULE) CGO_ENABLED=0 GOARCH=amd64
FLAGS_LD=-ldflags "-w -s -X github.com/gnames/gnapi.Build=${DATE} \
                  -X github.com/gnames/gnapi.Version=${VERSION}"
NO_C = CGO_ENABLED=0
GOCMD=go
GOINSTALL=$(GOCMD) install $(FLAGS_LD)
GOBUILD=$(GOCMD) build $(FLAGS_LD)
GOCLEAN=$(GOCMD) clean
GOGET = $(GOCMD) get

all: install

test: deps install
	$(FLAG_MODULE) go test ./...

tools: deps
	@echo Installing tools from tools.go
	@cat gnapi/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %


deps:
	@echo Download go.mod dependencies
	$(GOCMD) mod download; \

build:
	cd gnapi; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) $(NO_C) $(GOBUILD);

release: dockerhub
	cd gnapi; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=linux $(GOBUILD); \
	tar zcvf /tmp/gnapi-${VER}-linux.tar.gz gnapi; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=darwin $(GOBUILD); \
	tar zcvf /tmp/gnapi-${VER}-mac.tar.gz gnapi; \
	$(GOCLEAN); \
	$(FLAGS_SHARED) GOOS=windows $(GOBUILD); \
	zip -9 /tmp/gnapi-${VER}-win-64.zip gnapi.exe; \
	$(GOCLEAN);

install:
	cd gnapi; \
	$(FLAGS_SHARED) $(GOINSTALL);

docker: build
	docker build -t gnames/gnapi:latest -t gnames/gnapi:${VERSION} .; \
	cd gnapi; \
	$(GOCLEAN);

dockerhub: docker
	docker push gnames/gnapi; \
	docker push gnames/gnapi:${VERSION}
