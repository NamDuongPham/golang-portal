# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Copy go.mod và go.sum để cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build binary từ main.go
WORKDIR /app/cmd/server
RUN go build -o main .

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/cmd/server/main .
EXPOSE 8080
CMD ["./main"]
