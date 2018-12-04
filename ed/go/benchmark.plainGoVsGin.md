Plain Go vs Gin
-

The aim of this benchmark is to answer the question
do you need framework for your tiny microservice.

#### Prepare

Please run next command with purpose to make go web-server available
from one docker container to another docker container and especially for `apachebench`.

````bash
# OSX
sudo ifconfig lo0 alias 10.254.254.254

# Ubuntu
sudo ifconfig lo 10.254.254.254
````

#### Prepare Gin

````bash
mkdir -p /tmp/benchmark/src/gin

curl -o /tmp/benchmark/src/gin/main.go \
  "https://raw.githubusercontent.com/cn007b/my/master/ed/go.gin/examples/one/src/one/main.go"

docker run -it --rm -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/gin-gonic/gin && go install gin'
````

#### Prepare Plain Go

````bash
mkdir -p /tmp/benchmark/src/plain

curl -o /tmp/benchmark/src/plain/main.go \
  "https://raw.githubusercontent.com/cn007b/my/master/ed/go/examples/web.three.tiny/src/app/main.go"

docker run -it --rm -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c 'cd $GOPATH && go install plain'
````

#### Benchmark

````bash
# run gin
docker run -it --rm -p 8080:8080 -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin'

# run apachebench
docker run -ti --rm cn007b/ubuntu ab -k -n 5000 -c 100 -t 2 "http://10.254.254.254:8080/v1/file-info/id/7"

# run plain go
docker run -it --rm -p 8080:8080 -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c 'cd $GOPATH && ./bin/plain'

# run apachebench
docker run -ti --rm cn007b/ubuntu ab -k -n 5000 -c 100 -t 2 "http://10.254.254.254:8080/v1/file-info/id/7"
````

#### Result Gin

````
Finished 3088 requests


Server Software:
Server Hostname:        10.254.254.254
Server Port:            8080

Document Path:          /v1/file-info/id/7
Document Length:        10 bytes

Concurrency Level:      100
Time taken for tests:   2.002 seconds
Complete requests:      3088 ‼️
Failed requests:        0
Keep-Alive requests:    3088
Total transferred:      484816 bytes
HTML transferred:       30880 bytes
Requests per second:    1542.76 [#/sec] (mean) ‼️
Time per request:       64.819 [ms] (mean)
Time per request:       0.648 [ms] (mean, across all concurrent requests)
Transfer rate:          236.54 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2  17.4      0     145
Processing:     3   60  34.8     54     288
Waiting:        3   60  34.8     54     288
Total:          3   63  42.3     54     350

Percentage of the requests served within a certain time (ms)
  50%     54
  66%     67
  75%     78
  80%     84
  90%    108
  95%    128
  98%    190
  99%    273
 100%    350 (longest request)
````

#### Result Plain Go

````
Finished 7493 requests


Server Software:
Server Hostname:        10.254.254.254
Server Port:            8080

Document Path:          /v1/file-info/id/7
Document Length:        11 bytes

Concurrency Level:      100
Time taken for tests:   2.001 seconds
Complete requests:      7493 ‼️
Failed requests:        0
Keep-Alive requests:    7493
Total transferred:      1138936 bytes
HTML transferred:       82423 bytes
Requests per second:    3744.57 [#/sec] (mean) ‼️
Time per request:       26.705 [ms] (mean)
Time per request:       0.267 [ms] (mean, across all concurrent requests)
Transfer rate:          555.83 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1  10.8      0     151
Processing:     3   26   9.4     25      92
Waiting:        3   25   9.4     25      92
Total:          3   26  15.6     25     228

Percentage of the requests served within a certain time (ms)
  50%     25
  66%     29
  75%     31
  80%     33
  90%     38
  95%     42
  98%     51
  99%     58
 100%    228 (longest request)
````

#### Conclusion

100% sure you've answered the main question of this benchmark! 🙂
