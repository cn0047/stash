Architecture
-

````sh
# Time-based availability
availability = uptime / (uptime + downtime)

# Aggregate availability
availability = successful requests / total requests
````

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

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

````sh
SLA=99.99 -> 100 fails for 1M requests
````

#### CAP

CAP theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following (pick 2 out of 3):
* Consistency - data is the same across the cluster (read from replica return data written on master).
* Availability - ability to access cluster even if node goes down (can read & write data).
* Partition tolerance - cluster continues to function even if communication break between 2 nodes (network partitioning).

CAP means: network partition happens - you have to make decision tolerate failure or not,
so you may have consistent system or available, but not both.
