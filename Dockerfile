# Use the official Golang image as the base image for building
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o bin/translator ./cmd/translator

# Use a minimal base image for the final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN adduser -D translator

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/translator .

# Copy configuration files
COPY --from=builder /app/config.yaml.example ./config.yaml

# Copy web assets
COPY --from=builder /app/web ./web

# Change ownership to the non-root user
RUN chown -R translator:translator /app

# Switch to the non-root user
USER translator

# Expose the port the application listens on
EXPOSE 8080

# Command to run the application
ENTRYPOINT ["./translator"]
