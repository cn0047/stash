Architecture
-

````sh
# Time-based availability
availability = uptime / (uptime + downtime)

# Aggregate availability
availability = successful requests / total requests
````

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

Architectural plan:
1) Introduction and goals (fundamental requirements).
2) Constraints.
3) Context and scope (external systems and interfaces).
4) Solution strategy (core ideas and approaches).
5) Building block view (structure of source code modularisation).
6) Runtime view (important runtime scenarios).
7) Deployment view (hardware and infrastructure).
8) Crosscutting concepts.
9) Architectural decisions.
10) Quality requirements.
11) Risk and technical depth (known problems and risks).
12) Glossary (ubiquitous language).

Architectural view model:
* Logical view - functionality that system provides to end-users.
* Process view - run-time behavior of system (concurrency, performance, scalability, etc.).
* Development view - development view (implementation view: packages, components, etc.).
* Physical view - system engineer's point of view (instances topology).
* Scenarios - use cases.

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

#### CAP

CAP theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following (pick 2 out of 3):
* Consistency - data is the same across the cluster (read from replica return data written on master).
* Availability - ability to access cluster even if node goes down (can read & write data).
* Partition tolerance - cluster continues to function even if communication break between 2 nodes (network partitioning).

CAP means: network partition happens - you have to make decision tolerate failure or not,
so you may have consistent system or available, but not both.
