# Use official Golang image to build the app
FROM golang:1.23.7 AS builder

WORKDIR /app

# Copy source files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go binary
RUN go build -o devtracker .

# Use a lightweight container for runtime
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/devtracker .

# Expose the port
EXPOSE 8080

# Run the binary
CMD ["./devtracker"]

