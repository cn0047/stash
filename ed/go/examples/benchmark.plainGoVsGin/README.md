Plain Go vs Gin
-

The aim of this benchmark is to answer the question
do you need framework for your **tiny** microservice.

#### Prerequisites

Let's consider simple example where we need framework only to parse URL
<br>(someone may decide that it's sufficient reason to start use framework for tiny microservice).

For gin example will use [this](https://github.com/cn007b/my/blob/master/ed/go.gin/examples/one/src/one/main.go) implementation.
For plain go example will use [this](https://github.com/cn007b/my/blob/master/ed/go/examples/web.three.tiny/src/app/main.go) implementation.

Code looks in next way:

![code](https://gist.github.com/cn007b/6083c25407e1f19317f3a513e7ae2a28/raw/6a8b388cea838490d79a09eec71cb41c9ad4a775/code.png)

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
    cn007b/go sh -c 'go get github.com/gin-gonic/gin && go install gin'
````

#### Prepare Plain Go

````bash
mkdir -p /tmp/benchmark/src/plain

curl -o /tmp/benchmark/src/plain/main.go \
  "https://raw.githubusercontent.com/cn007b/my/master/ed/go/examples/web.three.tiny/src/app/main.go"

docker run -it --rm -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    cn007b/go sh -c 'go install plain'
````

#### Benchmark

````bash
# run gin
docker run -it --rm -p 8080:8080 -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    cn007b/go sh -c './bin/gin'

# run apachebench
docker run -ti --rm cn007b/ubuntu ab -k -n 5000 -c 100 -t 2 "http://10.254.254.254:8080/v1/id/7"

# run plain go
docker run -it --rm -p 8080:8080 -v /tmp/benchmark:/app -w /app -e GOPATH='/app' \
    cn007b/go sh -c './bin/plain'

# run apachebench
docker run -ti --rm cn007b/ubuntu ab -k -n 5000 -c 100 -t 2 "http://10.254.254.254:8080/v1/id/7"
````

#### Result Gin

````
Finished 7637 requests


Server Software:
Server Hostname:        10.254.254.254
Server Port:            8080

Document Path:          /v1/id/7
Document Length:        10 bytes

Concurrency Level:      100
Time taken for tests:   2.000 seconds
Complete requests:      7637 ‼️
Failed requests:        0
Keep-Alive requests:    7637
Total transferred:      1199009 bytes
HTML transferred:       76370 bytes
Requests per second:    3818.42 [#/sec] (mean) ‼️
Time per request:       26.189 [ms] (mean)
Time per request:       0.262 [ms] (mean, across all concurrent requests)
Transfer rate:          585.44 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   8.3      0     129
Processing:     1   25  15.1     24     155
Waiting:        1   25  15.1     24     155
Total:          1   26  17.3     24     196

Percentage of the requests served within a certain time (ms)
  50%     24
  66%     28
  75%     31
  80%     33
  90%     41
  95%     56
  98%     73
  99%     90
 100%    196 (longest request)
````

#### Result Plain Go

````
Finished 15349 requests


Server Software:
Server Hostname:        10.254.254.254
Server Port:            8080

Document Path:          /v1/id/7
Document Length:        11 bytes

Concurrency Level:      100
Time taken for tests:   2.000 seconds
Complete requests:      15349 ‼️
Failed requests:        0
Keep-Alive requests:    15349
Total transferred:      2333048 bytes
HTML transferred:       168839 bytes
Requests per second:    7672.79 [#/sec] (mean) ‼️
Time per request:       13.033 [ms] (mean)
Time per request:       0.130 [ms] (mean, across all concurrent requests)
Transfer rate:          1138.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   5.9      0     123
Processing:     1   13   2.9     12      49
Waiting:        1   13   2.9     12      49
Total:          1   13   6.7     12     142

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     13
  75%     14
  80%     15
  90%     16
  95%     17
  98%     20
  99%     23
 100%    142 (longest request)
````

#### Conclusion

100% sure you've answered the main question of this benchmark! 🙂
