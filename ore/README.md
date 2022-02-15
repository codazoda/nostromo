# Ore

## Compile

`go build ore.go`

## Compile for linux

`GOOS=linux GOARCH=amd64 go build -o files/ore ore.go`

## Build the Docker image

`docker build --no-cache --tag ore .`

## Create and start the container

```
docker run \
    --hostname ore \
    --publish "9801:9801" \
    --name ore \
    --detach \
    ore \
    /ore
```
