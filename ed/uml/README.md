UML
-

### Building blocks:

* Things ([example 1](https://github.com/cn007b/my/blob/master/ed/uml/things.png), [example 2](https://github.com/cn007b/my/blob/master/ed/uml/messagesAndActions.png)):

![things](https://github.com/cn007b/my/blob/master/ed/uml/things.png)
![messagesAndActions](https://github.com/cn007b/my/blob/master/ed/uml/messagesAndActions.png)

  * Class.
  * Use case (filled oval).
  * Component (rectangle).
  * Node (cube).

  * Message (solid line with arrow).
  * Return (dashed line with arrow).

  * State (rounded rectangle).
  * Action (filled rounded rectangle).

* [Relationships](https://github.com/cn007b/my/blob/master/ed/uml/relationships.png):

  * Association.
  * Generalization (inheritance) (solid line with not filled arrow)
    (arrow head points on ancestor).
  * Implementation (not filled dashed arrow)
    (arrow head points on interface).
  * Dependency (dashed arrow with open arrow head).

### Structural Diagrams:

* **[Class diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.class.png)**:

  * Basic (solid line).
  * Aggregation (solid line with open diamond)
    (diamond points on parent or container).
  * Composition (solid line with filled diamond)
    (diamond points on parent or owner).
  * Uses (dashed line with arrow)
    (arrow point on class which are using another class (start of arrow)).

* [Component diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.component.png):

  * Loosely connected:
    1. Interface
      (solid line with circle (aka lollipop) in 1 end and square (aka port) in 2nd).
    2. Required interface (like lollipop but half of circle (aka socket)).
  * Tightly connected (solid line and squares in both ends).

* [Package diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.package.png):

  * Basic (like file folder).

* [Deployment diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.deployment.png):

  * Nodes:
    1. Processors: Cubes (db); Stereotyped (any server);
    2. Devices: Network (cloud); Specialty (pc, scanner, etc image);

### Behavioral Diagrams:

* [Use case diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.useCase.png):
  
  * User tasks.
  * System interactions.
  * What not how.

  * Actors (man).
  * Use cases (filled oval).
  * Relationships (solid lines).

* **[Sequence diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.sequence.png)**:

  * Lifeline (dashed vertical line with filled square in top).
  * Focus of control.
  * Object Lifetime.

  * Loop (rectangle over the top of sequence diagram).
  * Optional (like loop but with tag opt).
  * Conditional
    (like loop but with tag alt and with dashed line in middle).
  * Parallel (like conditional but with tag par).

* [State diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.state.png):

  * States:
    * Basic (rounded rectangle).
    * Internal behavior (rounded rectangle with: entry-do-exit).
    * Special:
      1. Initial - underlined solid circle.
      2. Final - underlined circle with filled circle inside.

  * Transitions:
    * Basic (solid open arrow).
    * Transition (solid open arrow with caption).
    * Event (like transition).

  * Composite state - state inside state.

* [Activity diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.activity.png):

  * Ations (single step).
  * Activity (multiple step).
  * Special: Initiation (solid circle); Completion (circle with filled circle inside);
  * Flow control:
    1. Decision/Branch (true/false diamond).
    2. Fork and Join (solid line with filled arrow).

### Interaction Diagrams:

* [Collaboration diagram](https://github.com/cn007b/my/blob/master/ed/uml/diagram.collaboration.png).
* Sequence diagram.
