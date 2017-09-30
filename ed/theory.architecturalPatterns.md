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

An architectural pattern is a proven structural organization schema for software systems.

## Architectural patterns:

* Layers
* Client-server
* Master-slave
* Pipe-filter
* Broker
* Peer-to-peer
* Blackboard
* Event-bus

* Event-driven architecture
* Active record
* Data mapper
* Implicit invocation
* Naked objects
* Operational Data Store
* Service-oriented architecture
* Microservices

* Model-View-Controller
* Presentation-Abstraction-Control
* Model-View-Presenter
* Model-View-ViewModel (MVVM)
* Entity-Control-Boundary

#### Layers (AKA multitier or three-tier or n-tier)

* Presentation layer (UI layer)
* Application layer (service layer)
* Business logic layer (domain layer)
* Data access layer (persistence layer)

Usage: desktop app, E-commerce web app.

#### Client-server

Usage: online app like email, document sharing and banking.

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

#### Event-bus

`event source`, `event listener`, `channel` and `event bus`.

Usage: android development, notification services.

#### Model-View-Controller

1. model — contains the core functionality and data.
2. view — displays the information to the user.
3. controller — handles the input from the user.

Usage: architecture for World Wide Web app.

#### Blackboard

Useful for problems for which no deterministic solution strategies are known.

Usage: speech recognition, vehicle identification and tracking,
protein structure identification, sonar signals interpretation.

#### Data mapper

(Object Relational Mapper (ORM) and the Database Abstraction Layer (DAL))
Object-relational mapping - is a programming technique
for converting data between incompatible type systems in object-oriented programming languages.

#### Entity-Control-Boundary

Action -> Boundary -> Control (mediator) -> Entity.

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

Microservices is a variant of the SOA
architectural style that structures an application as a collection of loosely coupled services.
It is modern interpretation of SOA used to build distributed software systems.

Microservices-based architectures enable continuous delivery and deployment.

Orchestration - it is when one microservice supervise and manage else microservices.
<br>Choreography - it is when microservice independent, and only have to know about message-queue.

Philosophy: Do one thing and do it well.

The less well you understand a domain,
the harder it will be for you to find proper bounded contexts for your services.

Greenfield development is also quite challenging.
So again, consider starting monolithic first and break things out when you’re stable.

Many of the challenges you’re going to face with microservices get worse with scale.

* Model Around Business Concepts
  (use bounded contexts)
* Adopt a Culture of Automation
  (automated testing, deploy the same way everywhere, continuous delivery)
* Hide Internal Implementation Details
  (modeling bounded contexts, services should also hide their databases, consider using REST)
* Decentralize All the Things
  (teams own their services, align teams to the organization, prefer choreography over orchestration)
* Independently Deployable
  (coexist versioned endpoints, one-service-per-host)
* Isolate Failure
  (expect failure will occur anywhere and everywhere)
* Highly Observable
  (aggregate your logs, aggregate your stats)

Anti-pattern - nanoservice,
is a service whose overhead (communications, maintenance etc.) outweighs its utility.
