docker-go
-

# go
version=1.11
docker build -t cn007b/go:$version ./docker/$version
# check
docker run -it --rm cn007b/go:$version go version
docker run -it --rm cn007b/go:$version golint --help
docker run -it --rm cn007b/go:$version gometalinter --help
docker run -it --rm cn007b/go:$version goveralls --help
docker run -it --rm cn007b/go:$version dlv version
docker run -it --rm cn007b/go:$version gin --help
docker run -it --rm cn007b/go:$version pprof --help
docker run -it --rm cn007b/go:$version dot -V
# push
docker push cn007b/go:$version

# gae
docker build -t cn007b/go:$version-gae ./docker/$version-gae
# check
docker run -it --rm cn007b/go:$version-gae gcloud version
docker run -it --rm cn007b/go:$version-gae goapp version
# push
docker push cn007b/go:1.10-gae

# protobuf
docker build -t cn007b/go:$version-protobuf ./docker/$version-protobuf
# check
docker run -it --rm cn007b/go:$version-protobuf protoc --version
# push
docker push cn007b/go:1.10-protobuf

# latest
docker build -t cn007b/go:latest ./docker/1.10-gae
docker push cn007b/go:latest
