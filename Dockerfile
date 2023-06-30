# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /rever

# Copy the entire project
COPY . .

# location of the config file, can be overwritten to point to a different file
ENV TEST_CONFIG=/rever/test/config.json

# Install dependencies
RUN RUN apt install build-essential -y --no-install-recommends
RUN go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest

# Run the tests agains the config file 
CMD ["make", "in-docker-test"]