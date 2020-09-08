# python

# profiling
pip3 install pyprof2calltree
python3 ed/l/python/examples/http/server.prof.py
curl localhost:8080
#
f=ed/l/python/examples/whatever/hw.py
python3 -m cProfile -s cumtime $f
#
python3 -m cProfile -o callgrind.cprof $f
pyprof2calltree -k -i callgrind.cprof

# blur
docker run -it --rm -v $PWD:/gh -w /gh/ed/l/python/examples/blur xpy sh -c '
  pip3 install opencv-python
  python3 bench.py
'

# blur 2
cd ed/l/python/examples/blur/
virtualenv -p python3.7 local && source ./local/bin/activate
pip3 install \
  numpy==1.17.4 \
  lap==0.4.0 \
  filterpy==1.4.5 \
  opencv-python==4.4.0.42
python3 blur.2.py

# desktop
brew install pyqt5 qt-creator
python3 ed/l/python/examples/desktop/one.py

# http
python3 ed/l/python/examples/http/server.simple.py
curl localhost:8080
curl -X POST localhost:8080 -d v=9



#### Python Flask

# one
cd ed/l/python/python.flask/examples/one
virtualenv -p python3.7 local && source ./local/bin/activate
pip3 install flask flask_api flask_cors
python3 main.py
curl -X GET localhost:8080/
curl -X GET localhost:8080/x
curl -X POST localhost:8080/x
