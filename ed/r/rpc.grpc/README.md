gRPC
-

[grpc](https://grpc.io/docs)
[golang](https://grpc.io/docs/languages/go/quickstart/)
[grpcurl](https://github.com/fullstorydev/grpcurl)

````sh
# golang struct -> protobuf
# @see: https://github.com/anjmao/go2proto
go2proto -f ./protoDir -p "git.org.tech/scope/pkg"
````

````sh
Content-Type: application/grpc+json
Content-Type: application/grpc+proto
````

gRPC - open source, language & platform neutral RPC framework, based on Protocol Buffers.

Provides features such as:
* authentication.
* cancellation.
* timeouts.
* unary and bidirectional streaming and flow control.
* blocking or nonblocking bindings.

gRPC uses HTTP/2 for transport.

RPC based on actons & functions (for example REST based on resources).

By default, gRPC limits incoming messages to 4 MB.

Kinds of service method:
* Unary RPCs
  `rpc SayHello(HelloRequest) returns (HelloResponse){}`.
* Server streaming
  `rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}`.
* Client streaming
  `rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse){}`.
* Bidirectional streaming
  `rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){}`.
