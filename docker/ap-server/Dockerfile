FROM golang:1.17-alpine3.13 AS builder

WORKDIR /app
COPY . /app

RUN go mod tidy -compat=1.17
RUN go get github.com/cosmtrek/air
CMD ["air"]