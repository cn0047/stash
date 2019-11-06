System Design
-

1: Clarify constraints and use cases.
2: Abstract design.
3: Understand bottlenecks.
4: Scale your abstract design.

Primary concerns:
* Reliability (fault tolerant)
* Scalability (increasing load)
* Maintainability (code that can easily be understood, refactored and upgraded)

CAP Theorem states that in a distributed system,
it is impossible to simultaneously guarantee all of the following (pick 2out of 3):
* Consistency
* Availability
* Partition tolerance

Storage Scalability:
* What is the amount of data that we need to store?
* Will the data keep growing over time? If yes, then at what rate?

Elements of a System:
* Architecture
* Modules
* Components
* Interfaces
* Data
