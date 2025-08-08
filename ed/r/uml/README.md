UML
-

### Building blocks:

* Things:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/things.png" width="70%" />
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/messagesAndActions.png" width="70%" />

  * Class.
  * Use case (filled oval).
  * Component (rectangle).
  * Node (cube).

  * Message (solid line with arrow).
  * Return (dashed line with arrow).

  * State (rounded rectangle).
  * Action (filled rounded rectangle).

* Relationships:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/relationships.1.png" width="70%" />

  * Association.
  * Generalization (inheritance) (solid line with not filled arrow)
    (arrow head points on ancestor).
  * Implementation (not filled dashed arrow)
    (arrow head points on interface).
  * Dependency (dashed arrow with open arrow head).

<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/relationships.2.webp" width="70%" />

<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/relationships.3.png" width="70%" />

### Structural Diagrams:

* Class diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.class.png" width="70%" />

  * Basic (solid line).
  * Aggregation (solid line with open diamond)
    (diamond points on parent or container).
  * Composition (solid line with filled diamond)
    (diamond points on parent or owner).
  * Uses (dashed line with arrow)
    (arrow point on class which are using another class (start of arrow)).

* Component diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.component.png" width="70%" />

  * Loosely connected:
    1. Interface
      (solid line with circle (aka lollipop) in 1 end and square (aka port) in 2nd).
    2. Required interface (like lollipop but half of circle (aka socket)).
  * Tightly connected (solid line and squares in both ends).

* Package diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.package.png" width="70%" />

  * Basic (like file folder).

* Deployment diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.deployment.png" width="70%" />

  * Nodes:
    1. Processors: Cubes (db); Stereotyped (any server);
    2. Devices: Network (cloud); Specialty (pc, scanner, etc image);

### Behavioral Diagrams:

* Use case diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.useCase.png" width="70%" />

  * User tasks.
  * System interactions.
  * What not how.

  * Actors (man).
  * Use cases (filled oval).
  * Relationships (solid lines).

* Sequence diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.sequence.png" width="70%" />

  * Lifeline (dashed vertical line with filled square in top).
  * Focus of control.
  * Object Lifetime.

  * Loop (rectangle over the top of sequence diagram).
  * Optional (like loop but with tag opt).
  * Conditional
    (like loop but with tag alt and with dashed line in middle).
  * Parallel (like conditional but with tag par).

* State diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.state.png" width="70%" />

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

* Activity diagram:
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.activity.png" width="70%" />

  * Ations (single step).
  * Activity (multiple step).
  * Special: Initiation (solid circle); Completion (circle with filled circle inside);
  * Flow control:
    1. Decision/Branch (true/false diamond).
    2. Fork and Join (solid line with filled arrow).

### Interaction Diagrams:

* Collaboration diagram.
<img src="https://github.com/cn007b/my/blob/master/ed/r/uml/img/diagram.collaboration.png" width="70%" />
* Sequence diagram.

### Design Diagrams:

* Application diagrams.
