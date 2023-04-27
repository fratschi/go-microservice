# go-microservice

# basic usage

## run locally

```bash
go run cmd/main.go
```

## run docker

Building the docker image

```bash
docker build -t "go-microservice" .
```

Run the docker image without a config file. This will use some defaults.

```bash
docker run "go-microservice"
```

Run docker immage with a config file. 
```bash
docker run --mount type=bind,source=<ABSOLUTE PATH>/config.yaml,target=/config/config.yaml "go-microservice"

```
