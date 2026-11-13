# Orders API

A small learning microservice written in Go that manages orders, using Redis as its data store. Based on the Dreams of Code microservice tutorial.

## Stack

- [Go](https://go.dev/) with the standard `net/http` server
- [chi](https://github.com/go-chi/chi) for routing and middleware
- [go-redis](https://github.com/redis/go-redis) as the Redis client
- Docker Compose for the local Redis instance
- [Task](https://taskfile.dev/) as the task runner

## Project structure

```
.
├── main.go              # Entry point, graceful shutdown on Ctrl+C
├── application/         # App wiring: server, Redis client, routes
├── handler/             # HTTP handlers for orders
├── model/               # Order and LineItem data models
├── repository/order/    # Redis-backed order repository
└── deployments/         # Docker Compose files (Redis)
```

## Getting started

Start Redis:

```sh
task up
```

Run the API (listens on `:3000`):

```sh
go run .
```

Other tasks:

```sh
task logs          # Follow container logs
task exec:redis    # Open redis-cli inside the Redis container
```

## Endpoints

| Method   | Path           | Description                                       |
| -------- | -------------- | ------------------------------------------------- |
| `POST`   | `/orders`      | Create an order                                   |
| `GET`    | `/orders`      | List orders (paginated with `?cursor=`)           |
| `GET`    | `/orders/{id}` | Get an order by ID                                |
| `PUT`    | `/orders/{id}` | Update status (`shipped` or `completed`)          |
| `DELETE` | `/orders/{id}` | Delete an order                                   |

### Examples

```sh
# Create
curl -X POST localhost:3000/orders \
  -d '{"customer_id":"2f1c6a9e-8a3b-4c1d-9e2f-1a2b3c4d5e6f","line_items":[{"item_id":"7d3e1b2a-4c5d-4e6f-8a9b-0c1d2e3f4a5b","quantity":2,"price":1500}]}'

# List
curl localhost:3000/orders

# Ship, then complete
curl -X PUT localhost:3000/orders/<id> -d '{"status":"shipped"}'
curl -X PUT localhost:3000/orders/<id> -d '{"status":"completed"}'

# Delete
curl -X DELETE localhost:3000/orders/<id>
```

An order must be `shipped` before it can be `completed`.

## Storage

Each order is stored as JSON under the key `order:<id>`, and every key is tracked in the `orders` set, which is used for paginated listing with `SSCAN`.
