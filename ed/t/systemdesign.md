System Design
-

````sh
# Time-based availability
availability = uptime / (uptime + downtime)

# Aggregate availability
availability = successful requests / total requests
````

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

System design plan:
* Clarify constraints and use cases, functional & non-functional requirements.
* Abstract design, propose high-level design.
* Understand bottlenecks.
* Scale your abstract design.
* Wrap up.

Steps:
* Do math (how many reads/writes, total users, daily active users, size of message/post/etc., PRS, etc.)
* Draw diagram app/topology/design/interactions/use case/etc.
* Iteratively go into details and deep dive.
* Q/A session.

Storage scalability:
* What is the amount of data that we need to store?
* Will the data keep growing over time? If yes, then at what rate?

Primary concerns:
* Reliability (fault tolerant).
* Scalability (increasing load).
* Maintainability (code that can easily be understood, refactored and upgraded).

CAP Theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following (pick 2 out of 3):
* Consistency - data is the same across the cluster.
* Availability - ability to access cluster even if node goes down.
* Partition tolerance - cluster continues to function even if communication break between 2 nodes (network partitioning).

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
* Latency.
* Keep all geographically close to each other.
* Scale & Performance.

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
