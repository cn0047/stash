DDD
-

Domain - A sphere of knowledge, influence, or activity. The subject area.

Model - A system of abstractions that describes selected aspects of a domain.

Ubiquitous Language - A language structured around the domain model
and used by all team members.

Bounded Context - Multiple models are in play on any large project.
It is often unclear in what context a model should not be applied.

Continuous integration.

Context map.

Layered Architecture:
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

Entities - An object that is not defined by its attributes,
but rather by a thread of continuity and its identity.
Is a category of objects which seem to have an identity.

Value Objects - An object that contains attributes but has no conceptual identity.

Modules - For a large and complex application, the model tends to grow
bigger and bigger. The model reaches a point where it is hard to
talk about as a whole, and understanding the relationships and
interactions between different parts becomes difficult. For that
reason, it is necessary to organize the model into modules.
Modules are used as a method of organizing related concepts
and tasks in order to reduce complexity.

Aggregates - A collection of objects that are bound together by a root entity.
A model can contain a large number of domain objects. No
matter how much consideration we put in the design, it happens
that many objects are associated with one another, creating a
complex net of relationships.
There are several types of associations (one-to-one, many-to-many...).

Domain Event - A domain object that defines an event (something that happens).

Repositories - Methods for retrieving domain objects
should delegate to a specialized Repository object
such that alternative storage implementations may be easily interchanged.

Factory - Methods for creating domain objects
should delegate to a specialized Factory.

Service - When an operation does not conceptually belong to any object.
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
