# VEM-GO

This is a VEM example microservice repo using the golang services and api gateway.
This VEM offers a great starting point for anyone looking to build their microservice
architecture using primarily golang.

## Quickstart

In order to run this microservice application, be sure to have docker installed. (More
docs will be coming soon to walk through this process, but for now, go
[here](https://docs.docker.com/engine/installation/))

Clone this repo recursivly to include all of the submodule services.

```bash
$ git clone --recursive https://github.com/vems/vem-go.git
```

`cd` into the directory where it was cloned.

```bash
$ cd vem-go
```

Build all docker containers required to run this application.

```bash
$ docker-compose build
```

Start all of the docker containers required to run this application.

```bash
$ docker-compose up
```

The application is now running. Visit http://localhost:8080/api/sum?first=2&second=2 to see
the result of the sum service calculating `2 + 2`.
