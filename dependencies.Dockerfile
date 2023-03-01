FROM golang:1.20 AS dep
# Add the module files and download dependencies.
ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
RUN go mod download
# Add the shared packages.
COPY ./DB /app/data
COPY ./utils /app/util