# e_commerce-microservices

Rebuilding [e_market backend](https://github.com/nikhil0929/e_market) with microservice architecture

## How to run Products Microservice

#### (The only microservice I have implemented thus far)

```
$ go run main.go
```

This should spin up the products microservice on localhost:4001

### Build Products image from dockerfile

```
docker build -t products_ms -f Dockerfile.products .
```

This will create a docker image called `products_ms`

### Create Postgres DB container (instead of hosting on local machine)

```
docker run -d --name my-postgres \
  -e POSTGRES_USER=postgres_user \
  -e POSTGRES_PASSWORD=postgres_password \
  -e POSTGRES_DB=your_database_name \
  -p 5432:5432 \
  postgres:latest
```

postgres_user = postgres
postgres_password = A\*\*\*\*9!
postgres_DB = e_commerce

## TODOs

- Implement other services
  - ~~products~~
  - cart
  - orders
  - users
- Kafka Event Bus(so everything can run more asynchronusly)
- Dockerize microservices
- Kubernetes???
- Sensu monitoring for each service
