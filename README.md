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

### Make a REST request

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

### Make a Connect request

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

### Make a GraphQL request

Thanks to [Tailcall][tailcall], we've able to generate a GraphQL API frmo our
Vanguard HTTP APIs. (See [**`tailcall.graphql`**][tailcall-graphql]).

[tailcall]: https://tailcall.run/
[tailcall-graphql]: ./tailcall.graphql

Once you've started the Tailcall server with:

```shell
pkgx tailcall start ./tailcall.graphql
```

You can visit the [GraphQL Playground][playground] and run the following query:

```graphql
{
  foo(input: { id: "1" }) {
    foo {
      id
      name
    }
  }
}
```

[playground]: http://127.0.0.1:8000/
