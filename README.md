# e_commerce-microservices

Rebuilding [e_market backend](https://github.com/nikhil0929/e_market) with microservice architecture

## Docker Products

#### Build Products image from dockerfile

```
docker build -t products_ms -f dockerfiles/Dockerfile.products .
```

This will create a docker image called `products_ms`

#### Run `products_ms` as container

```
docker run --publish 4001:4001 products_ms
```

## Docker Users

#### Build Users image from dockerfile

```
docker build -t users_ms -f dockerfiles/Dockerfile.users .
```

This will create a docker image called `users_ms`

#### Run `products_ms` as container

```
docker run --publish 4002:4002 users_ms
```

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
  - ~~users~~
- Kafka Event Bus(so everything can run more asynchronusly)
- Dockerize microservices (in progress)
- Kubernetes???
- Sensu monitoring for each service
