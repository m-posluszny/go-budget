FROM golang:1.20.5-alpine

ENV GIN_MODE=release

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 7000
ENV APP_MODE PRODUCTION
COPY ./src ./src
RUN go build  -o go-ynab-exec ./src/main.go
COPY .env .env

CMD "./go-ynab-exec"