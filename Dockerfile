FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o ./golang-api

EXPOSE 8001

CMD ["./golang-api"]
