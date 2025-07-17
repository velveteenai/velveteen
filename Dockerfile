# Start from golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy only the go.mod to download and cache dependencies
COPY go.mod go.sum Makefile ./
RUN make setup

# Copy everything from the current directory to the PWD inside the container
COPY . .

CMD ["go", "run", "./cmd/http-server/."]
