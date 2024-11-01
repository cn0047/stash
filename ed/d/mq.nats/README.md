NATS
-

[docs](https://docs.nats.io/)

NATS - high performance message system and connective fabric.

NATS Core.
JetStream.

Service.
Stream.

Leaf Node extends existing NATS system, optionally bridging both operator and security domains.
Leaf Node transparently routes messages from local clients to one or more remote NATS system(s) and vice versa.

````sh
# server

nats server run
nats server run --jetstream
````

````sh
# client

nats context add localTest --description 'localTest'
nats context add localTest --server localhost:4222 --description 'localTest'
nats context ls
nats context select localTest

nats reply hello.bob "Hi"
nats req hello.bob ""

nats sub hello.world
nats sub hello.*.any
nats sub hello.>
nats sub >           # all
nats pub hello.world "Hey"

nats stream ls -a
nats stream add

nats consumer ls
nats consumer create
nats consumer info
nats consumer next orders pull_consumer

nats kv add mybucket
nats kv put mybucket foo bar
````

https://github.com/nats-io/natscli
