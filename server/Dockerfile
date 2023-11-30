# syntax=docker/dockerfile:1
FROM golang:1.21.3-alpine AS build_base
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /go-app-code

# Copy the entire project
COPY . .

# Build the server
RUN go build -o /app/server ./server/main.go

# Start fresh from a smaller image
FROM alpine:latest
RUN apk add ca-certificates

WORKDIR /app

# Copy just the needed go binary compiled
COPY --from=build_base /app/server .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary 
CMD ["./server"]