# Otus Homework 03

## Install

```shell
helm install app ./03/helm -n otus-03 --create-namespace
```

## Grafana dashboard

[Download](https://github.com/wuzyk/otus-microservice-arch/blob/main/03/helm/charts/grafana/dashboard.json)

![dashboard](./assets/dashboard.png)

## Uninstall

```shell
helm uninstall app ./03/helm -n otus-03
```
