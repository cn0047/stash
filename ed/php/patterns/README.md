Patterns
-

####Design Patterns:

1. Creational.

    * abstractFactory
    * builder
    * factoryMethod
    * prototype
    * singleton

2. Structural.

    * adapter
    * bridge - decouple an abstraction from its implementation
    * composite - "compose" objects into tree structures
    * facade - object that provides a simplified interface to a larger body of code
    * flyweight - lazy load
    * proxy

3. Behaviour.

    * chainOfResponsibility - each object contains logic and link to the next
    * command - object encapsulate all information needed to call a method
    * decorator - dynamically allows behavior to be added to object
    * dependencyInjection - IoC
    * interpreter
    * iterator - used to traverse a container
    * mediator - object that encapsulates how a set of objects interact
    * memento - provides the ability to restore an object to its previous state
    * observer
    * state - encapsulate varying behavior based on an object's state
    * strategy - behavior to be selected at runtime
    * templateMethod
    * visitor

This grouping described using the concepts of **delegation**, **aggregation**, and **consultation**.

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
Adapter vs Bridge:
Adapter works with defined objects,
bridge can accept objects on fly, because it is abstraction over implementation.

Prototype vs Interpreter: ?

Command vs Strategy: ?
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
* Service-oriented architecture (SOA)

####Else patterns:

* Data mapper pattern (Object Relational Mapper (ORM) and the Database Abstraction Layer)
  Object-relational mapping - is a programming technique
  for converting data between incompatible type systems in object-oriented programming languages.
