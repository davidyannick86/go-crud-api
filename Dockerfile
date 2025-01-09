# Build stage
FROM golang:1.23.4-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Production stage
FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]