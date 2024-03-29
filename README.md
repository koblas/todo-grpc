# grpc-todo

This project was originally conceived as a way to test gRPC and TypeScript it has since
morphed into a tour-de-force of how to build a application framework that builds a 
micro services AWS lambda framework. 

## USeful

https://www.obytes.com/blog/go-serverless-part-4-realtime-interactive-and-secure-applications-with-aws-websocket-api-gateway

## Detailed History

* TypeScript gRPC / Docker
*  

## Components



---
## Original README that needs to be cleaned up 


This is designed to show how to do basic gRPC between a web browser and
a server. It uses envoy as the proxy to convert from GRPC-WEB or JSON to GRPC.

## Running

The fastest way to get all of the components up and running is via the included `docker-compose` setup.  

- `docker-compose build`
- `docker-compose up`

### Build and run locally

If you're more for playing around without having to go through the world of docker.  You need three windows open and can run each of these commands.  *Note: Skipping over the npm install in the directories*

- `cd go-server ; go run server.go`  or `cd ts-server ; npm start`
- `cd ts-client-react ; npm start`
- `cd envoy ; envoy -c envoy.yaml`
- Open browser on port specificed (localhost:1234)

### Building the protobuf / gRPC definitions

Base requirements:
* OSX: `brew install protobuf`
* Linux: install protobuf (via your favorite package manager)

Protobuf files (which are checked into the repo, but if you're changing things)

`cd protos ; make`

The directory `google/api` was created by:

    mkdir -p google/api
    cd protos/google/api
    wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
    wget https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto

## References

- [A TODO app using grpc-web and Vue.js](https://medium.com/@aravindhanjay/a-todo-app-using-grpc-web-and-vue-js-4e0c18461a3e)
- https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

### Graveyard of notes


#### Go Example 
  - production ready gRPC server - https://github.com/apssouza22/grpc-production-go
  - This is good https://www.youtube.com/watch?v=bJEEGJa_C4o

#### Install Protobuf Generator for Go

```console
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### Install Protobuf Generator for Web

```console
git clone https://github.com/grpc/grpc-web /tmp/grpc-web
cd /tmp/grpc-web && sudo make install-plugin
rm -rf /tmp/grpc-web
cd -
```


### Running on AWS Lambda

Create the following SSM Values

* /common/jwt_secret -- 32 characters

### Running K8S (minikube)

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
kubectl delete secret common-secrets --ignore-not-found
kubectl create secret generic common-secrets --from-env-file=.env
```

    --namespace=new-namespace


### Day-to-day startup

```
minikube start
minikube addons enable ingress
minikube tunnel
tilt up
kubectl port-forward service/minio 9000:9000 9001:9001
```

TODO -- Need to setup the MINIO->NATS subscription (the queue is created, but the subscription is not)

Cleanup disk space
```
minikube ssh -- docker system prune
```

## K3S

```
k3d cluster create mycluster --registry-create mycluster-registry --k3s-arg "--disable=traefik@server:0" -p "8000:80@loadbalancer"
kubectl create secret generic common-secrets --from-env-file=.env
sudo kubectl port-forward svc/ingress-nginx-controller 80:80
```
