FROM cn007b/pi:ai.cuda.0

Maintainer V. Kovpak <cn007b@gmail.com>

RUN pip3 install tensorflow-gpu==1.14
RUN pip3 install numpy matplotlib
WORKDIR /tmp
RUN wget https://raw.githubusercontent.com/cn007b/my/master/ed/ai/ml/tensorflow/examples/regression.py
