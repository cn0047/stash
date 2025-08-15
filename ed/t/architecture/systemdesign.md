System Design
-

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
