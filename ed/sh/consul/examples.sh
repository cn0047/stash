# consul examples

# hw
go run ed/l/go/examples/http/http.server.hw2.go
go run ed/l/go/examples/http/http.client.hw.go
#
consul agent -dev
consul services register ed/sh/consul/examples/hw/service.hws.1.hcl
consul services register ed/sh/consul/examples/hw/service.hwc.1.hcl
consul services deregister ed/sh/consul/examples/hw/service.hws.1.hcl
consul services deregister ed/sh/consul/examples/hw/service.hwc.1.hcl
# or
consul agent -dev -config-dir=ed/sh/consul/examples/hw/service.hws.1.hcl -node=cnhws
consul agent -dev -config-dir=ed/sh/consul/examples/hw/service.hwc.1.hcl -node=cnhwc
# vet
consul members
consul catalog services
#
consul connect proxy -sidecar-for hw_1
consul connect proxy -sidecar-for hw_2
# test
curl localhost:8081
curl localhost:8082

curl http://127.0.0.1:8500/v1/agent/services

docker run -it --rm --net=xnet -p 8080:8080 -p 8070:8070 --name chws \
  -v $PWD/ed/sh/consul/examples/hw:/app/consul.conf \
  --hostname dhhws cn007b/pi:hws

docker run -it --rm --net=xnet -p 8081:8081 -p 8071:8071 --name chwc \
  -v $PWD/ed/sh/consul/examples/hw:/app/consul.conf \
  -e TARGET_URI='http://localhost:8060' cn007b/pi:hwc

docker exec -it chws sh -c '
  consul agent -config-file=/app/consul.conf/service.hws.1.hcl -data-dir=/app/consul.data -node=cnhws
'

docker exec -it chwc sh -c '
  consul agent -config-file=/app/consul.conf/service.hwc.1.hcl -data-dir=/app/consul.data -node=cnhwc
'
