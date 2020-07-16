FROM nvidia/cuda:10.0-cudnn7-runtime

Maintainer V. Kovpak <cn007b@gmail.com>

# ENV LD_LIBRARY_PATH=/usr/local/cuda-10.0/lib64
ENV LD_LIBRARY_PATH=/usr/local/cuda-10.0/lib64:/usr/local/cuda-10.0/compat/

RUN apt-get update --fix-missing
RUN apt-get install -y software-properties-common build-essential pkg-config wget
RUN add-apt-repository ppa:deadsnakes/ppa
RUN apt-get update

ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8
RUN apt-get install -y python3.7 python3.7-dev
RUN wget https://bootstrap.pypa.io/get-pip.py
RUN python3 get-pip.py
RUN \
     ln -sf /usr/bin/python3.7 /usr/local/bin/python3 \
  && ln -sf /usr/local/bin/pip /usr/local/bin/pip3 \
  && ln -s $(which python3.7) /usr/local/bin/python

RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/*

RUN pip3 --no-cache-dir install --upgrade pip setuptools
