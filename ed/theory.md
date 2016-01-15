Theory
-

Tell, Don’t Ask (Law of Demeter).

#### ACID

<ul>
    <li><i>Atomicity</i> - all or nothing.</li>
    <li><i>Consistency</i> ensures that any transaction will bring the database from one valid state to another (constraints, cascades, triggers).</li>
    <li><i>Isolation</i> ensures that the concurrent execution of transactions will executed serially, i.e., one after the other.</li>
    <li><i>Durability</i> ensures that once a transaction has been committed, it will remain so, even in the event of power loss, crashes, or errors...</li>
</ul>

#### SOLID

<ul>
    <li>
        (SRP) Single responsibility principle - states that every class should have responsibility over a single part of the functionality provided by the software.
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

#### IoC:

Inversion of control - is used to increase modularity of the program and make it extensible.
<br>Software frameworks, callbacks, schedulers, event loops and dependency injection
are examples of design patterns that follow the inversion of control principle.

IoC serves the following design purposes:
<ul>
    <li>To decouple the execution of a task from implementation.</li>
    <li>To focus a module on the task it is designed for.</li>
    <li>
        To free modules from assumptions about how other systems do
        what they do and instead rely on contracts.
    </li>
    <li>To prevent side effects when replacing a module.</li>
</ul>
("Hollywood Principle: Don't call us, we'll call you".)
<br>There are several basic techniques to implement inversion of control:
<ul>
    <li>Using a factory pattern</li>
    <li>Using a service locator pattern</li>
    <li>
        Using a dependency injection, for example
        (constructor injection, parameter injection, setter injection, interface injection).
    </li>
    <li>Using a contextualized lookup</li>
    <li>Using template method design pattern</li>
    <li>Using strategy design pattern</li>
</ul>

#### Programming:

Imperative programming is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

Declarative programming focuses on what the program should accomplish
without specifying how the program should achieve the result.

#### Extreme Programming:

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

#### Agile:

The Agile Manifesto is based on 12 principles:

1.  Customer satisfaction by rapid delivery of useful software
2.  Welcome changing requirements, even late in development
3.  Working software is delivered frequently (weeks rather than months)
4.  Close, daily cooperation between business people and developers
5.  Projects are built around motivated individuals, who should be trusted
6.  Face-to-face conversation is the best form of communication (co-location)
7.  Working software is the principal measure of progress
8.  Sustainable development, able to maintain a constant pace
9.  Continuous attention to technical excellence and good design
10. Simplicity—the art of maximizing the amount of work not done—is essential
11. Self-organizing teams
11. Regular adaptation to changing circumstance

Agile software development methods:
* Adaptive software development (ASD)
* Agile modeling
* Agile Unified Process (AUP)
* Crystal Clear Methods
* Disciplined agile delivery
* Dynamic systems development method (DSDM)
* Extreme programming (XP)
* Feature-driven development (FDD)
* Lean software development
* Kanban (development)
* Scrum
* Scrum ban

Agile practices:
* Acceptance test-driven development (ATDD)
* Agile modeling
* Backlogs (Product and Sprint)
* Behavior-driven development (BDD)
* Cross-functional team
* Continuous integration (CI)
* Domain-driven design (DDD)
* Information radiators (scrum board, task board, visual management board, burndown chart)
* Iterative and incremental development (IID)
* Pair programming
* Planning poker
* Refactoring
* Scrum events (sprint planning, daily scrum, sprint review and retrospective)
* Test-driven development (TDD)
* Agile testing
* Timeboxing
* Use case
* User story
* Story-driven modeling
* Retrospective
* Velocity tracking
