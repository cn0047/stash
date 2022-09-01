System Design
-

````sh
# Time-based availability
availability = uptime / (uptime + downtime)

# Aggregate availability
availability = successful requests / total requests
````

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

System Design Plan:
1: Clarify constraints and use cases.
2: Abstract design.
3: Understand bottlenecks.
4: Scale your abstract design.

Storage Scalability:
* What is the amount of data that we need to store?
* Will the data keep growing over time? If yes, then at what rate?

Primary concerns:
* Reliability (fault tolerant).
* Scalability (increasing load).
* Maintainability (code that can easily be understood, refactored and upgraded).

CAP Theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following (pick 2out of 3):
* Consistency - data is the same across the cluster.
* Availability - ability to access cluster even if node goes down.
* Partition tolerance - cluster continues to function even if communication break between 2 nodes.

Elements of a System:
* Architecture.
* Modules.
* Components.
* Interfaces.
* Data.

Think about:
* DB (SQL, NoSQL), schema, indexes, ACID/BASE, replication, writes/reads.
* Cache, CDN.
* API, REST, gRPC, etc.
* LB.

#### SLA/SLO/SLI

Availability:
• What level of service will users expect?
• Does service tie directly to revenue?
• Is it paid service or free?
• What level of service do competitors provide?
• Is service for consumers or enterprises?

Service-level indicator (SLI) - measure of the service level provided by a service provider to a customer.
Common SLIs include latency, throughput, availability, error rate, durability, correctness, etc.

Service-level objective (SLO) - (key element of SLA) specific measurable characteristics
of the SLA such as availability, throughput, frequency, response time, or quality.
SLOs - `lower bound ≤ SLI ≤ upper bound`.
Publishing SLOs to users sets expectations about how service will perform.

Service-level agreement (SLA) - commitment (contract) between a service provider and a client:
quality, availability, responsibility, performance, costs, etc.
