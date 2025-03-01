# Этап сборки Go приложения
FROM golang:1.23.3-alpine AS auto-tg-app


WORKDIR /app

# Копирование и установка зависимостей
COPY go.mod go.sum ./
COPY vendor/ ./vendor/

# Копирование остального
COPY .. ./
COPY /docker/.env ./

CMD go run ./cmd/main.go