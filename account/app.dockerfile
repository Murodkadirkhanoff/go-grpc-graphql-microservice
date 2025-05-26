# Stage 1: Build
FROM golang:1.22-alpine3.20 AS build

# Устанавливаем инструменты сборки
RUN apk --no-cache add gcc g++ make ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной исходный код
COPY . .

# Сборка бинарника
RUN go build -o /go/bin/app ./account/cmd/account

# Stage 2: Minimal runtime
FROM alpine:3.20

WORKDIR /usr/bin

# Копируем только скомпилированное приложение
COPY --from=build /go/bin/app .

EXPOSE 8080

CMD ["./app"]
