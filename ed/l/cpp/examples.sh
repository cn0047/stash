# c++

gh=$HOME/web/kovpak/gh
d=$gh/ed/l/cpp/examples

f=$d/whatever/hw.cpp
f=$d/whatever/exception.cpp
g++ -w -o x $f && ./x

# include
g++ -w -o x $d/whatever/1.lib.cpp $d/whatever/include.cpp && ./x
g++ -w -o x $d/whatever/1.lib.cpp $d/whatever/include.h.cpp && ./x

# class.one
g++ -w -o x $d/class.one/a.cpp $d/class.one/t.cpp $d/class.one/main.cpp && ./x



# blur
docker run -it --rm -v $PWD:/gh -w /gh spmallick/opencv-docker:opencv sh -c '
  f=ed/l/cpp/examples/blur/bench.cpp
  g++ -w $f -o x `pkg-config --cflags --libs opencv`
  ./x
'
# and
docker run -it --rm -v $gh:/gh -w /gh spmallick/opencv-docker:opencv sh -c '
  f=ed/l/cpp/examples/blur/one.cpp
  g++ -w $f -o x `pkg-config --cflags --libs opencv`
  ./x /gh/ed/l/python/examples/blur/z.png
'
