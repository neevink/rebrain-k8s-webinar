FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY main.go .
COPY go.mod .
RUN go build -o app

FROM alpine:3.21
WORKDIR /root
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
