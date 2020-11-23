Service Meshe
-

istio
envoy
traefik
consul
k8s

The idea behind service meshe - is to add sidecar proxy
to you microservice to help with:
service discovery, load balancing, rate limiting, circuit breake, retry, etc.

## Service meshe infrastructure:

Control plane:
* Authentication
* Routing
* Retry
* Timeout
* Circuit Breaker

````
pilot, mixer, cit
````
