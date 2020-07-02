FROM cn007b/python

MAINTAINER V. Kovpak <cn007b@gmail.com>

RUN pip3 install numpy matplotlib tensorflow tensorboard

RUN mkdir -p /app
ADD \
  https://raw.githubusercontent.com/cn007b/my/master/ed/ai/ml/tensorflow/examples/regression.py \
  /app/regression.py
