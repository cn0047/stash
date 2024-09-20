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

Basic auth:
````go
"crypto/tls"
"google.golang.org/grpc"
"google.golang.org/grpc/credentials"
"google.golang.org/grpc/metadata"

// client
creds, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
auth := basicAuth{username: "admin", password: "admin"}
opts := []grpc.DialOption{
  grpc.WithPerRPCCredentials(auth),
  grpc.WithTransportCredentials(creds),
}
conn, err := grpc.Dial(address, opts...)

// server
cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
opts := []grpc.ServerOption{
  grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
  grpc.UnaryInterceptor(ensureValidBasicCredentials),
}
func ensureValidBasicCredentials(
  ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !valid(md["authorization"]) {
    return nil, errInvalidToken
  }
  return handler(ctx, req)
}
func valid(authorization []string) bool {
  if len(authorization) < 1 {
    return false
  }
  token := strings.TrimPrefix(authorization[0], "Basic ")
  return token == base64.StdEncoding.EncodeToString([]byte("admin:admin"))
}
````

OAuth 2.0:
````go
"google.golang.org/grpc"
"google.golang.org/grpc/credentials"
"google.golang.org/grpc/credentials/oauth"

// client
auth := oauth.NewOauthAccess(fetchToken())
creds, err := credentials.NewClientTLSFromFile(crtFile, hostname)
opts := []grpc.DialOption{
  grpc.WithPerRPCCredentials(auth),
  grpc.WithTransportCredentials(creds),
}
conn, err := grpc.Dial(address, opts...)
func fetchToken() *oauth2.Token {
  return &oauth2.Token{AccessToken: "some-secret-token"}
}

// server
cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
opts := []grpc.ServerOption{
  grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
  grpc.UnaryInterceptor(ensureValidToken),
}
func valid(authorization []string) bool {
  if len(authorization) < 1 {
    return false
  }
  token := strings.TrimPrefix(authorization[0], "Bearer ")
  return token == "some-secret-token"
}
func ensureValidToken(
  ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
  md, ok := metadata.FromIncomingContext(ctx)
  if !valid(md["authorization"]) {
    return nil, errInvalidToken
  }
  return handler(ctx, req)
}
func valid(authorization []string) bool {
  if len(authorization) < 1 {
    return false
  }
  token := strings.TrimPrefix(authorization[0], "Bearer ")
  return token == "some-secret-token"
}
````

JWT auth:
````go
// client
jwtCreds, err := oauth.NewJWTAccessFromFile("token.json")
creds, err := credentials.NewClientTLSFromFile("server.crt", "localhost")
opts := []grpc.DialOption{
  grpc.WithPerRPCCredentials(jwtCreds),
  grpc.WithTransportCredentials(creds),
}
conn, err := grpc.Dial(address, opts...)

// or
perRPC, err := oauth.NewServiceAccountFromFile("service-account.json", scope)
pool, _ := x509.SystemCertPool()
creds := credentials.NewClientTLSFromCert(pool, "")
opts := []grpc.DialOption{
  grpc.WithPerRPCCredentials(perRPC),
  grpc.WithTransportCredentials(creds),
}
conn, err := grpc.Dial(address, opts...)

// server
@TODO
````
