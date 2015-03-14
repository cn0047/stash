Principles
-

* DRY (don't repeat yourself)
* TDD (Test-driven development)


####SOLID
* (SRP) Single responsibility principle - states that every class should have responsibility over a single part of the functionality provided by the software.
* (OCP) Open/closed principle - (classes, modules, functions, etc.) should be open for extension, but closed for modification.
* (LSP) Liskov substitution principle - if S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of the program.
* (ISP) Interface segregation principle - splits interfaces which are very large into smaller and more specific ones.
* (DIP) Dependency inversion principle - refers to a specific form of decoupling software modules.
    <br>*High-level modules should not depend on low-level modules. Both should depend on abstractions.*
    <br>*Abstractions should not depend on details. Details should depend on abstractions.*
    * Ownership inversion - both high- and lower-level layers should depend on abstractions that draw the behavior.
    * Abstraction dependency:
        * All member variables in a class must be interfaces or abstracts.
        * All concrete class packages must connect only through interface/abstract classes packages.
        * No class should derive from a concrete class.
        * No method should override an implemented method.
        * All variable instantiation requires the implementation of a Creational pattern as the Factory Method or the Factory pattern, or the more complex use of a Dependency Injection framework.

####Extreme Programming
Core Practices:
* Whole Team
* Planning Game
<br>*Release Planning*
<br>*Iteration Planning*
* Customer Tests
* Small Releases
* Simple Design
* Pair Programming
* Test-Driven Development
* Design Improvement
<br>*Refactoring*
* Continuous Integration
* Collective Code Ownership
* Coding Standard
* Metaphor
<br>*(common vision of how the program works)*
* Sustainable Pace
