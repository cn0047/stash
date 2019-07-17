Theory
-

````
┌─┬─┐
├─┼─┤
│ │ │   ╔═══╗
│ │ │   ║ + ║
│ │ │   ╚═══╝
└─┴─┘
````

Server-Sent Events - servers can initiate data transmission towards clients
once an initial client connection has been established.

IO throughput - data transfer speed in megabytes per second (MB/s or MBPS).

To deal with big complicated problems - have to perform breakdown or decomposition.
One of the best ways to deal with complexity is divide and conquer — split the problem
into smaller problems and solve each one separately.

Bandwidth - count of lines on high-way.
Latency - speed limit on high-way.

Concurrency - single process use multiple threads.
But each thread uses the same resource concurrently.
Parallel - concurrency without using same resource.

Resilience - is the ability of a system to adapt or keep working when challenges occur.

Internationalization (i18n) - process of changing your software so that it isn't hardwired to one language.

Localization (l10n) - the process of adding the appropriate resources to your software so that a particular language/locale is supported.

Tell, Don’t Ask.

Law of Demeter - object A can call method of object B,
but object A should not "reach through" object B to access yet another object C, to request its services.

Consistent Hashing - used in .io

ACL (often simpler): List of users (IDs) and privileges to resource (what can do or see).
RBAC: Permissions based on role.

PC: CPU, Main Memory, Secondary Memory (SSD).

#### TL

Avoid conflicts.
Let offer new ideas.
Ask do people need help.
Help feel okay to admit being wrong.
Know team: skils, strengths, goals, interests.

Bit of delegating & directing and more supporting & coaching.

TL must be:
* like a teacher and have authority to teach team
* like a coach

#### Full stack:

* presentation layer (html, css, js)
* business layer (node, php, etc)
* data access layer (mongo, mysql, etc)

#### Low coupling and high cohesion.

COUPLING refers to the interdependencies between modules.
<br>Tightly coupled code - bad code.

LOW COUPLING is often a sign of a well-structured computer system and a **good** design.
synonym: lose coupling, antonym: coupling.

COHESION describes how related the functions within a single module are.
<br>
Classes should have a small number of instance variables.
In general the more variables a method manipulates the more cohesive that method is to its class.
A class in which each variable is used by each method is maximally cohesive.

Basically, classes are the tightest form of coupling in object-oriented programming.

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

Problems:

* db migrations (slow migrations, few steps migrations)
* `cron`

#### Deployments Strategies

Dark launching - deploying the very first version of a service
so no users available yet.

Highlander (most traditional deployment patten) - all instances are deploying simultaneously.

Canary Deployment - deploys to only a small portion of the available servers.
Some kind of A/B testing.

Rolling Deploy (continuation of the canary deploy) - update one server after another.

Blue-Green - Once you have deployed and fully tested the software in Green,
you switch the router so all incoming requests now go to Green instead of Blue.
Green is now live, and Blue is idle.

Canary with two groups - Blue-Green and add 1 node from new cluster into old one.

Rolling Deploy with two groups - continuation of the canary with two groups.

Don't forget about:
Users may see V1 of a page on one click, then see V2 on refresh and back to V1 on yet another refresh.
As solution you can suggest:
* to use separated site for some users (http://beta.yourcompany.com)
* or use "Feature Toggles"
* or A/B

#### Code quality

Maintainable
Extendable
Reusable
Readable
Understandable
Testable
Documentable
Well designed (patterns)
Follow SOLID
Haven't memory leaks
Haven't vulnerabilities and secure
Cares about backward compatibility

Predictability
Resilience
Elasticity
Fault tolerance

Rigidity
Fragility
Immobility
Complexity

#### Comments

Good comment must explain: `what? why? how?`

#### [Clean Code](https://monosnap.com/file/9UGwycGbfCus8TRIXPjFWGsI2pKOKW)

Clean code always looks like it was written by someone who cares.

Methods should have verb or verb phrase names like postPayment, deletePage, or save.
Accessors, mutators, and predicates should be named for their value and prefixed
with get, set, and is according to the javabean standard.

Flag arguments are ugly. Passing a boolean into a function is a truly terrible practice.
It immediately complicates the signature of the method,
loudly proclaiming that this function does more than one thing.
It does one thing if the flag is true and another if the flag is false!

Noise words are another meaningless distinction.
Imagine that you have a Product class. If you have another called ProductInfo or ProductData,
you have made the names different without making them mean anything different.
Info and Data are indistinct noise words like a, an, and the.

Note that there is nothing wrong with using prefix conventions like a and the
so long as they make a meaningful distinction. For example you might use
a for all local variables and the for all function arguments.
The problem comes in when you decide to call a variable theZork
because you already have another variable named zork.

Noise words are redundant. The word `variable` should never appear in a variable name.
The word `table` should never appear in a table name. How is NameString better than Name?
Would a Name ever be a floating point number? If so, it breaks an earlier rule about disinformation.
Imagine finding one class named Customer and another named CustomerObject.
What should you understand as the distinction? Which one will represent the best path to a customer’s payment history?

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

````
Odd  (1, 3, 5...) -> x % 2 = 1.
Even (2, 4, 6...) -> x % 2 = 0.

3/4 # 3 - numerator; 4 - denominator
x^2 # x - base; 2 - power
````

````
byte          - 8    bits      ~ 
kilobyte (KB) - 1024 bytes     ~ 1 000
megabyte (MB) - 1024 kilobytes ~ 1 000 000
gigabyte (GB) - 1024 megabytes ~ 1 000 000 000
terabyte (TB) - 1024 gigabytes ~ 1 000 000 000 000
petabyte (PB) - 1024 terabytes ~ 1 000 000 000 000 000
exabyte  (EB) - 1024 petabytes ~ 1 000 000 000 000 000 000

millisecond - 0,001
microsecond - 0,000 001
nanosecond  - 0,000 000 001
````

````
uint8  (tinyint)  - 1byte, 0 to 255
uint16 (smallint) - 2byte, 0 to 65535
uint32 (int)      - 4byte, 0 to 4294967295
uint64 (bigint)   - 8byte, 0 to 18446744073709551615

int8              - 1byte, -128 to 127
int16             - 2byte, -32768 to 32767
int32             - 4byte, -2147483648 to 2147483647
int64             - 8byte, -9223372036854775808 to 9223372036854775807

boolean           - 1byte

float             - 4byte, 1.175494351E-38 to 3.402823466E+38
decimal           - 8byte, 2.2250738585072014E-308 to 1.7976931348623158E+308

ASCII character   - 1byte
UTF-8 character   - 1byte (for ASCII equivalents)
UTF-8 character   - 2byte (for special chars)
````

````
nil | null | none | blank | undefined | unknown | void | empty
anonymous |
environments: local; dev; stage|staging; prod|production;
````

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
` Back quote (back tick)
/ Slash (forward slash)
\ Backward Slash
# Pound
$ Dollar
% Percent
& Ampersand
* Asterisk
— Dash (long dash)
- Hyphen
. Dot (period)
@ At Sign
^ Caret
_ Underscore
| Pipe
~ Tilde
≈ Almost Equal
````