Theory
-

Even (2, 4, 6...) - `x % 2 = 0`.
Odd (1, 3, 5...) - `x % 2 = 1`.

AJAX polling - send request in loop.
Long-polling - ...
Forewer frame - ...
HTM5 Server-sent events - ...

Bandwidth - count of lines on high-way.
Latency - speed limit on high-way.

Concurrency - single process use multiple threads.

Race condition - when events do not happen in the order the programmer intended.

The ideal task size is not bigger than a day's work...

Two-pizza teams - where no team should be so big that it could not be fed with two pizzas.

Resilience - is the ability of a system to adapt or keep working when challenges occur.

Pure function - 1) always return the same result for same arguments, and 2) no mutation or output to I/O etc.

Function is DETERMINISTIC if it always produces the same result for the same input parameters.

Internationalization (i18n) - process of changing your software so that it isn't hardwired to one language.

Localization (l10n) - the process of adding the appropriate resources to your software so that a particular language/locale is supported.

Tell, Don’t Ask.

Law of Demeter - object A can call method of object B,
but object A should not "reach through" object B to access yet another object C, to request its services.

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

#### Security

Symmetric encryption (symmetric key cryptography) - sender and receiver have same secret key.
Same key for encrypting and decrypting.

Public key cryptography - sender have public key, receiver have private key.
It is expensive type of cryptography.

#### Full stack:

* presentation layer (html, css, js)
* business layer (node, php, etc)
* data access layer (mongo, mysql, etc)

#### Low coupling and high cohesion.

COUPLING refers to the interdependencies between modules.

LOW COUPLING is often a sign of a well-structured computer system and a **good** design.

COHESION describes how related the functions within a single module are.
<br>
Classes should have a small number of instance variables.
In general the more variables a method manipulates the more cohesive that method is to its class.
A class in which each variable is used by each method is maximally cohesive.

#### Continuous Integration/Continuous Delivery

How long would it take your organization to deploy a change that involves just one single line of code?
Do you do this on a repeatable, reliable basis?

A deployment pipeline is, in essence, an automated implementation of your application’s
build, deploy, test, and release process.

The practice of building and testing your application on every check-in
is known as continuous integration.

Continuous Integration - is a software development practice
where members of a team use a version control system and integrate their work frequently to the same location,
such as a master branch.
Each change is built and verified by tests and other verifications
in order to detect any integration errors as quickly as possible.

Continuous Delivery - is a software development methodology where the release process is automated.
Every software change is automatically built, tested, and deployed to production.

Continuous Deployment - is a synonym to Continuous Delivery.

Pre-alpha ⇒ Alpha ⇒ Beta ⇒ Release candidate ⇒ Gold

#### RPC

Problems:

* How should the client react if there are no servers running or RPC server is down for a long time?
* Should a client have some kind of timeout for the RPC?
* If the server malfunctions and raises an exception, should it be forwarded to the client?
* Protecting against invalid incoming messages (eg checking bounds, type) before processing.

#### Code quality

* Maintainable
* Extendable
* Reusable
* Understandable
* Testable
* Documentable
* Well designed (patterns)
* Follow SOLID
* Haven't memory leaks
* Haven't vulnerabilities and be secure
* Predictability
* Cares about backward compatibility

#### ACID

<ul>
    <li><i>Atomicity</i> - all or nothing.</li>
    <li><i>Consistency</i> ensures that any transaction will bring the database from one valid state to another (constraints, cascades, triggers).</li>
    <li><i>Isolation</i> ensures that the concurrent execution of transactions will executed serially, i.e., one after the other.</li>
    <li><i>Durability</i> ensures that once a transaction has been committed, it will remain so, even in the event of power loss, crashes, or errors...</li>
</ul>

#### 10 Questions Developers Should be Asking Themselves

* Is there a pattern here?
* How can I make this simpler?
* Why does it work like that?
  (Knowing that something works and knowing why it works that way are two very  different things.)
* Has somebody done this before?
* Who said it first?
  (Always try read the original source of a concept or theory.)
* Do I love what I’m doing?
* Where else could I use this?
* What did I fail at today?
* How can we make this possible?
  (Start from the assumption that whatever you want to do is possible.)
* Who can I learn from?
  (You should never work anywhere where you are the smartest person in the room.)

#### Programming:

Imperative programming is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

Declarative programming focuses on what the program should accomplish
without specifying how the program should achieve the result.

The declarative layer describes what the code will do,
while the implementation layer describes how the code does it.
(The declarative layer is, in effect, a small  domain-specific language).

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
