# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.18 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/server

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS calls and timezone data
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the timezone (optional)
ENV TZ=UTC

# Use an unprivileged user.
USER appuser

WORKDIR /app

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Expose port 50051 for the application
EXPOSE 50051

# Command to run the executable
CMD ["./server"]