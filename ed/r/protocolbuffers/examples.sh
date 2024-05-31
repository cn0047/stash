# protoc examples

# one
cd ed/r/protocolbuffers/examples/one
protoc --gofast_out=plugins=grpc:. job.proto
protoc --js_out=import_style=commonjs:test \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:test job.proto

# go.grpc
go install github.com/gogo/protobuf/protoc-gen-gofast@latest
protoc \
  --gofast_out=plugins=grpc:. \
  --proto_path="ed/r/protocolbuffers/examples/go.grpc/" \
  ed/r/protocolbuffers/examples/go.grpc/main.proto
# or
cd ed/r/protocolbuffers/examples/go.grpc
protoc --gofast_out=plugins=grpc:. main.proto
