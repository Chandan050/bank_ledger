# Use official Golang image as the build environment
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o banking-ledger cmd/main.go

# Use a minimal image for running the app
FROM alpine:latest

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/banking-ledger .

# Expose port 8080 for the API
EXPOSE 8080

# Command to run the binary
CMD ["./banking-ledger"]
