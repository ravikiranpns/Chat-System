# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN go build -o main ./cmd

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh
CMD ["./wait-for-it.sh", "cassandra:9042", "--", "./main"]
