Service Meshe
-

````
istio
envoy
traefik
consul
k8s
````

The idea behind service meshe - is to add sidecar proxy to you microservice to help with:
* Service discovery.
* Load balancing.
* Rate limiting.
* Circuit breake.
* Retry.

## Service meshe infrastructure:

Control plane:
* Authentication.
* Routing.
* Retry.
* Timeout.
* Circuit Breaker.

````
pilot, mixer, cit
````
