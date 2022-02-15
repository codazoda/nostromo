# Refinery

## Compile

`go build refinery.go`

## Compile for linux

`GOOS=linux GOARCH=amd64 go build -o files/refinery refinery.go`

## Build the Docker image

`docker build --no-cache --tag refinery .`

## Create and start the container

```
docker run \
    --hostname refinery \
    --publish "9800:9800" \
    --name refinery \
    --detach \
    refinery \
    /refinery
```

## Create and start an image based off the published container

```
docker run \
    --hostname refinery \
    --publish "9800:9800" \
    --name refinery \
    --detach \
    deseretdigital/nostromo:refinery \
    /refinery
```
