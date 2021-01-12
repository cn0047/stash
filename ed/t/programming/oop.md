OOP
-

Since 1950. Released in 1967.

**Encapsulation** - hide the values or state of a structured data object inside a class.

**Polymorphism** - provision of a single interface to entities of different types
(same name, different behavior).

**Inheritance** - when an object or class is based on another object (prototypal inheritance)
or class (class-based inheritance).

**Abstraction** - is the development of classes, objects, types
in terms of their interfaces and functionality, instead of their implementation details.

Class - program-code-template for creating objects.

Object - runtime concrete occurrence of class instance.

#### Problems

Inheritance:

* Banana Monkey Jungle Problem.
* The Fragile Base Class Problem.
* The Hierarchy Problem (deep, unclear, multiple).

Design patterns and delegation brakes inheritance.

Encapsulation:

* The Reference Problem (update object by reference)

Polymorphism:

No need in OOP, just need interface-based polymorphism.

#### OOD

Concepts of **delegation**, **aggregation**, and **consultation**.

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
of a message send to another object.

#### Aggregation and Composition

Aggregation - child can exist independently of the parent (Car -> Tires).

Composition - child can NOT exist independent of the parent (House -> Room).
