FROM nvidia/cuda:10.0-cudnn7-runtime

Maintainer V. Kovpak <cn007b@gmail.com>

RUN apt-get update --fix-missing
RUN apt-get install -y software-properties-common build-essential wget
RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get update

RUN apt-get install -y python3.7 python3.7-dev \
  && ln -sf /usr/bin/python3.7 /usr/local/bin/python3 \
  && ln -sf /usr/local/bin/pip /usr/local/bin/pip3 \
  && ln -s $(which python3.7) /usr/local/bin/python
RUN wget https://bootstrap.pypa.io/get-pip.py
RUN python3 get-pip.py
RUN pip3 install tensorflow-gpu==1.14
# RUN pip3 install tensorflow
RUN apt-get clean
ENV LD_LIBRARY_PATH=/usr/local/cuda-10.0/lib64:/usr/local/cuda-10.0/compat/
