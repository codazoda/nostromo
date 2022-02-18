# Use the alpine base image (small)
FROM alpine:latest

# Copy all the files into place
COPY files/ /

# Run all the things
RUN apk update && \
    apk update

# Ports we want in the container
EXPOSE 9802

# Run the command
CMD /clock
