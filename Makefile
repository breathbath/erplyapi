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

now=$(shell date +"%Y%m%d%H%M%S")

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
	@echo "gendoc			 : Generates API docs"
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

gendoc:
	docker build -f ./docker/APIDoc_Dockerfile.txt -t erply_apidoc:latest .
	docker run --rm -it -v ${CURDIR}/apidoc:/home/apidoc/apidoc -v ${CURDIR}:/home/apidoc/source erply_apidoc:latest apidoc --input /home/apidoc/source --output /home/apidoc/apidoc -v
genmig:
	#cat migrations/template.txt | sed -e 's/MigName/Migration${now}/g' > migrations/items/Migration${now}.go
	docker run --rm -it -v ${CURDIR}/migrations:/home/migrations busybox sh -c 'sed "s/MigName/Migration${now}/g" /home/migrations/template.txt  > /home/migrations/items/Migration${now}.go'
	docker run --rm -it -v ${CURDIR}/migrations:/home/migrations busybox sed -i "s/^}/	registry.RegisterMigration(items.Migration${now}{})\\n}/g" /home/migrations/registryInitialiser.go
