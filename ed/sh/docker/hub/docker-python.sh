docker-python
-

# python
version=3.6
version=3.7
version=3.7-ml
version=3.8
version=3.8-tensorboard-2.4.0
docker build -t cn007b/python:$version ./docker/$version
# check
docker run -it --rm cn007b/python:$version /bin/bash
docker run -it --rm cn007b/python:$version python --version
docker run -it --rm cn007b/python:$version python3 --version
docker run -it --rm cn007b/python:$version python3.7 --version
docker run -it --rm cn007b/python:$version python3.8 --version
docker run -it --rm cn007b/python:$version pip --version
docker run -it --rm cn007b/python:$version pip3 --version
# push
docker push cn007b/python:$version

# ml
declare -a arr=('numpy' 'matplotlib' 'torch' 'torchvision' 'tensorflow' 'tensorboard' 'kubernetes' 'kfp')
for k in "${arr[@]}"; do
  echo -e "\n=======\n $k \n=======\n"
  docker run -it --rm cn007b/python:$version sh -c "python3 -c 'import "$k"; print("$k".__version__)'"
done

# latest
docker tag cn007b/python:$version cn007b/python:latest
docker push cn007b/python:latest
