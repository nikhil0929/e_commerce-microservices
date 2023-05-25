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

## Notes:

When you run pgAdmin4 and PostgreSQL in separate Docker containers and on the same machine, they are essentially on two different "machines" from a networking perspective, even though they're both technically on your Mac.

The term localhost inside the pgAdmin4 container refers to the pgAdmin4 container itself, not your Mac host where the PostgreSQL container is running. This is why you're unable to connect to the PostgreSQL server.

The easiest way to resolve this issue is by using Docker's built-in networking capabilities. By default, Docker provides a special network named bridge, and all containers run in this network unless specified otherwise.

When containers are in the same network, they can reach each other by the container name. So you should be able to use cart_postgres as the hostname to reach your PostgreSQL server from within the pgAdmin4 container.

I created my own network using

```bash
docker network create e_commerce_microservices
```

and then i ran the cart_postgres container using the following where i specify the network name:

```bash
docker run -d --network=e_commerce_microservices --name cart_postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Arnik0929! -e POSTGRES_DB=cart_postgres -p 5434:5432 postgres:latest
```

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
