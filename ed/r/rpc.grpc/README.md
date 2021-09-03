gRPC
-

[grpc](https://grpc.io/docs)

````sh
Content-Type: application/grpc+json
Content-Type: application/grpc+proto
````

gRPC - open source RPC framework.

RPC based on actons & functions (for example REST based on resources).

Kinds of service method:
* Unary RPCs
  `rpc SayHello(HelloRequest) returns (HelloResponse){}`.
* Server streaming
  `rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){}`.
* Client streaming
  `rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse){}`.
* Bidirectional streaming
  `rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){}`.
