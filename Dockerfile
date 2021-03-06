# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Start from the latest golang base image
FROM golang:latest as builder
# Add Maintainer Info
LABEL maintainer="igor <ighyss@gmail.com>"
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
WORKDIR /app/cmd/whats-app-api-server/
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webapi


######## Start a new stage from scratch #######
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/cmd/whats-app-api-server/ .
COPY appconfig.json .
# Expose port 8080 to the outside world
EXPOSE 9000
# Command to run the executable
CMD ["./webapi"]