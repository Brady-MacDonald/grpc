# Protocol Buffers

## Protocol Buffer Compiler

The `protoc` executable is used to build the language specific implementation

The schema created in the `.proto` is language neutral
This contract is written so that it can be implemented in various languages
It is up to the protocol buffer compiler to generate the language

To generate go code from a `.proto` file we need to use 

[protoc-gen-go](https://google.golang.org/protobuf) 
> Base level proto buf code generator
> Pure proto buf logic
> Marshaling / Unmarshaling code and message definitions

[protoc-gen-go-grpc](https://google.golang.org/grpc/cmd/protoc-gen-go-grpc) 
> gRPC code generator for Go
> Used for service code in gRPC context
> Server / Client interfaces

These packages can be installed using the `go install {pkg}` command to your `$GOBIN` location

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
protoc 
    --go_out=. /
    --go_opt=module=github.com/your-username/grpc-demo \  # Your module path  
    --go-grpc_out=. \      
    --go-grpc_opt=module=github.com/your-username/grpc-demo \  # Your module path  
    proto/hello/hello.proto
```

`--go_out` for pure proto buf code generation
`--go-grpc_out` for gRPC specific code generation
