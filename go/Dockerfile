# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/app

# Runtime Stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

# Install ca-certificates for SSL/TLS
RUN apk --no-cache add ca-certificates

# Expose port if needed (not strictly needed for this CLI worker, but good practice)
# EXPOSE 8080

ENTRYPOINT ["./main"]
