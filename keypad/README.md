# Chain

## Run with php built-in server

`php -S localhost:9801`

## Build the Docker image

`docker build --no-cache --tag keypad .`

## Create and start the container

```
docker run \
    --hostname keypad \
    --publish "9801:9801" \
    --name keypad \
    --detach \
    keypad \
    /usr/bin/php -S 0:9801 -t /var/www
```

## Create and start an image based off the published container

```
docker run \
    --hostname chain \
    --publish "9801:9801" \
    --name chain \
    --detach \
    deseretdigital/nostromo:chain \
    /usr/bin/php -S 0:9801 -t /var/www
```
