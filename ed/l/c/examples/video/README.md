README
-

````sh
docker build -t ubuntuvideo -f ed/sh/docker/examples.Dockerfile/c.video.Dockerfile .
docker run -ti --rm --net=xnet -v $HOME/Downloads:/d -w /d -v $PWD:/gh ubuntuvideo /bin/bash

f=/gh/ed/l/c/examples/video/video.2.c
gcc -L/opt/ffmpeg/lib -I/opt/ffmpeg/include/ $f \
  -lavcodec -lavformat -lavfilter -lavdevice -lswresample -lswscale -lavutil -o x
./x v1.mp4
````
