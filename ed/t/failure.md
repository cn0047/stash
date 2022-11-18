Failure
-

Fault - when one component of the system deviating from its spec.
Failure - when the system as a whole stops providing the required service to the user.
Generally prefer tolerating faults over preventing faults.
For example, hardware faults:
disks may be set up in a RAID configuration,
servers may have dual power supplies and hot-swappable CPUs,
and datacenters may have batteries and diesel generators for backup power...

#### Cascading Failures

Conditions for Cascading Failures:
* Process death (crash).
* Process updates.
* New rollouts.
* Organic growth.

Steps to Address Cascading Failures:
* Increase resources.
* Restart servers.
* Drop traffic.
* Enter degraded mode.
* Eliminate bad traffic.
