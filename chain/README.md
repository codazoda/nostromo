# Chain

## Run with php built-in server

`php -S localhost:9802`

## Build the Docker image

`docker build --no-cache --tag chain .`

## Create and start the container

```
docker run \
    --hostname chain \
    --publish "9802:9802" \
    --name chain \
    --detach \
    chain \
    /usr/bin/php -S 0:9802 -t /var/www
```

## Create and start an image based off the published container

```
docker run \
    --hostname chain \
    --publish "9802:9802" \
    --name chain \
    --detach \
    deseretdigital/nostromo:chain \
    /usr/bin/php -S 0:9802 -t /var/www
```
