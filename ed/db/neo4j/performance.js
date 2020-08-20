// Performance

// replan=default
// replan=force   - force a replan.
// replan=skip    - f a valid plan already exists, it will be used.
CYPHER replan=force MATCH ...
CYPHER replan=force EXPLAIN MATCH ...

EXPLAIN ... // see the execution plan but not run.
PROFILE ... // run and see which operators are doing most of the work.

// WHERE p.name STARTS WITH '' - may improve performance.
// WITH p                      - berore RETURN, may improve performance.
