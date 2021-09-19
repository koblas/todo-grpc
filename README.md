# grpc-todo

This is designed to show how to do basic gRPC between a web browser and
a server. It uses envoy as the proxy to convert from HTTP to

https://medium.com/@aravindhanjay/a-todo-app-using-grpc-web-and-vue-js-4e0c18461a3e

https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

## Running

- `envoy -c envoy/envoy_v3.yaml`
- `cd ts-client-react; npm start`
- `cd go-server ; go run server.go` or `cd ts-server ; npm start`
- Open browser on port specificed (localhost:1234)

# grpc-todo

A simple todo app using gRPC on both the server and client side

Code for the article [A TODO app using grpc-web and Vue.js](https://medium.com/@aravindhanjay/a-todo-app-using-grpc-web-and-vue-js-4e0c18461a3e)

```console
git clone git@github.com:thearavind/grpc-todo.git
```

## Requirements

### On OSX

```console
brew install protobuf
```

### On Linux

```console
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip

# Unzip
unzip protoc-3.6.1-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
```

## Install Protobuf Generator for Go

```console
go get -u github.com/golang/protobuf/protoc-gen-go
```

## Install Protobuf Generator for Web

```console
git clone https://github.com/grpc/grpc-web /tmp/grpc-web
cd /tmp/grpc-web && sudo make install-plugin
rm -rf /tmp/grpc-web
cd -
```

## Build process on your local machine

### To start the gRPC server

```console
go run server.go
```

### To start the Envoy proxy

```bash
sudo -E docker build -t envoy:v1 .
```

```bash
sudo docker run  -p 8080:8080 --net=host  envoy:v1
```

### To start the client side frontend app

```bash
cd todo-client/
yarn serve
```

## Build process with docker-compose

```console
docker-compose build
docker-compose up
```
