FROM golang:1.20.5-alpine

ENV GIN_MODE=release

WORKDIR /app
COPY go.mod ./

EXPOSE 7000
ENV APP_MODE PRODUCTION
COPY . .
RUN go build  -o go-ynab-exec ./src/main.go

CMD "./go-ynab-exec"