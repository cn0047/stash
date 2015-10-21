Patterns
-

####Design Patterns:

1. Creational.
2. Structural.
3. Behaviour.

This grouping described using the concepts of delegation, aggregation, and consultation.

Delegation is the simple yet powerful concept of handing a task over to another part of the program.
Object assigns a task to another object, known as the delegate.

In aggregation, the object may only contain a reference or pointer to the object
(and not have lifetime responsibility for it).
<i>
For example, a university owns various departments, and each department has a number of professors.
If the university closes, the departments will no longer exist,
but the professors in those departments will continue to exist.
University can be seen as a composition of departments,
whereas departments have an aggregation of professors.
In addition, a Professor could work in more than one department,
but a department could not be part of more than one university.
</i>

Consultation in object-oriented programming occurs when an object's method implementation consists
of a message send of the same message to another constituent object.

````
Adapter vs Bridge: adapter works with defined objects, bridge can accept objects on fly.
Facade vs Proxy: facade cover defined objects, proxy provide access to objects and filtrate or restrict access to object.
Prototype vs Interpreter: ?
Bridge vs Mediator: ?
Command vs Strategy: ?

Singletons are bad, they create dependency between classes.
````

####Architectural patterns:

* Active record
* Blackboard system
* Data mapper
* Event-driven architecture
* Implicit invocation
* Layers
* Microservices
* Model-View-Controller, Presentation-abstraction-control, Model View Presenter, and Model View ViewModel
* Multitier architecture (often three-tier or n-tier)
* Naked objects
* Operational Data Store (ODS)
* Peer-to-peer
* Pipe and filter architecture
* Service-oriented architecture

####Else patterns:

* Data mapper pattern (Object Relational Mapper (ORM) and the Database Abstraction Layer)
