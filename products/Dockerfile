# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY */ ./

RUN go build -o /ecm-ms

EXPOSE 4001

CMD ["/ecm-ms"]