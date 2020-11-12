# c

gh=$HOME/web/kovpak/gh
d=$gh/ed/l/c/examples

objdump -d /tmp/x

cr() {
    gcc -o /tmp/x -Wall $1 && /tmp/x
}

cr ed/l/c/examples/whatever/assert.c
cr ed/l/c/examples/whatever/hw.c
cr ed/l/c/examples/whatever/func.one.c
cr ed/l/c/examples/whatever/func.variadic.c
cr ed/l/c/examples/whatever/enum.c
cr ed/l/c/examples/whatever/rand.c
cr ed/l/c/examples/whatever/rand.1.c
cr ed/l/c/examples/whatever/stdin.c
cr ed/l/c/examples/whatever/for.c
cr ed/l/c/examples/whatever/recursion.c
cr ed/l/c/examples/whatever/digitsSum.c
cr ed/l/c/examples/whatever/char.c
cr ed/l/c/examples/whatever/array.one.c
cr ed/l/c/examples/whatever/array.permute.c
cr ed/l/c/examples/whatever/str.c
cr ed/l/c/examples/whatever/str.sort.c
cr ed/l/c/examples/whatever/print.pattern.c
cr ed/l/c/examples/whatever/struct.one.c
cr ed/l/c/examples/whatever/struct.two.c

# in docker
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh xubuntu /bin/bash
# ⬆

# video
docker build -t ubuntuvideo -f $gh/ed/sh/docker/examples.Dockerfile/c.video.Dockerfile .
docker run -ti --rm --net=xnet -v $HOME/Downloads:/d -w /d -v $gh:/gh ubuntuvideo /bin/bash
# and
f=$d/video/video.2.c
gcc -L/opt/ffmpeg/lib -I/opt/ffmpeg/include/ $f \
  -lavcodec -lavformat -lavfilter -lavdevice -lswresample -lswscale -lavutil -o x
./x v1.mp4
