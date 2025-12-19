Distributed Systems
-

Observability - not only monitoring cpu/mem/net/etc usage, but also:
* Logging.
* Tracing.
* Health checking.
* Distributed events monitoring.
* SLI metrics monitoring.

Consensus - getting all of the nodes to agree on something.

Make all nodes to behave idempotent as much as possible,
because some actions may repeat twice:
* DB incremented value but network response failed.
* Retry due to timeout, but node processed request it just busy handling spikes.

Dynamic coupling:
* Communication: sync, async.
* Consistency: atomic, eventual.
* Coordination: orchestration, choreography.

Fallacies of distributed systems:
* Network is reliable.
* Latency is zero.
* Bandwidth is infinite.
* Network is secure.
* Topology doesn't change.
* There is only one administrator (devops, infrastructure engineer).
* Transport cost is zero (network data send price).
* Network is homogeneous (systems, networks, etc).

**Gosip broadcast algorithm** - broadcast message only to 3-5 neighbors.

**Fault** - when one component of the system deviating from its spec.
**Failure** - when the system as a whole stops providing the required service to the user.
Generally prefer tolerating faults over preventing faults.
For example, hardware faults:
disks may be set up in a RAID configuration,
servers may have dual power supplies and hot-swappable CPUs,
and datacenters may have batteries and diesel generators for backup power...

**Failure detector** - like health check,
checks that node response with expected message during expected time.

**Eventual failure detector** - like failure detector,
but aware that network not reliable: timeout may happen or delayed response, etc.

Conditions for Cascading Failures:
* Process death (crash).
* Process updates.
* New rollouts.
* Organic growth.

Steps to address Cascading Failures:
* Increase resources.
* Restart servers.
* Drop traffic.
* Enter degraded mode.
* Eliminate bad traffic.
