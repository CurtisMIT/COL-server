# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Curtis Lin"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

RUN ls

RUN pwd
RUN ls -a
CMD ["/app/main"]
