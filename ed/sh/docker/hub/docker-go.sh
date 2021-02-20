docker-go
-

/opt/cn007b/go/ # reserved
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# go
version=1.13
version=1.13-neo4j
version=1.14
version=1.14-alpine
version=1.15
version=1.16-protobuf
docker build -t cn007b/go:$version ./docker/$version
# check
docker run -it --rm cn007b/go:$version sh -c 'hash go'
docker run -it --rm cn007b/go:$version sh -c 'go version'
docker run -it --rm cn007b/go:$version sh -c 'hash dep'
docker run -it --rm cn007b/go:$version sh -c 'hash golint'
docker run -it --rm cn007b/go:$version sh -c 'hash golangci-lint'
docker run -it --rm cn007b/go:$version sh -c 'hash goveralls'
docker run -it --rm cn007b/go:$version sh -c 'hash dlv'
docker run -it --rm cn007b/go:$version sh -c 'hash gin'
docker run -it --rm cn007b/go:$version sh -c 'hash gosec'
docker run -it --rm cn007b/go:$version sh -c 'hash pprof'
docker run -it --rm cn007b/go:$version sh -c 'hash dot'
# push
docker push cn007b/go:$version

# # gae
# docker build -t cn007b/go:$version-gae ./docker/$version-gae
# # check
# docker run -it --rm cn007b/go:$version-gae gcloud version
# docker run -it --rm cn007b/go:$version-gae goapp version
# # push
# docker push cn007b/go:1.10-gae

# # protobuf
# docker build -t cn007b/go:$version-protobuf ./docker/$version-protobuf
# # check
# docker run -it --rm cn007b/go:$version-protobuf protoc --version
# # push
# docker push cn007b/go:1.10-protobuf

# latest
docker tag cn007b/go:$version cn007b/go:latest
docker push cn007b/go:latest
