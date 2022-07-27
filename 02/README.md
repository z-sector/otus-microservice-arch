# Otus Homework 02

## Apply manifests

```shell
kubectl create namespace otus-01
kubectl apply -f ./01/k8s/1-deployment.yaml
kubectl apply -f ./01/k8s/2-ingress.yaml
```

## Request service

```shell
curl http://arch.homework/otusapp/vladimir.ratsev/health
```

## Delete namespace

```shell
kubectl delete namespace otus-01
```
