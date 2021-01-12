Functional Programming
-

Important:
Separate mutation from calculation.
Separate function from rule.

Basic principles of functional programming:
* Immutability - once you assign a value to something, that value won't change.
* Disciplined state - shared mutable state is evil, don't use global vars.
* Pure functions.
* First class functions (which can take a function as input or return a function as output).
* Referential transparency - pure functions + immutable data = referential transparency.

**Pure** function:
1. always return the same result for same arguments
2. no mutation or output to I/O etc
3. rely only on input args
4. it isn't void function

Pure function - no side effects!
Pure function can be parallelized!

**Deterministic** function always produces the same result for the same input parameters.

**First class function** - passing function as argument into another function,
assign function to variable, return function from function.
