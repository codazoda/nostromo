# Refinery

## Compile

`go build refinery.go`

## Run locally

`./refinery`

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

## Create and start an image based off the published container (or local if we have it)

```
docker run \
    --hostname refinery \
    --publish "9800:9800" \
    --name refinery \
    --detach \
    deseretdigital/sms_refinery \
    /refinery
```
