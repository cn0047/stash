NATS
-

[docs](https://docs.nats.io/)
[REPL](https://github.com/nats-io/natscli)
[leaf node](https://docs.nats.io/running-a-nats-service/configuration/leafnodes)

NATS - high performance message system and connective fabric.

Core NATS:
* Publish-Subscribe (fan-out, one-to-many communication).
* Request-Reply.
* Queue Groups (fan-out + queue group - only one randomly chosen subscriber of the queue group will consume message).

Multi-Tenancy:
* JetStream accounts - use accounts to logically isolate different groups of users.
* Subject isolation - topics in one account are isolated from another.
* Resource management - each account can have own quotas.
* Authentication and authorization.

JetStream - streaming and persistence layer.
Stream - message store, which defines how messages are stored and limits (duration, size, interest).
Stream consumes normal NATS subject, any message published on subjects will be captured in the defined storage system.
Consumer - stateful view of a stream.
With message persistence + acknowledgment mechanisms, JetStream ensures messages are delivered at least once.
Durable consumer persists state, so it can resume processing even after restart.
Storage options: file-based, memory-based.
Headers - used for: de-duplication, auto-purging, metadata from republished messages, etc.

Subjects:

````sh
time.us.east
time.us.east.atlanta
time.eu.east
time.eu.east.warsaw

time.*.east    # ok for: time.us.east, time.eu.east
time.us.*      # ok for: time.us.east
time.New*.east # no substring match
time.us.>      # ok for: time.us.east, time.us.east.atlanta
*.*.east.>     # ok for: time.us.east.atlanta
````

````sh
# server

nats server run
nats server run --jetstream

nats server list
nats server info
nats server ping
nats server check connection
````

````sh
# client

h=localhost
p=53341
u=local
pas=''

nats --server $h:$p context add localTest --description 'localTest'
nats --server $h:$p context ls
nats --server $h:$p context select localTest

# key/value
nats --server $h:$p --user $u --password $pas kv add myBucket
nats --server $h:$p --user $u --password $pas kv put myBucket foo bar
nats --server $h:$p --user $u --password $pas kv get myBucket foo

# request/reply
# reply
nats --server $h:$p --user $u --password $pas reply reqReply.test.1 "Hello"
# request
nats --server $h:$p --user $u --password $pas req reqReply.test.1 "Hi"

# pub/sub
# subscribe
nats --server $h:$p --user $u --password $pas sub 'pubSub.test.simple'
nats --server $h:$p --user $u --password $pas sub 'pubSub.*.simple'
nats --server $h:$p --user $u --password $pas sub 'pubSub.*.*'
nats --server $h:$p --user $u --password $pas sub '*.test.>'
nats --server $h:$p --user $u --password $pas sub '>' # all
# publish
nats --server $h:$p --user $u --password $pas pub 'pubSub.test.simple' "ping 1"
nats --server $h:$p --user $u --password $pas pub 'pubSub.test.simple' "ping 2"
nats --server $h:$p --user $u --password $pas pub 'pubSub.test.simple' "ping 3"

# jetstream
s=s1 # stream name
sbj='jetStream.test.simple'
nats --server $h:$p --user $u --password $pas stream ls -a
nats --server $h:$p --user $u --password $pas stream add # interactive mode
nats --server $h:$p --user $u --password $pas stream add --replicas=3
nats --server $h:$p --user $u --password $pas stream info
nats --server $h:$p --user $u --password $pas stream report
nats --server $h:$p --user $u --password $pas stream rm
nats --server $h:$p --user $u --password $pas stream purge

# if create new stream for subject used in other stream:
# nats: error: could not create Stream: subjects overlap with an existing stream (10065)

# consumer
c=c1 # consumer name
nats --server $h:$p --user $u --password $pas consumer ls
nats --server $h:$p --user $u --password $pas consumer create # interactive mode
nats --server $h:$p --user $u --password $pas consumer info
nats --server $h:$p --user $u --password $pas consumer rm
# publish
nats --server $h:$p --user $u --password $pas pub $sbj "test 1"
nats --server $h:$p --user $u --password $pas pub $sbj "test 2"
nats --server $h:$p --user $u --password $pas pub $sbj "test 3"
# consume
nats --server $h:$p --user $u --password $pas consumer sub # interactive mode
nats --server $h:$p --user $u --password $pas consumer next $s $c --no-ack

# if create new consumer for subject used in other consumer: ok, no error.
````

## Leaf Nodes

Leaf - special type of server connection that bridges two separate NATS servers or clusters together.
Leaf Node extends existing NATS system, optionally bridging both operator and security domains.
Leaf Node transparently routes messages from local clients to one or more remote NATS system(s) and vice versa.

````sh
# leaf.conf
leafnodes {
  remotes = [{url: "nats://<main-server-host>:7422"}]
}

nats-server -c leaf.conf
````
