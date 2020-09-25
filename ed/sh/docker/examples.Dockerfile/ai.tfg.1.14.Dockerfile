FROM cn007b/python:3.7

MAINTAINER V. Kovpak <cn007b@gmail.com>

RUN pip3 install numpy matplotlib tensorflow-gpu==1.14 tensorboard==1.14

RUN mkdir -p /app
ADD \
  https://raw.githubusercontent.com/cn007b/my/master/ed/ai/ml/tensorflow/examples/regression.py \
  /app/regression.py
