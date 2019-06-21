System Design
-

Replication needed for high availability (in case node die - you'll have all data).

1: Clarify constraints and use cases.
2: Abstract design.
3: Understanding bottlenecks.
4: Scaling your abstract design.

Primary concerns:
* reliability (fault tolerant)
* scalability (increasing load)
* maintainability (code that can easily be understood, refactored and upgraded)

Elements of a System:
* architecture
* modules
* components
* interfaces
* data

CAP Theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following:
* Consistency
* Availability
* Partition Tolerance
