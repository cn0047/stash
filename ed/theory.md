Theory
-

Bandwidth - count of lines on high-way.
Latency - speed limit on high-way.

Concurrency - single process use multiple threads.

Race condition - when events do not happen in the order the programmer intended.

The ideal task size is not bigger than a day's work...

Two-pizza teams - where no team should be so big that it could not be fed with two pizzas.

Orchestration - it is when one microservice supervise and manage else microservices.
<br>Choreography - it is when microservice independent, and only have to know about message-queue.

Resilience - is the ability of a system to adapt or keep working when challenges occur.

Pure function - 1) always return the same result for same arguments, and 2) no mutation or output to I/O etc.

Function is DETERMINISTIC if it always produces the same result for the same input parameters.

Internationalization (i18n) - process of changing your software so that it isn't hardwired to one language.

Localization (l10n) - the process of adding the appropriate resources to your software so that a particular language/locale is supported.

Tell, Don’t Ask (Law of Demeter).

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

#### Architecture

* Durability
* Utility
* Beauty

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

#### Service Oriented Architecture

SOA is less about how to modularize an application,
and more about how to compose an application by integration
of distributed, separately-maintained and deployed software components.

In SOA, services use protocols that describe how they pass and parse messages using description metadata.

Each SOA building block can be:

* Service provider
* Service broker
* Service requester/consumer

#### Microservices

Microservices is a variant of the SOA
architectural style that structures an application as a collection of loosely coupled services.
It is modern interpretation of SOA used to build distributed software systems.

Microservices-based architectures enable continuous delivery and deployment.

Philosophy: Do one thing and do it well.

The less well you understand a domain,
the harder it will be for you to find proper bounded contexts for your services.

Greenfield development is also quite challenging.
So again, consider starting monolithic first and break things out when you’re stable.

Many of the challenges you’re going to face with microservices get worse with scale.

* Model Around Business Concepts
  (use bounded contexts)
* Adopt a Culture of Automation
  (automated testing, deploy the same way everywhere, continuous delivery)
* Hide Internal Implementation Details
  (modeling bounded contexts, services should also hide their databases, consider using REST)
* Decentralize All the Things
  (teams own their services, align teams to the organization, prefer choreography over orchestration)
* Independently Deployable
  (coexist versioned endpoints, one-service-per-host)
* Isolate Failure
  (expect failure will occur anywhere and everywhere)
* Highly Observable
  (aggregate your logs, aggregate your stats)

Anti-pattern - nanoservice,
is a service whose overhead (communications, maintenance etc.) outweighs its utility.

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

#### IoC:

When high-level module depends on low-level module - you cann't reuse high-level module.

Bob Martin's (uncle Bob) CopyProgram:

````
              +------+
              | Copy |
              +------+
                 /\
                /  \
+---------------+  +------------------+
| Read keyboard |  | Write to printer |
+---------------+  +------------------+
````

Now we need add ability to: Write to disc...

Inversion of control - is used to increase modularity of the program and make it extensible.
<br>Software frameworks, CALLBACKS, schedulers, event loops and dependency injection
are examples of design patterns that follow the inversion of control principle.

IoC serves the following design purposes:

* To decouple the execution of a task from implementation.
* To focus a module on the task it is designed for.
* To free modules from assumptions about how other systems do
  what they do and instead rely on contracts.
* To prevent side effects when replacing a low-level module.

There are several basic techniques to implement inversion of control:

* Using a factory pattern
* Using a service locator pattern
* Using a dependency injection, for example

  * constructor injection
  * setter injection
  * interface injection - define injector method in interface
  * [parameter injection]

* Using a contextualized lookup
* Using template method design pattern
* Using strategy design pattern

Dependency inversion principle -
High-level modules should not depend on low-level modules. Both should depend on abstractions.
IoC - way that we provide this abstraction.

**Interface inversion** - interface must have more than 1 implementation,
not 1 interface for 1 particular class (there is no benefit).
Because in this case 2 classes will have 2 interfaces which may differ
and to solve it we inverse interfaces and create 1 interfaces for all any implementations.

**Flow inversion** (Hollywood Principle - Don't call us, we'll call you).

**Creation inversion** - use factory pattern or service locator or dependency injection.

#### Programming:

Imperative programming is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

Declarative programming focuses on what the program should accomplish
without specifying how the program should achieve the result.

The declarative layer describes what the code will do,
while the implementation layer describes how the code does it.
(The declarative layer is, in effect, a small  domain-specific language).
