FROM golang:1.16.0-alpine AS builder


RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .
CMD ["/app/main"]
