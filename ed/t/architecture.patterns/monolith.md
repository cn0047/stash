Monotilth
-

**Monolith** - unit of deployment (single process monolith).
Monolith is not synonymous with legacy, it has advantages & disadvantages.

**Modular monolith** - subset of single process monolith,
which consists of separate modules (like vertical layers for domains).

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

Monolith isn't so awful in case of server-side-rendering.
But:
* [10K SLOC](https://en.wikipedia.org/wiki/Source_lines_of_code).
* Extremely difficult to navigate the code and isolate your code.
