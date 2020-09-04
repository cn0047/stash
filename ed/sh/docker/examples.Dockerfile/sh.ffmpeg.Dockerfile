FROM python:3.7.6-slim-stretch

ENV PYTHONPATH $PYTHONPATH:.

RUN  apt-get -y update && \
     apt-get -y install ssh git build-essential wget curl

RUN apt-get -y update && \
    apt-get -y install --no-install-recommends autoconf automake libass-dev libfreetype6-dev \
    libsdl2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev \
    libxcb-xfixes0-dev pkg-config texinfo wget zlib1g-dev yasm \
    libx264-dev libx265-dev && \
    rm -rf /var/lib/apt/lists/*

# ffmpeg with x264
RUN DIR=$(mktemp -d) && cd ${DIR} && \
    curl -s http://ffmpeg.org/releases/ffmpeg-3.2.14.tar.gz | tar zxvf - -C . && \
    cd ffmpeg-3.2.14 && \
    ./configure \
    --enable-version3 --enable-gpl --enable-nonfree --enable-pic --enable-shared --enable-small --enable-libx264 \
    --enable-libx265 --enable-libtheora --enable-libvorbis --enable-libass --enable-postproc --enable-avresample \
    --enable-libfreetype --disable-debug && \
    make -j"$(nproc)" && \
    make install && \
    make distclean && \
    rm -rf ${DIR}

RUN echo "include /usr/local/lib" >> /etc/ld.so.conf && ldconfig
