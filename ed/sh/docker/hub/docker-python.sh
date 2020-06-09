docker-python
-

# python
version=3.6
version=3.7
version=3.7-ml
docker build -t cn007b/python:$version ./docker/$version
# check
docker run -it --rm cn007b/python:$version python3 --version
docker run -it --rm cn007b/python:$version python3.7 --version
docker run -it --rm cn007b/python:$version pip3 --version
# push
docker push cn007b/python:$version

# latest
docker tag cn007b/python:$version cn007b/python:latest
docker push cn007b/python:latest
