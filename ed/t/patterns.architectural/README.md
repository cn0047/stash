Architectural patterns
-

Architecture must care about:
* Durability
* Utility
* Beauty

Common questions:
1. MVC, MVP, MVVM etc:
    * What about helpers?
    * What about stuff common for 2, 3 models?

An architectural pattern is a proven structural organization schema for software systems.

* Model-View-Controller
* Presentation-Abstraction-Control
* Model-View-Presenter
* Model-View-ViewModel (MVVM)
* Entity-Control-Boundary
* Layers

* Active record
* Data mapper
* CQRS
* Operational Data Store

* Client-server
* Master-slave
* Peer-to-peer
* Pipe-filter

* API Gateway
* Broker
* Event-bus

* Event-driven architecture
* Service-oriented architecture
* Microservices
* Blackboard

* Implicit invocation
* Naked objects

**Client-server**
Usage: online app like email, document sharing and banking.

**Pipe-filter**
For data streams. Data to be processed is passed through pipes.
<br>Usage: compilers, lexical analysis, parsing, semantic analysis, and code generation.

**Broker**
For distributed systems with decoupled components.
Components can interact with each other by remote service invocations.
A `broker` component is responsible for the coordination of communication among components.
<br>Usage: message broker (Kafka, RabbitMQ).

**Peer-to-peer**
Individual components are known as `peers`.
Peers may function both as a `client`, requesting services from other peers,
and as a `server` providing services to other peers.
<br>Usage: file-sharing networks, multimedia protocols, something like spotify.

**Blackboard**
Useful for problems for which no deterministic solution strategies are known.
<br>
Usage: speech recognition, vehicle identification and tracking,
protein structure identification, sonar signals interpretation.

**Event-bus**
`event source`, `event listener`, `channel` and `event bus`.
<br>Usage: android development, notification services.

**Entity-Control-Boundary**
Action -> Boundary -> Control (mediator) -> Entity.
