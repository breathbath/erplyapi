# set default shell
SHELL = bash -e -o pipefail

# Variables
VERSION                  ?= $(shell cat ./VERSION)


## Docker related
DOCKER_EXTRA_ARGS        ?=
DOCKER_REGISTRY          ?=
DOCKER_REPOSITORY        ?=
DOCKER_TAG               ?= ${VERSION}
DOCKER_BUILD_ARGS        ?=${DOCKER_EXTRA_ARGS} --build-arg version="${VERSION}"
IMAGE_NAME				 ?=erply_api

default: build

help:
	@echo "Usage: make [<target>]"
	@echo "where available targets are:"
	@echo
	@echo "build             : Build ERPLY API binary"
	@echo "build-docker      : Build ERPLY API docker image"
	@echo "help              : Print this help"
	@echo "test              : Run unit tests, if any"
	@echo "up				 : Bring all services up with docker-compose"
	@echo "down				 : Bring all services down with docker-compose"
	@echo

build:
	mkdir -p bin
	go build -race -o bin/erply_server \
	    main.go

build-docker:
	docker build $(DOCKER_BUILD_ARGS) -t ${IMAGE_NAME}:${VERSION} -t ${IMAGE_NAME}:latest  -f docker/Dockerfile .

test:
	go test -race -v -p 1 ./...

up:
	docker-compose -f docker-compose.yml up -d

down:
	docker-compose -f docker-compose.yml down
