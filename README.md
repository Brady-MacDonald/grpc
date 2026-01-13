# URL Shortener

Project structure

```sh
url-shortener/
├── go.mod
├── gateway/
│   ├── main.go
│   └── handlers.go
│
├── proto/
│    └── idgen/
│        └── base62.go
│
├── proto/
│   ├── shortener.proto
│   └── analytics.proto
│
├── services/
│   ├── shortener/
│   │   ├── main.go
│   │   ├── server.go
│   │   └── store.go
│   │
│   └── analytics/
│       ├── main.go
│       ├── server.go
│       └── store.go
```

> `/gateway`    - Entry point of project, handles the HTTP -> gRPC
> `/internal`   - Shared logic
> `/proto`      - Used for contract layer
> `/services`   - Independently run able services

## Go Dependencies

```bash
go mod tidy
```

## Protocol Buffer Compiler

Ensure you have the `protoc` compiler available on your `$PATH`
Additional plugins required for Go code generation

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Run the compile command

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto
```

```bash
protoc \ 
  --go_out=. \ 
  --go_opt=module=github.com/Brady-MacDonald/grpc \ 
  --go-grpc_out=. \ 
  --go-grpc_opt=module=github.com/Brady-MacDonald/grpc \ 
  proto/hello/hello.proto

```
