# Build STAGE
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]