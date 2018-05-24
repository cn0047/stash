docker-go
-

# go
docker build -t cn007b/go:1.10 ./docker/1.10
# check
docker run -it --rm cn007b/go:1.10 go version
docker run -it --rm cn007b/go:1.10 golint --help
docker run -it --rm cn007b/go:1.10 dlv version
docker run -it --rm cn007b/go:1.10 pprof
# push
docker push cn007b/go:1.10

# latest
docker build -t cn007b/go:latest ./docker/1.10
docker push cn007b/go:latest
