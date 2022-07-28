# Otus Homework 02

## Install

```shell
helm install app ./02/helm -n otus-02
```

## Test

Download [Postman collection](https://github.com/wuzyk/otus-microservice-arch/02/tools/postman/collection.json)

```shell
newman run ~/Downloads/collection.json
```

## Uninstall

```shell
helm uninstall app ./02/helm -n otus-02
```
