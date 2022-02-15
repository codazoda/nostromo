# Use the alpine base image (small)
FROM alpine:latest

# Copy just the index file into place
COPY /index.php /var/www/index.php

# Install php
RUN apk update && \
    apk add php

# Ports we want in the container
EXPOSE 9802

# Run the php dev server (cheating)
CMD /usr/bin/php -S 0:9802 -t /var/www/
