# Use official Go image as builder
FROM golang:1.21 as builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN go build -o app

# Final lightweight container
FROM debian:bullseye-slim

WORKDIR /app

# Copy compiled Go binary from builder
COPY --from=builder /app/app .

# Copy any static files or .env if needed (optional)
# COPY .env .   # Only if you're not using runtime ENV vars

# Port app listens on
EXPOSE 8080

# Start the app
CMD ["./app"]
