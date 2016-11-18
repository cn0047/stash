DDD
-

### Domain

A sphere of knowledge, influence, or activity. The subject area.

### Model

A system of abstractions that describes selected aspects of a domain.

### Ubiquitous Language

A language structured around the domain model
and used by all team members.

### Bounded Context

Multiple models are in play on any large project.
It is often unclear in what context a model should not be applied.

### Continuous integration.

### Context map.

### Layered Architecture

* User Interface - Responsible for presenting information to the user and
interpreting user commands.
* Application - This is a thin layer which coordinates the application
activity. It does not contain business logic. It does not
hold the state of the business objects, but it can hold
the state of an application task progress.
* Domain - This layer contains information about the domain. This
is the heart of the business software. The state of
business objects is held here. Persistence of the
business objects and possibly their state is delegated to
the infrastructure layer.
* Infrastructure - This layer acts as a supporting library for all the other
layers. It provides communication between layers,
implements persistence for business objects, contains
supporting libraries for the user interface layer, etc.
(Order is a domain concept whereas Table, Column and so on are infrastructure concerns.)

Example:

````
src/Domain/Model/ParticularModel/ParticularModel.php - Doctrine entity.
src/Domain/Model/ParticularModel/DTO/ParticularModel.php.
src/Domain/Model/ParticularModel/Service/Command/CreateParticularModel.php - (CQRS) Call persister commands.
src/Domain/Model/ParticularModel/Service/Query/CreateParticularModel.php - (CQRS) Call persister commands.
src/Domain/Service/
src/Domain/VO/
src/Infrastructure/Command/ParticularModel/PersisterDoctrine.php - Doctrine EM Wrapper.
src/Infrastructure/DataProvider/
src/Infrastructure/Persistence/Doctrine/ORM/Model.ParticularModel.ParticularModel.orm.xml
src/ProjectFrameworkFiles/
````

````
Application\Service
BuyIt\Billing\Infrastructure\FullTextSearching\Elastica
Ddd\Billing\Domain\Model\Order
Ddd\Billing\Domain\Model\Order\Order
Ddd\Billing\Domain\Model\Order\OrderId
Ddd\Billing\Domain\Model\Order\OrderRepository
Ddd\Billing\Infrastructure\Doctrine\Order\DoctrineOrderRepository
Ddd\Catalog\Domain\Model\Book
Ddd\Common\Domain\Model
Ddd\Domain\Model\Currency
Ddd\Domain\Model\Money
Ddd\Identity\Domain\Model
Doctrine\DBAL\Platforms\AbstractPlatform
Doctrine\DBAL\Types\Type
Doctrine\ORM\EntityManager
Doctrine\ORM\EntityRepository
Doctrine\ORM\Tools
Doctrine\ORM\Tools\Setup
Domain\Model\Body
Domain\Model\Post
Domain\Model\PostId
Domain\Model\PostRepository
Domain\Model\PostSpecificationFactory
Idy\Console\Command
Infrastructure\Persistence\Doctrine
Infrastructure\Persistence\Doctrine\Types
Infrastructure\Persistence\InMemory
Infrastructure\Persistence\Redis
Infrastructure\Persistence\Sql
````

### Entities

An object that is not defined by its attributes,
but rather by a thread of continuity and its identity.
Is a category of objects which seem to have an identity.

### Value Objects

An object that contains attributes but has no conceptual identity.

### Modules

For a large and complex application, the model tends to grow
bigger and bigger. The model reaches a point where it is hard to
talk about as a whole, and understanding the relationships and
interactions between different parts becomes difficult. For that
reason, it is necessary to organize the model into modules.
Modules are used as a method of organizing related concepts
and tasks in order to reduce complexity.

### Aggregates

A collection of objects that are bound together by a root entity.
A model can contain a large number of domain objects. No
matter how much consideration we put in the design, it happens
that many objects are associated with one another, creating a
complex net of relationships.
There are several types of associations (one-to-one, many-to-many...).

### Domain Event

A domain object that defines an event (something that happens).

### DBAL

Active Record ORMs not good for DDD, because:

* Active Record pattern assumes a one-to-one relation between an entity and a database table.
And in a rich domain model sometimes entities are constructed with information
that may come from different data sources.

* Advanced things like collections or inheritance are tricky to implement

* Possible persistence leakage into the domain model
by coupling the domain model with the ORM.

ORM Doctrine is an implementation of the Data Mapper pattern.

Doctrine annotations is bad for DDD, because:
* domain concerns are mixed with infrastructure concerns
* if the entity were required to be persisted using another entity
manager and with a different mapping metadata, it would not be possible.

So better use XML mapping files.

### DTO - Data transfer object.

### Repositories

Methods for retrieving domain objects
should delegate to a specialized Repository object
such that alternative storage implementations may be easily interchanged.

### Factory

Methods for creating domain objects
should delegate to a specialized Factory.

### CQRS

Command Query Responsibility Segregation.

### Service

When an operation does not conceptually belong to any object.
An object does not have an internal state, and its purpose is to simply provide
functionality for the domain.
We should not create a Service for
every operation needed. But when such an operation stands out
as an important concept in the domain, a Service should be
created for it.

There are three characteristics of a Service:

1. The operation performed by the Service refers to a domain
concept which does not naturally belong to an Entity or Value
Object.

2. The operation performed refers to other objects in the domain.

3. The operation is stateless.
