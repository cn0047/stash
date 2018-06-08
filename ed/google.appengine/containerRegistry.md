Container Registry
-

````
# init
gcloud auth configure-docker

# build & check
docker build -t cn007b/ubuntu-gcloud https://raw.githubusercontent.com/cn007b/docker-ubuntu/master/docker/17.10/Dockerfile
docker run -ti --rm cn007b/ubuntu-gcloud echo "OK"

# push
docker tag cn007b/ubuntu-gcloud gcr.io/clique-dev/kovpak-test:latest
docker push gcr.io/clique-dev/kovpak-test:latest
# open
open http://gcr.io/clique-dev/kovpak-test

# pull
docker pull gcr.io/clique-dev/kovpak-test:latest
````
