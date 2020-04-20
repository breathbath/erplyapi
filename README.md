# ERPLY API Proxy

## Project description
This project is intended to proxy calls to [ERPLY API](https://learn-api.erply.com/requests).

## Supported API methods

- [getCustomers](https://learn-api.erply.com/requests/getcustomers)
- [saveCustomer](https://learn-api.erply.com/requests/savecustomer)

## Prerequisites
- Docker ^19.03.8
- Docker-compose ^3.4
- GNU Make ^3.81

## Getting started

The main entry point is make commands list defined in Makefile.
To get started you need to build all docker images:

    make build-docker

Place .env file in the root folder with the needed configuration options. Use env.dist file for the inspiration.

Now you can start all services with the docker-compose:

    make up
    
To stop all services:

    make down

To build local binary:

    make build
    
To run unit-tests:

    make test
    
To generate API docs:

    make gendoc
    
To show full help:

    make help
