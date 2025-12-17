FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY do-data-go/go.mod do-data-go/go.sum ./
RUN go mod download

# Copy the rest of the Go app
COPY do-data-go/ ./

# Build custom PocketBase Go app
RUN go build -o pocketbase-app main.go

FROM alpine:latest

WORKDIR /pb

# Copy the built binary from the builder stage
COPY --from=builder /app/pocketbase-app /pb/pocketbase-app

# Copy static/public files
# COPY do-data-go/pb_public ./pb_public

# Set the entrypoint to your Go app
CMD ["/pb/pocketbase-app", "serve", "--http=0.0.0.0:8080"]