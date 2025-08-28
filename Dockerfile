# --- STAGE 1: Building app ---
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY backend/ .

# Build the application binary
RUN go build -o main ./cmd/api

# --- STAGE 2: Minimal runtime---
FROM debian:12-slim

WORKDIR /app
COPY --from=builder /app/main .

# Open port 8080
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
