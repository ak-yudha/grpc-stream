# grpc-stream

Installing the gRPC Go plugin

```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go

go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

export PATH="$PATH:$(go env GOPATH)/bin"
```

# Running the application

1. Install the dependencies

```bash
go mod tidy
```

2. Run the server

```bash
go run server/main.go
```

3. Run the client

```bash
go run client/main.go
```