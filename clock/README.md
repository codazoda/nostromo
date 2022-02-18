# Ore

## Compile

`go build ore.go`

## Compile for linux

`GOOS=linux GOARCH=amd64 go build -o files/clock clock.go`

## Build the Docker image

`docker build --no-cache --tag clock .`

## Create and start the container

```
docker run \
    --hostname clock \
    --publish "9802:9802" \
    --name clock \
    --detach \
    clock \
    /clock
```
