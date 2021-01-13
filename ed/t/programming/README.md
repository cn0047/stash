Programming
-

CSS - fault tolerant.
JS - NOT fault tolerant.

Mixin better than OOP.

#### Process and Thread

Process - any program. It's isolated from other processes.
Process has a virtual address space, executable code, open handles to system objects,
a security context, a unique process identifier, environment variables,
minimum and maximum working set sizes, and at least one thread (primary thread) of execution.

Thread is the segment of a process, run in a shared memory space.
A thread is an entity within a process that can be scheduled for execution.
All threads of a process share its virtual address space and system resources.

## Common stuff

An `expression` evaluates to a value. A `statement` does something.
````sh
y = x + 1   # an expression
print y     # a statement
````

#### Stack and Heap

Stack - function's parameters and local variables allocated on the stack.
If the stack needs to grow then heap operations (allocate new, copy old to new, free old) will occur.

Heap - does not have a single partition of allocated and free regions, set of of free regions.
Unlike the stack, the heap is not owned by one function.
(manipulating the set of free regions in the heap requires synchronization).
OS specifies min heap size (`ulimit` in linux).

## Language

* lexer
* parser
* compiler

## Paradigms

**Imperative** programming - is a programming paradigm that uses statements that change a program's state.
Imperative program consists of commands for the computer to perform.

**Structured** programming - is a programming paradigm
aimed at improving the clarity, quality, and development time of a computer program
by making extensive use of subroutines (procedure, function, method or subprogram),
block structures, for and while loops.

**Procedural** programming - is a programming paradigm, derived from structured programming,
based upon the concept of the procedure call.
(Fortran, Pascal).

**Declarative** programming - focuses on what the program should accomplish
without specifying how the program should achieve the result (SQL, CSS, regex).
<br>The declarative layer describes what the code will do,
while the implementation layer describes how the code does it.
(The declarative layer is, in effect, a small domain-specific language).

**Functional** programming - is a programming paradigm,
a style of building the structure and elements of computer programs
that treats computation as the evaluation of mathematical (pure & deterministic) functions
and avoids changing-state and mutable data.
(JavaScript, Scala).

**Object-oriented** programming - is a programming paradigm based on the concept of objects,
which may contain data, in the form of fields, often known as attributes,
and code, in the form of procedures, often known as methods.

**Event-driven** programming – program control flow is determined by events.
(JavaScript).

**Metaprogramming** - is a programming technique
in which computer programs have the ability to treat programs as their data.
It means that a program can be designed to read, generate, analyse or transform
other programs, and even modify itself while running.

#### Low coupling and high cohesion.

**Coupling** refers to the interdependencies between modules.
<br>Tightly coupled code - bad code.

**Low Coupling** is often a sign of a well-structured computer system and a **good** design.
synonym: lose coupling, antonym: coupling.

**Cohesion** describes how related the functions within a single module are.
<br>
Classes should have a small number of instance variables.
In general the more variables a method manipulates the more cohesive that method is to its class.
A class in which each variable is used by each method is maximally cohesive.

Basically, classes are the tightest form of coupling in object-oriented programming.

#### Code quality

Readable
Understandable
Maintainable
Extendable
Reusable
Testable
Documentable
Well designed (patterns)
Follows SOLID
Don't have memory leaks
Don't have vulnerabilities and security issues
Cares about backward compatibility

Fault tolerance
Predictability
Resilience - is the ability of a system to adapt or keep working when challenges occur.
Elasticity

Rigidity
Fragility
Immobility
Complexity

Self-optimization
Self-protection
Self-adaptation
Self-configuration
Self-healing

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
