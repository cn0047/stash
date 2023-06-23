Consul
-

[docs](https://developer.hashicorp.com/consul/docs)
[glossary](https://developer.hashicorp.com/consul/docs/install/glossary)

````sh
consul version
consul validate $f
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

````sh
# consul-server.mesh.hcl

node_name = "consul-server"
server    = true
bootstrap = true
ui_config {
  enabled = true
}
datacenter = "dc1"
data_dir   = "consul/data"
log_level  = "INFO"
addresses {
  http = "0.0.0.0"
}
connect {
  enabled = true
}
````
