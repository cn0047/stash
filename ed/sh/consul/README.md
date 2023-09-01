Consul
-

[docs](https://developer.hashicorp.com/consul/docs)
[glossary](https://developer.hashicorp.com/consul/docs/install/glossary)
[services conf](https://developer.hashicorp.com/consul/docs/services/configuration/services-configuration-reference)
[service mesh conf](https://developer.hashicorp.com/consul/docs/connect/config-entries)
[api-gateway conf](https://developer.hashicorp.com/consul/docs/connect/gateways/api-gateway/configuration/api-gateway)
[http-route conf](https://developer.hashicorp.com/consul/docs/connect/gateways/api-gateway/configuration/http-route)
[tcp-route conf](https://developer.hashicorp.com/consul/docs/connect/gateways/api-gateway/configuration/tcp-route)
[agent conf](https://developer.hashicorp.com/consul/docs/agent/config)
[agent conf file](https://developer.hashicorp.com/consul/docs/agent/config/config-files)
[rate limiting](https://developer.hashicorp.com/consul/docs/agent/limits)

````sh
consul version
consul validate $f

consul agent -config-dir configs

consul services register -name=web services.hcl
````

````sh
curl -H "X-Consul-Token: $CONSUL_TOKEN"
````

Consul - tool for discovering and configuring services.

Consul offers:
* service discovery (client-side discovery, server-side discovery).
* service mesh.
* traffic management.
* automated updates to network infrastructure device.

Service discovery - uses service identity instead of IP (web.svc1.dc1.consul).

Service mesh - dedicated network layer that provides secure service-to-service communication.

Service mesh provides:
* service discovery.
* application health monitoring.
* load balancing.
* automatic failover.
* traffic management.
* encryption.
* observability and traceability.
* authentication and authorization.
* network automation.

Consul server: leader, follower.
Control plane: consul client.
Data plane: app.

LAN gossip pool - client and server agents participate in a LAN gossip pool.

Consul agent - core process of Consul, daemon on every member of Consul cluster,
running in client or server mode.

Raft used as consensus protocol in consul.

Proxy - enables applications to connect to other services in the service mesh.

Service intention - control traffic communication between services at the L4 or L7.

Gateway:
* mesh gateway - service-to-service traffic between Consul datacenters or between Consul admin partitions.
* ingress gateway - connectivity from services outside Consul service mesh to services in the mesh.
* terminating gateway - from services in the Consul service mesh to services outside the mesh.

API gateway - allows external network clients to access services running in Consul datacenter.
* Control access at the point of entry (with TLS certificates).
* Simplify traffic management (load balance requests).

Cluster peering - connect two or more independent Consul clusters.

Failover - route traffic to and from unhealthy or unreachable service.

Dataplane - manages Envoy proxies and leaves responsibility
for other functions (communications between service instances, their sidecar proxies, and the servers)
to the orchestrator (for example kubelet in k8s).

KV Store - simple Key/Value store.

ACLs - authenticate requests and authorize access to resources.
Rule -> Policy -> Token (PolicyA, PolicyB).
