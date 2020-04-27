FROM alpine:3.8 as compiler

RUN echo -e '@edgunity http://nl.alpinelinux.org/alpine/edge/community \
    @edge http://nl.alpinelinux.org/alpine/edge/main \
    @testing http://nl.alpinelinux.org/alpine/edge/testing \
    @community http://dl-cdn.alpinelinux.org/alpine/edge/community' \
    >> /etc/apk/repositories

RUN apk add --update --no-cache \
      build-base \
      openblas-dev \
      unzip \
      wget \
      cmake \
      g++ \
      libjpeg  \
      libjpeg-turbo-dev \
      libpng-dev \
      jasper-dev \
      tiff-dev \
      libwebp-dev \
      clang-dev \
      linux-headers

ENV CC /usr/bin/clang
ENV CXX /usr/bin/g++
ENV OPENCV_VERSION='3.4.2' DEBIAN_FRONTEND=noninteractive

RUN mkdir /opt && cd /opt && \
  wget https://github.com/opencv/opencv/archive/${OPENCV_VERSION}.zip && \
  unzip ${OPENCV_VERSION}.zip && \
  rm -rf ${OPENCV_VERSION}.zip

RUN mkdir -p /opt/opencv-${OPENCV_VERSION}/build && \
  cd /opt/opencv-${OPENCV_VERSION}/build && \
  cmake \
    -D BUILD_DOCS=OFF \
    -D BUILD_EXAMPLES=OFF \
    -D BUILD_opencv_apps=OFF \
    -D BUILD_opencv_python2=OFF \
    -D BUILD_opencv_python3=OFF \
    -D BUILD_PERF_TESTS=OFF \
    -D BUILD_SHARED_LIBS=OFF \
    -D BUILD_TESTS=OFF \
    -D CMAKE_BUILD_TYPE=RELEASE \
    -D ENABLE_PRECOMPILED_HEADERS=OFF \
    -D FORCE_VTK=OFF \
    -D WITH_FFMPEG=OFF \
    -D WITH_GDAL=OFF \
    -D WITH_IPP=OFF \
    -D WITH_OPENEXR=OFF \
    -D WITH_OPENGL=OFF \
    -D WITH_QT=OFF \
    -D WITH_TBB=OFF \
    -D WITH_XINE=OFF \
    -D BUILD_JPEG=ON  \
    -D BUILD_TIFF=ON \
    -D BUILD_PNG=ON \
  .. && \
  make -j$(nproc) && \
  make install && \
  rm -rf /opt/opencv-${OPENCV_VERSION}

# RUN wget --progress=dot:giga https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.0.0-linux-x86-64.tar.gz && \
#     pwd && \
#     tar -xzf libwebp-1.0.0-linux-x86-64.tar.gz && \
#     mv /libwebp-1.0.0-linux-x86-64/lib/libwebp.a /usr/lib && \
#     rm -rf /libwebp*

# RUN wget --progress=dot:giga http://www.ece.uvic.ca/~frodo/jasper/software/jasper-2.0.10.tar.gz && \
#     tar -xzf jasper-2.0.10.tar.gz && \
#     cd jasper-2.0.10 && \
#     mkdir BUILD && \
#     cd BUILD && \
#     cmake -DCMAKE_INSTALL_PREFIX=/usr \
#       -DCMAKE_BUILD_TYPE=Release \
#       -DCMAKE_SKIP_INSTALL_RPATH=YES \
#       -DCMAKE_INSTALL_DOCDIR=/usr/share/doc/jasper-2.0.10 \
#       -DJAS_ENABLE_SHARED=FALSE \
#       ..  && \
#     make install && \
#     rm -rf /jasper-2.0.10*

ENV PKG_CONFIG_PATH=/usr/local/lib64/pkgconfig:/usr/lib/pkgconfig

# COPY app.cpp app.cpp

# RUN g++ -Wl,-Bstatic -static-libgcc -std=c++11 \
#     app.cpp \
#     -o /app \
#     $(pkg-config --cflags --libs -static opencv) \
#     -lgfortran -lquadmath

FROM alpine
# COPY --from=compiler /app /bin/app
