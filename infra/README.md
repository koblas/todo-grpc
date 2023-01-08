https://golangscript.com/g/kubernetes-cluster-that-shows-how-to-deploy-and-connect-go-backend-react-frontend-mongodb-redis-cache-with-each-other

### Enable registry

```
minikube addons enable registry
kubectl port-forward --namespace kube-system service/registry "$REGISTRY_PORT":80 &

[ -n "$(docker images -q alpine)" ] || docker pull alpine
docker run --rm -it --network=host alpine ash -c "apk add socat && socat TCP-LISTEN:$REGISTRY_PORT,reuseaddr,fork TCP:$(minikube ip):$REGISTRY_PORT"
```

### Enable Ingress

```
minikube addons enable ingress
```

### Secrets

```
kubectl apply -f infra/dev-variables-env-configmap.yaml
kubectl delete secret common-secrets --ignore-not-found
kubectl create secret generic common-secrets --from-env-file=../.env
```

    --namespace=new-namespace


### Day-to-day startup

```
minikube start
minikube addons enable ingress
minikube tunnel
tilt up
```
