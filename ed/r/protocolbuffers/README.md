Protocol Buffers (protobuf)
-

[docs](https://developers.google.com/protocol-buffers/)
[versions](https://github.com/protocolbuffers/protobuf/tags)
[programming guides](https://protobuf.dev/programming-guides/style/)
[size limit](https://protobuf.dev/programming-guides/encoding/#size-limit)
[buf](https://buf.build/)

````sh
protoc --proto_path=$pathToProtoFiles

cat foo.msg | protoc --encode=Package.Foo foo.proto

cat cmdmWindows.Query.msq | protoc --encode=cmdmWindows.Query cmdmWindows.proto
cat cmdmWindows.Query.msq \
  | protoc --encode=cmdmWindows.Query cmdmWindows.proto \
  | curl -k -X PUT --data-binary @- https://10.254.254.254:443/command/windows \
  | protoc --decode cmdmWindows.Response cmdmWindows.proto
````

Protocol buffers - language & platform neutral, extensible mechanism for serializing structured data.

Canonically, messages are serialized into a binary wire format.

Scalar Value Types:
* double.
* float.
* int32, int64.
* uint32, uint64 - Uses variable-length encoding (integer in php).
* sint32, sint64 - (integer in php).
* fixed32, fixed64 - (integer in php).
* sfixed32, sfixed64 - (integer in php).
* bool.
* string.
* bytes - any arbitrary sequence of bytes (string in php).

`repeated` keyword for arrays.

`enum` keyword for enumerations.
First enum value must have value zero and should have name `*_UNSPECIFIED` or `*_UNKNOWN`,
and should have no semantic meaning other than "this value was unspecified".

Protocol Buffers are not designed to handle large messages.
If you are dealing in messages larger than a megabyte each,
it may be time to consider an alternate strategy.

Size Limitations: protos must be smaller than 2 GiB when serialized.

````go
import "google/protobuf/any.proto";
import "google/protobuf/timestamp.proto";

import "proto/common.proto";

message Job {
  reserved 20;
  string id = 1;
  repeated string tags = 2;
  map<string, string> params = 3; // grpcwebtext will not generate seter for map
  google.protobuf.Timestamp created_at  = 5;

  oneof kind {
    UnitKind unit = 1;
    ModuleKind module = 2;
  }

  message UnitKind {}

  message ModuleKind {}
}
````
