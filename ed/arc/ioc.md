IoC (Inversion of Control).
-

When high-level module depends on low-level module - you can't reuse high-level module.

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

**Ownership inversion** - both high and lower level layers should depend on abstractions
that draw the behavior.
