NATS
-

[docs](https://docs.nats.io/)

NATS - high performance message system and connective fabric.

NATS Core.
JetStream.

Service.
Stream.

````sh
nats server run
nats server run --jetstream

nats context select dev
nats reply hello.bob "Hi"
nats req hello.bob ""

nats sub hello.world
nats sub hello.*.any
nats sub hello.>
nats sub >           # all
nats pub hello.world "Hey"

nats stream ls -a
nats stream add

nats consumer create
nats consumer info
nats consumer next orders pull_consumer

nats kv add mybucket
nats kv put mybucket foo bar
````
