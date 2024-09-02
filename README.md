## ProdSub Service

### Requirements

- Go 1.23.x
- Postgres DB
- protobuf
- grpcurl (for testing grpc endpoints)

### Setup instructions

- Run command `go mod tidy`
- Create a local postgres database using name `prodsub_db` with default configurations i.e
  `host=localhost user=postgres dbname=prodsub_db port=5432 sslmode=disable`
- Enable uuid generation in postgres by creating extension `create extension if not exists "uuid-ossp";`

### Unit tests

Run command to execute all unit tests on the project

```bash
    make test
```

### Run Grpc server

To run the server use the command

```bash
  go run cmd/server/main.go
```

## Usuage

- List available endpoints using command

```bash
     grpcurl -plaintext localhost:50051 list
```

- Create Product

```bash
grpcurl -plaintext -d '
        {
        "name": "shoe",
        "description": "latest nikey shoes",
        "price": 89.99,
        "product_type": 0,
        "product_attribute": {
            "weight": 49.5,
            "dimensions": "16*16"
        }
        }' localhost:50051 prodsub.product.v1.ProductService/CreateProduct
```

- Create subscription

```bash

grpcurl -plaintext -d '
            {
            "product_id": "52f257b6-1bae-4bc3-af6f-ee552bcfa2b8",
            "plan_name": "Monthly",
            "duration": 88,
            "price": 4.5
            }' localhost:50051 prodsub.product.v1.SubscriptionService/CreateSubscription
```

### Limitations due to time

- No Docker setup
- No Grpc server auth
- All non-validation errors resolve the same error status code (Internal)
