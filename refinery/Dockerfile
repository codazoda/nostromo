# Use the alpine base image (small)
FROM alpine:latest

# Copy all the files into place
COPY files/ /

# Run all the things
RUN apk update && \
    apk update

# Ports we want in the container
EXPOSE 9800

# Run the ore command
CMD /refinery
