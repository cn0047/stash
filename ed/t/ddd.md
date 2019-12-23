DDD (Domain-driven design)
-

#### Domain.

A sphere of knowledge, influence, or activity. The subject area.

#### Model.

A system of abstractions that describes selected aspects of a domain.
(aggregates, entities, factories)

#### Ubiquitous Language.

A language structured around the domain model
and used by all team members.

#### Bounded Context.

Multiple models are in play on any large project.
It is often unclear in what context a model should not be applied.

#### Continuous integration.

#### Context map.

#### Layered Architecture.

* User Interface - Responsible for presenting information to the user and
interpreting user commands.

* Application - This is a thin layer which coordinates the application activity.
It does not contain business logic.
It does not hold the state of the business objects,
but it can hold the state of an application task progress.
(Used by external consumers to talk to your system).
(Application services should generally represent all possible use cases).
(It can check whether a domain object exists or not and throw exceptions accordingly).

* Domain - This layer contains information about the domain.
This is the heart of the business software.
The state of business objects is held here.
Persistence of the business objects and possibly their state is delegated to
the infrastructure layer.

* Infrastructure - This layer acts as a supporting library for all the other layers.
It provides communication between layers,
implements persistence for business objects, contains
supporting libraries for the user interface layer, etc.
(Order is a domain concept whereas Table, Column and so on are infrastructure concerns.)
(Here we put all the implementations of the interfaces defined in the domain layer.)

Example:

````
src/Domain/Model/ParticularModel/ParticularModel.php - Doctrine entity.
src/Domain/Model/ParticularModel/DTO/ParticularModel.php.
src/Domain/Model/ParticularModel/Service/Command/CreateParticularModel.php - (CQRS) Call persister commands.
src/Domain/Model/ParticularModel/Service/Query/CreateParticularModel.php - (CQRS) Call persister commands.
src/Infrastructure/Command/ParticularModel/PersisterDoctrine.php - Doctrine EM Wrapper.
src/Infrastructure/DataProvider/
````

#### Entities.

An object that is not defined by its attributes,
but rather by a thread of continuity and its identity.
Is a category of objects which seem to have an identity.

#### Value Objects (VO).

An object that contains attributes but has no conceptual identity.

VOs must be placed in application layer
because only this layer aware how to interact with domain layer from the outside world.

#### Modules.

For a large and complex application, the model tends to grow
bigger and bigger. The model reaches a point where it is hard to
talk about as a whole, and understanding the relationships and
interactions between different parts becomes difficult. For that
reason, it is necessary to organize the model into modules.
Modules are used as a method of organizing related concepts
and tasks in order to reduce complexity.

Each module must have src directory.
This folder contains all the code necessary for this bounded context to work:
domain code and infrastructure code.

Something like:

````
├──composer.json
├──composer.lock
├──src
│  └── BuyIt
│      ├── Billing
│      │   ├── Domain
│      │   │   ├── Model
│      │   │   │   ├ Bill (contains: factory + model|entity + events + interface for repository & cqrs)
│      │   │   │   ├ Order
│      │   │   │   └ Waybill
│      │   │   └── Service
│      │   └── Infrastructure
│      │       ├── Logging
│      │       ├── Messaging
│      │       ├── FullTextSearching
│      │       │   └── Elastica
│      │       └── Persistence
│      │           ├── Doctrine
│      │           ├── SQL
│      │           ├── InMemory
│      │           └── Redis
│      ├── Catalog
│      ├── Common
│      └── Identity
└── tests
````

#### Aggregates.

A collection of objects that are bound together by a root entity.
A model can contain a large number of domain objects.
No matter how much consideration we put in the design, it happens
that many objects are associated with one another, creating a
complex net of relationships.
There are several types of associations (one-to-one, many-to-many...).

A DDD aggregate is a cluster of domain objects
that can be treated as a single unit.

Car it is aggregate for: wheels, engine, spark and fuel, etc.

#### Domain Event.

A domain object that defines an event (something that happens).

* Modeling a Domain Event is like writing a news article
* Publishing a Domain Event is like printing the article on the paper
* Spreading a Domain Event is like sending the newspaper so everyone can read the article

#### DBAL - Database Abstraction Layer.

Active Record ORMs not good for DDD, because:

* Active Record pattern assumes a one-to-one relation between an entity and a database table.
And in a rich domain model sometimes entities are constructed with information
that may come from different data sources.

* Advanced things like collections or inheritance are tricky to implement.

* Possible persistence leakage into the domain model
by coupling the domain model with the ORM.

ORM Doctrine is an implementation of the Data Mapper pattern.

Doctrine annotations is bad for DDD, because:
* domain concerns are mixed with infrastructure concerns
* if the entity were required to be persisted using another entity
manager and with a different mapping metadata, it would not be possible.

So better use XML mapping files.

#### Factory.

Methods for creating domain objects
should delegate to a specialized Factory.

#### CQRS.

Command Query Responsibility Segregation.

#### Service.

When an operation does not conceptually belong to any object.
An object does not have an internal state, and its purpose is to simply provide
functionality for the domain.
We should not create a Service for every operation needed.
But when such an operation stands out as an important concept in the domain,
a Service should be created for it.

There are three characteristics of a Service:

1. The operation performed by the Service refers to a domain
concept which does not naturally belong to an Entity or Value Object.

2. The operation performed refers to other objects in the domain.

3. The operation is stateless.

There are typically three different types of service:
* Application (middleware between the outside world and the domain logic)
* Domain (domain services are stateless)
* Infrastructure (sending emails, logging meaningful data etc)

In DDD, transactions are handled at the Application Service level (for example TransactionalApplicationService).

Domain Services are used to describe things into the domain,
operations that don’t belong to entities nor value objects.
(Cross-aggregate behavor, repositories, external services).

#### DTO - Data transfer object.

Communication between the delivery mechanism
and the domain is carried by data structures called DTO.

DTO it is something like request/response VO for domain.
DTO does not have any behavior except for storage and retrieval of its own data.
DTOs are simple objects that should not contain any business logic.

Interface to DTO must be placed in domain layer.
<br>Particular DTO implementation (mysql, mongo, etc) must be placed in infrastructure layer
because it contains specific stuff (related to rows in mysql, how to get data, how to transform, etc).

DTO it's just data container which used to transport data between different layers.

#### Repositories.

Methods for retrieving domain objects
should delegate to a specialized Repository object
such that alternative storage implementations may be easily interchanged.

Repositories are not DAOs.

#### DAO - Data access object.

Typically a DAO would contain CRUD methods for a particular domain object.
DAOs must be placed in domain layer.

## Real problems.

* Configs and DI (especially DI) must be placed outside any layer (user-interface, application, domain, infrastructure)
  because they are not part of any layer.

* Exceptions must be present in each layer,
  with purpose to describe particular problems of certain layer.

Domain event publisher, and technical stuff (publisher-subscriber or bus patterns internal implementations)?
Switch from one php framework to another?
Switch from one front-end framework to another? ~~And server-side rendering?~~

Confusion:
* VO, DTO - own meaning in Java world.
* Entity - in symfony world.
