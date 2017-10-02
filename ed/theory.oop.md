OOP
-

#### Problems

Inheritance:

* Banana Monkey Jungle Problem.
* The Fragile Base Class Problem.
* The Hierarchy Problem (deep, unclear, multiple).

Design patterns and delegation brakes inheritance.

Encapsulation:

* The Reference Problem (update object by reference)

Polymorphism:

No need OOP just need interface-based polymorphism.

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
of a message send of the same message to another constituent object.

#### SOLID

<ul>
    <li>
        (SRP) Single responsibility principle - states that every class should have responsibility over a single part
        of the functionality provided by the software. If description contains word AND or OR - it's not SRP.)
    </li>
    <li>
        (OCP) Open/closed principle - (classes, modules, functions, etc.) should be open for extension, but closed for modification.
    </li>
    <li>
        (LSP) Liskov substitution principle - if S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of the program.
    </li>
    <li>
        (ISP) Interface segregation principle - splits interfaces which are very large into smaller and more specific ones.
    </li>
    <li>
        (DIP) Dependency inversion principle - refers to a specific form of decoupling software modules.
        <ul>
            <li>High-level modules should not depend on low-level modules. Both should depend on abstractions.</li>
            <li>Abstractions should not depend on details. Details should depend on abstractions.</li>
            <li>Ownership inversion - both high- and lower-level layers should depend on abstractions that draw the behavior.</li>
            <li>
                Abstraction dependency:
                <ul>
                    <li>All member variables in a class must be interfaces or abstracts.</li>
                    <li>All concrete class packages must connect only through interface/abstract classes packages.</li>
                    <li>No class should derive from a concrete class.</li>
                    <li>No method should override an implemented method.</li>
                    <li>All variable instantiation requires the implementation of a Creational pattern as the Factory Method or the Factory pattern, or the more complex use of a Dependency Injection framework.</li>
                </ul>
            </li>
        </ul>
    </li>
</ul>
