
FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build -o /products

EXPOSE 4001

CMD [ "/products" ]