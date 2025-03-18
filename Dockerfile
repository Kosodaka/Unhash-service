FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /unhash-service ./cmd/main.go

# Финальная стадия
FROM alpine:3.18

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /unhash-service /app/unhash-service

COPY .env /app/.env
CMD ["/app/unhash-service"]