FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o sys-monitor .
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sys-monitor .
EXPOSE 9090
CMD ["./sys-monitor"]
