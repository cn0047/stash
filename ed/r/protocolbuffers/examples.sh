# protoc examples

# one
cd ed/r/protocolbuffers/examples/one
protoc --gofast_out=plugins=grpc:. job.proto
protoc --js_out=import_style=commonjs:test \
    --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:test job.proto

# simple
cd ed/r/protocolbuffers/examples/simple
t=/tmp/proto.test.txt
# pack
cat event.v1.msg | protoc --encode=EventV1 event.v1.proto > $t
# unpack
cat $t | protoc --decode=EventV1 event.v1.proto
cat $t | protoc --decode=EventV2 event.v2.proto
cat $t | protoc --decode=EventV3 event.v3.proto
cat $t | protoc --decode=EventV4 event.v4.proto
