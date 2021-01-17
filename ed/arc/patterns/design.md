Design Patterns
-

Design patterns offer a common solution for a common problem in the form of classes working together.

1. Creational (interface to create object):
    * abstractFactory
    * builder (constructor)
    * factoryMethod
    * prototype
    * singleton

2. Structural (way to put objects together):
    * adapter
    * bridge - decouple an abstraction from its implementation
    * composite - compose objects into tree structures
    * decorator - dynamically allows behavior to be added to object
    * facade - object that provides a simplified interface to a larger body of code
    * flyweight - lazy load
    * proxy

3. Behaviour (how classes interface each other and how program operates):
    * chainOfResponsibility - each object contains logic and link to the next
    * command - object encapsulate all information needed to call a method
    * dependencyInjection - IoC
    * interpreter
    * iterator - used to traverse a container
    * mediator - object that encapsulates how a set of objects interact
    * memento - provides the ability to restore an object to its previous state
    * observer (publish-subscribe)
    * state - encapsulate varying behavior based on an object's state
    * strategy - behavior to be selected at runtime
    * templateMethod
    * visitor

## Creational:

**Abstract Factory** offers the interface for creating a *family of related objects*,
without explicitly specifying their classes.
For example: DbProviderFactory (mysql/mongo/etc) - DbCommand, DbConnection, DbParameter, etc.

Unlike the abstract factory pattern and the factory method pattern
whose intention is to enable polymorphism,
the intention of the **Builder** pattern is to find a solution to the telescoping *constructor anti-pattern*.
It occurs when the increase of object constructor parameter combination
leads to an exponential list of constructors.

**Factory Method** defines an interface for creating an object,
but leaves the choice of its type to the subclasses,
*creation being deferred at run-time*.

**Prototype** it is used when the type of objects
to create is determined by a *prototypical instance*.
For example: daily meeting in calendar has same members, same time, etc.
You don't need to provide such information each time,
just instantiate prototype class and use it.

**Singleton** - *single instance*.
Singletons are bad, they create dependency between classes.

## Structural:

An **Adapter** (aka Wrapper) allows classes to work together
that normally could not because of *incompatible interfaces*
by providing its interface to clients while using the original interface.

**Bridge** - intended to "decouple an abstraction from its implementation
so that the two can vary independently".
It *belong to one domain but implement interfaces from another*.

The intent of a **Composite** is to compose objects into *tree structures*.

**Decorator** (aka Wrapper) - design pattern that *allows behavior to be added*
to an individual object, either statically or dynamically,
without affecting the behavior of other objects from the same class.

A **Facade** is an object that provides a *simplified interface*
to a larger body of code, such as a class library.

A **Flyweight** (*lazy load*) is an object that minimizes memory use
by sharing as much data as possible with other similar objects.

A **Proxy**, in its most general form,
is a class functioning as an *interface to something else*.

## Behaviour:

**Chain of Responsibility** is a design pattern consisting of objects,
*each object contains logic and link to the next processing object* in the chain,
that will be invoked.
 
**Command** - is a pattern in which an object is used to represent and encapsulate
*all the information needed to call a method*.

**Dependency Injection** separates the creation of a client's dependencies
from its own behavior, which allows program designs to be loosely coupled
and to follow the *dependency inversion* and single responsibility principles.
Types of dependency injection:
1) constructor injection
2) setter injection
3) interface injection

**Interpreter** - is a pattern that specifies *how to evaluate sentences
in a language*. The basic idea is to have a class
for each symbol (terminal or nonterminal) in a specialized computer language.

**Iterator** - is a pattern in which an iterator is used
to *traverse a container* and access the container's elements.

**Mediator** defines an object that *encapsulates how a set of objects interact*.
For example: brain is a mediator for all body parts.

**Memento** is a pattern that provides the *ability to restore an object
to its previous state* (undo via rollback).

**Observer** (aka *Publish-Subscribe*) - is a design pattern
in which an object, called the subject,
maintains a list of its dependents, called observers,
and notifies them automatically of any state changes.

**State** this pattern is used to encapsulate varying behavior
for the same routine *based on an object's state*.

**Strategy** - is a pattern that enables an algorithm's *behavior
to be selected at runtime*.

**Template method** - it’s another way of encapsulation piece of algorithms
so the sub-classes can hook themselves right into computation
any time they want.
For example: applying same (template) approach for bunch of stuff:
bus, plain, boat: start engine, leave terminal, entertainment, arrive, etc.

**Visitor** - is a way of separating an algorithm from an object structure
on which it operates. A practical result of this separation
is the ability to add new operations to existing object structures
without modifying those structures.
