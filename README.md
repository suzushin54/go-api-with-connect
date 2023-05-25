# Go API with Connect

This is a simple API written in Go that uses the [Connect](https://connect.build/).

## Getting Started

```shell
make setup
make install-check
make buf

go run ./cmd/main.go 

$ curl \                                                                                                                                                                              [master]
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/my_service.v1.MyService/HelloWorld
{"message":"Hello Connect!"}

```