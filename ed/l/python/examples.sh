# blur
docker run -it --rm -v $PWD:/gh -w /gh/ed/l/python/examples/blur xpy sh -c '
  pip3 install opencv-python
  python3 bench.py
'

# desktop
brew install pyqt5 qt-creator
python3 ed/l/python/examples/desktop/one.py

# http
python3 ed/l/python/examples/http/server.simple.py
curl localhost:8080
curl -X POST localhost:8080 -d v=9
