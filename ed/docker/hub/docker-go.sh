docker-go
-

# go
docker build -t cn007b/go:1.10 ./docker/1.10
# check
docker run -it --rm cn007b/go:1.10 go version
docker run -it --rm cn007b/go:1.10 golint --help
docker run -it --rm cn007b/go:1.10 gometalinter --help
docker run -it --rm cn007b/go:1.10 goveralls --help
docker run -it --rm cn007b/go:1.10 dlv version
docker run -it --rm cn007b/go:1.10 gin --help
docker run -it --rm cn007b/go:1.10 pprof --help
docker run -it --rm cn007b/go:1.10 dot -V
# push
docker push cn007b/go:1.10

# gae
docker build -t cn007b/go:1.10-gae ./docker/1.10-gae
# check
docker run -it --rm cn007b/go:1.10-gae gcloud version
docker run -it --rm cn007b/go:1.10-gae goapp version
# push
docker push cn007b/go:1.10-gae

# latest
docker build -t cn007b/go:latest ./docker/1.10-gae
docker push cn007b/go:latest
