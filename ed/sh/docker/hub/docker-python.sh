docker-python
-

# python
version=3.6
docker build -t cn007b/python:$version ./docker/$version
# check
docker run -it --rm cn007b/python:$version python3 --version
docker run -it --rm cn007b/python:$version pip3 --version
# push
docker push cn007b/python:$version

# latest
docker build -t cn007b/python:latest ./docker/$version
docker push cn007b/python:latest
