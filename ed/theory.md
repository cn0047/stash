Theory
-

nanosecond  - 0,000 000 001
microsecond - 0,000 001
millisecond - 0,001

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

Internationalization (i18n) - process of changing your software so that it isn't hardwired to one language.

Localization (l10n) - the process of adding the appropriate resources to your software so that a particular language/locale is supported.

Tell, Don’t Ask.

Law of Demeter - object A can call method of object B,
but object A should not "reach through" object B to access yet another object C, to request its services.

#### Security

Symmetric encryption (symmetric key cryptography) - sender and receiver have same secret key.
Same key for encrypting and decrypting.

Public key cryptography - sender have public key, receiver have private key.
It is expensive type of cryptography.

#### Full stack:

* presentation layer (html, css, js)
* business layer (node, php, etc)
* data access layer (mongo, mysql, etc)

#### Aggregation and Composition

Aggregation - child can exist independently of the parent (Car -> Tires).

Composition - child can NOT exist independent of the parent (House -> Room).

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
* Haven't vulnerabilities and secure
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

#### Simple stuff

`camelCase`
`kebab-case`
`snake_case`

Even (2, 4, 6...) - `x % 2 = 0`.
Odd (1, 3, 5...) - `x % 2 = 1`.

````
( Open Parenthesis
) Close Parenthesis
[ Open Bracket
] Close Bracket
{ Open Curly Bracket
} Close Curly Bracket
< Open Angle Bracket
> Close Angle Bracket
! Exclamation Mark
? Question Mark
' Single Quote
" Double Quote
` Back quote
/ Slash (forward slash)
\ Backward Slash
# Pound Sign
$ Dollar Sign
% Percent Sign
& Ampersand
* Asterisk
- Dash
. dot
@ At Sign
^ Caret
_ Underscore
| Pipe
~ Tilde
````
