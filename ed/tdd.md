TDD
-

Test-driven development.

````
Stub    - an object that provides predefined answers to method calls.
Mock    - an object on which you set expectations.
Fixture - is the fixed state that exists at the start of a test.
````

Better to use stubbing widely for large-scale components and subsystems,
but less for the components at the programming language level,
at this level, we generally prefer mocking.
Mocking allows you to effectively say, "Build me an object that can pretend to be a class of type X."

Tests should act as documentation for the code.

Each test should have just one assert.
But better rule is to think of one coherent feature per test,
which might be represented by up to a handful of assertions.
And better writing tests where each method exercises a unique aspect of the target code’s behavior.

All code should emphasize "what" it does over "how", including test code;

A common technique to isolate tests that use a transactional resource (such as a database)
is to run each test in a transaction which is then rolled back at the end of the test.
The problem with this technique is that it doesn’t test what happens on commit, which is a significant event.
The ORM flushes the state of the objects it is managing in memory to the database.
The database, in turn, checks its integrity constraints.
A test that never commits does not fully exercise how the code under test interacts with the database.

What not to automate with tests:
* Database schema changes - slow...
* Software and hardware configuration changes.

Regression testing - is designed to avoid regressing your application to a previous buggy state.

<ul>
    <li><b>Acceptance</b>: Does the whole system work?</li>
    <li><b>Integration</b>: Does our code work against code we can't change?</li>
    <li><b>Unit</b>: Do our objects do the right thing, are they convenient to work with?</li>
</ul>

## Acceptance

Acceptance criteria come in many different varieties; for one thing, they can be functional or nonfunctional.
Nonfunctional acceptance criteria include things like:
capacity, **performance**, modifiability, **availability**, **security**, usability, and so forth.
The key point here is that when the acceptance tests associated with a particular story or requirement pass.

Acceptance tests are business-facing, not developer-facing.
They test whole stories at a time against a running version of the application in a production-like environment.

The objective of acceptance tests is to prove that our application does
what the customer meant it to, not that it works the way its programmers think it should.

## Functional

Functional testing tests a slice of functionality of the whole system.
Functional testing bases on the specifications of the software.
Functions are tested by feeding them input and examining the output (type of black-box testing).

## Integration

Test multiple units.

## Unit

The aim of a unit test is to show that a single part of the application does what the programmer intends it to.
Unit tests MUST be fast.
