Distributed Systems
-

Make all nodes to behave idempotent as much as possible,
because some actions may repeat twice:
* DB incremented value but network response failed.
* Retry due to timeout, but node processed request it just busy handling spikes.

**Gosip broadcast algorithm** - broadcast message only to 3-5 neighbors.

**Failure detector** - like health check,
checks that node response with expected message during expected time.

**Eventual failure detector** - like failure detector,
but aware that network not reliable: timeout may happen or delayed response, etc.
