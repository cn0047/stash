Monotilth
-

Monolith - unit of deployment.
Monolith is not synonymous with legacy, it has advantages & disadvantages.

Single process monolith.

Modular monolith - subset of single process monolith,
which consists of separate modules.

Distributed monolith - when microservices tightly bounded,
and have high complexity of connections within the system.

Some examples of business problems caused by monotilth:
* Slow delivery.
* Buggy software releases.
* Poor scalability.
* Slow CI.

Strategies for refactoring a monolith to microservices:
* Implement new features as services.
* Separate presentation layer and backend.
* Break up the monolith by extracting functionality into services.
