# gRPC service and client example

## Building and Running Service

```
$ cd server
$ go build -i -v -o bin/server
$ ./bin/server
```

## Building and Running Client
```
$ cd client
$ go build -i -v -o bin/client
$ ./bin/client
```

## Generate Server and Client side code
```
$ protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:server/ecommerce
```

```
$ protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:client/ecommerce 
```