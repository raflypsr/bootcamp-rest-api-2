# Start from the official Go image
FROM golang:1.21-alpine
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go app
RUN go build main.go
# Expose port 8080 to the outside world
EXPOSE 9090
# Set environment variable for Gin mode
ENV GIN_MODE=release
# Run the executable
CMD ["/app/main"]