# web.three.tiny

docker run -it --rm -v $PWD/ed/go/examples/web.three.tiny:/app -e GOPATH='/app' \
  cn007b/go sh -c 'cd $GOPATH/src/app && go install'

docker run -it --rm -v $PWD/ed/go/examples/web.three.tiny:/app -e GOPATH='/app' \
  -p 8080:8080 \
  cn007b/go sh -c 'cd $GOPATH && ./bin/app'

# check
curl http://localhost:8080/v1/file-info/id/7
