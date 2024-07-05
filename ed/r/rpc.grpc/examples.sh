# grpc examples

# one
go install github.com/gogo/protobuf/protoc-gen-gofast@latest
protoc \
  --gofast_out=plugins=grpc:. \
  --proto_path="ed/r/rpc.grpc/examples/one/" \
  ed/r/rpc.grpc/examples/one/main.proto
# or
cd ed/r/protocolbuffers/examples/one
protoc --gofast_out=plugins=grpc:. main.proto
