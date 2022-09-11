# Otus Homework 06 - Idempotent and commutative APIs

## Solution

Solution - idempotent request id. Client creates unique uuid and sends it in request body

```
POST /orders
{
    "product_id": 111,
    "quantity": 10,
    "shipping_address": "10 Downing Street",
    "request_id": "ce2dc4a8-6d12-4734-9f1f-77e2025e0485"
}
```

Server responds with existing order for the same request

## Install

```shell
helm install app ./06/helm -n otus-06 --create-namespace
```

## Test

Download [Postman collection](https://raw.githubusercontent.com/wuzyk/otus-microservice-arch/main/06/tools/postman/collection.json)

```shell
newman run ~/Downloads/collection.json
```

## Uninstall

```shell
helm uninstall app ./06/helm -n otus-06
```
