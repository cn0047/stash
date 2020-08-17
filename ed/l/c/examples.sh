# c

gh=$HOME/web/kovpak/gh
d=$gh/ed/l/c/examples

f=ed/l/c/examples/whatever/hw.c
f=ed/l/c/examples/whatever/array.one.c
f=ed/l/c/examples/whatever/func.one.c
f=ed/l/c/examples/whatever/enum.c
f=ed/l/c/examples/whatever/rand.c
f=ed/l/c/examples/whatever/rand.1.c
gcc -o /tmp/x -Wall $f && /tmp/x



# video
docker build -t ubuntuvideo -f $gh/ed/sh/docker/examples.Dockerfile/c.video.Dockerfile .
docker run -ti --rm --net=xnet -v $HOME/Downloads:/d -w /d -v $gh:/gh ubuntuvideo /bin/bash
# and
f=$d/video/video.2.c
gcc -L/opt/ffmpeg/lib -I/opt/ffmpeg/include/ $f \
  -lavcodec -lavformat -lavfilter -lavdevice -lswresample -lswscale -lavutil -o x
./x v1.mp4
