# Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Required packages
RUN apk add --no-cache git

# Copy go mod and sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the source
COPY . .

# Build binary
RUN go build -o app ./cmd/app

# Run stage
FROM alpine:latest

WORKDIR /root/

# Copy binary from build stage
COPY --from=builder /app/app .

# Copy .env (optional for local dev)
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./app"]
