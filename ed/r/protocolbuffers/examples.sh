# protoc examples

# one
cd ed/r/protocolbuffers/examples/one
protoc --gofast_out=plugins=grpc:. job.proto
protoc --js_out=import_style=commonjs:test \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:test job.proto
