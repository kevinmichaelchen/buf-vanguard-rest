# buf-vanguard-rest

A demo showing how one [Buf][buf] API can offer REST, gRPC, and
[Connect][connect] protocols using the [Vanguard][vanguard] library.

Inspired by this [example][vanguard-example].

[buf]: https://buf.build/
[connect]: https://connectrpc.com/
[vanguard]: https://github.com/connectrpc/vanguard-go
[vanguard-example]:
  https://github.com/connectrpc/vanguard-go/blob/main/internal/examples/pets/internal/proto/io/swagger/petstore/v2/pets.proto

## Getting Started

### Run the server

Run the server on port 8080 with:

```shell
go run main.go
```

### Make a REST call

```shell
$ http localhost:8080/foos/1

HTTP/1.1 200 OK  
Accept-Encoding: gzip
Content-Encoding: gzip
Content-Length: 55
Content-Type: application/json
Date: Sat, 13 Jan 2024 18:28:44 GMT

{
    "foo": {
        "id": "1",
        "name": "Bar"
    }
}

```

### Make a Connect call

```shell
curl -v http://localhost:8080/bvr.v1beta1.FooService/CreateFoo \
  -H "Content-Type: application/json" \
  --data-binary @- <<EOF
  {
    "name": "BAR"
  }
EOF
```

or

```shell
pkgx http http://localhost:8080/bvr.v1beta1.FooService/CreateFoo name="BAR"
```