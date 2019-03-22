Architectural patterns
-

Architecture must care about:
* Durability
* Utility
* Beauty

Common questions.

MVC, MVP, MVVM etc:

* What about helpers?
* What about stuff common for 2, 3 models?

## Architectural patterns:

An architectural pattern is a proven structural organization schema for software systems.

* Layers
* **Client-server**
* **Master-slave**
* Pipe-filter
* **Broker**
* **Peer-to-peer**
* Blackboard
* **Event-bus**
* **CQRS**

* Event-driven architecture
* **Active record**
* **Data mapper**
* Implicit invocation
* Naked objects
* Operational Data Store
* **Service-oriented architecture**
* **Microservices**
* **API Gateway**

* Model-View-Controller
* Presentation-Abstraction-Control
* Model-View-Presenter
* Model-View-ViewModel (MVVM)
* Entity-Control-Boundary

#### Layers (AKA multitier or 3-tier or n-tier)

* Presentation layer (UI layer)
* Application layer (service layer)
* Business logic layer (domain layer)
* Data access layer (persistence layer)

Classic 3-tier: Presentation -> Business -> Data.

Usage: desktop app, E-commerce web app.

#### Client-server

Usage: online app like email, document sharing and banking.

#### Master-slave

#### Pipe-filter

For data streams. Data to be processed is passed through pipes.

Usage: compilers, lexical analysis, parsing, semantic analysis, and code generation.

#### Broker

For distributed systems with decoupled components.
Components can interact with each other by remote service invocations.
A `broker` component is responsible for the coordination of communication among components.

Usage: message broker (Kafka, RabbitMQ).

#### Peer-to-peer

Individual components are known as `peers`.
Peers may function both as a `client`, requesting services from other peers,
and as a `server` providing services to other peers.

Usage: file-sharing networks, multimedia protocols, something like spotify.

#### Blackboard

Useful for problems for which no deterministic solution strategies are known.

Usage: speech recognition, vehicle identification and tracking,
protein structure identification, sonar signals interpretation.

#### Event-bus

`event source`, `event listener`, `channel` and `event bus`.

Usage: android development, notification services.

#### CQRS

Type 1 - Single database CQRS.

Type 2 - Two database CQRS:

* Commands use write DB
* Queries use read DB

Type 3 - **Event Sourcing**:

Store events -> Replay events -> Modify entity -> Store new event -> Update read DB.

Pros:

* Roint-in-time reconstruction.
* Multiple read DBs.
* Rebuild PROD DB.

#### Event-driven architecture

#### Active record

#### Data mapper

(Object Relational Mapper (ORM) and the Database Abstraction Layer (DAL))
Object-relational mapping - is a programming technique
for converting data between incompatible type systems in object-oriented programming languages.

#### Implicit invocation

#### Naked objects

#### Operational Data Store

#### Service Oriented Architecture

SOA is less about how to modularize an application,
and more about how to compose an application by integration
of distributed, separately-maintained and deployed software components.

In SOA, services use protocols that describe how they pass and parse messages using description metadata.

Each SOA building block can be:

* Service provider
* Service broker
* Service requester/consumer

#### Microservices

#### API Gateway

#### Model-View-Controller

1. model — contains the core functionality and data.
2. view — displays the information to the user.
3. controller — handles the input from the user.

Usage: architecture for World Wide Web app.

#### Presentation-Abstraction-Control

#### Model-View-Presenter

#### Model-View-ViewModel (MVVM)

#### Entity-Control-Boundary

Action -> Boundary -> Control (mediator) -> Entity.
