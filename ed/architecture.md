Architecture
-

Architecture must care about:

* Durability
* Utility
* Beauty

#### Architectural patterns:

* Active record
* Blackboard system
* Data mapper pattern (Object Relational Mapper (ORM) and the Database Abstraction Layer)
  Object-relational mapping - is a programming technique
  for converting data between incompatible type systems in object-oriented programming languages.
* Event-driven architecture
* Implicit invocation
* Layers
* Microservices
* Model-View-Controller, Presentation-abstraction-control, Model View Presenter, and Model View ViewModel (MVVM)
* Multitier architecture (often three-tier or n-tier)
* Naked objects
* Operational Data Store (ODS)
* Peer-to-peer
* Pipe and filter architecture
* Service-oriented architecture (SOA)

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
