Architectural Methodologies
-
[awesome](https://github.com/simskij/awesome-software-architecture)
[stakeholder management](https://pubs.opengroup.org/architecture/togaf9-doc/arch/chap21.html)
[architectures ratings](https://gist.githubusercontent.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/53f77869c82fe76f612f2add377ce91e0a481f1c/architecturesRatings.png)

Fitness functions - describe how close an architecture is to achieving an architectural aim:
* Observability.
* Performance.
* Resiliency.
* Security.
* Compliance (regulatory, legal or corporate in a specific business or industry).
* Operability.
Both architects and developers maintain the fitness functions.

ADR (Architectural Decision Record) - document about why and how Architectural Decision was made.

Architect - mix: technology, enterprise architecture, leadership, financial strategy & implementation.

Enterprise Architecture:
* Discipline.
* Process (create -> evolve -> change -> manage architecture).
* Work products (roadmap, org chart, matrices, etc.).

Risk definition: $event due to $cause, which may result in $impact.

Business driver - resource, process or condition that is vital for continued success and growth of business.
Business goal - high-level statement of intent, direction, or desired end state.
Business objective - actionable items needed to achieve business goals.
Business value - generated value like concrete products or services.
Business capability - people, processes, technology.

Stakeholder - individual, group, or organization who may affect, be affected by,
or perceive itself to be affected by a decision, activity, or outcome of a project (someone interested in project).

Stakeholder management approach (power/interest matrix):
Power (Influence)
├──────────────────┬─────────────────┐ HIGH
│ Keep Satisfied   │ Key Players     │
├──────────────────┼─────────────────┤
│ Minimal Effort   │ Keep Informed   │
└──────────────────┴─────────────────┴ Level of interest (impact)
LOW

Requirements:
* Business requirements.
* Stakeholder requirements.
* Transition requirements (migration).

Solution requirements:
* Functional requirements.
* Non-functional requirements.

Architect Focus:
* Business drivers, goals, objectives.
* Stakeholders, business processes, requirements.

RAID - Risks, Assumptions, Issues, Dependencies.

RACI:
* Responsible - who do work.
* Accountable - sign off.
* Consulted - consultant or counsel.
* Informed.

SWOT analysis:
* Strengths.
* Weaknesses.
* Opportunities.
* Threats.

Architecture must care about:
* Durability.
* Utility.
* Beauty.

Modern approaches to enterprise architecture:
* Incremental approach - focus on small transformations.
* Value-driven approach - focus on specific business value.
* Adaptive approach - focus on highly adaptive systems that can evolve as business evolves.

Architectural view model:
* Logical view - functionality that system provides to end-users.
* Process view - run-time behavior of system (concurrency, performance, scalability, etc.).
* Development view - development view (implementation view: packages, components, etc.).
* Physical view - system engineer's point of view (instances topology).
* Scenarios - use cases.

Architectural plan:
1) Introduction and goals (fundamental requirements, for how long system will last).
2) Constraints.
3) Context and scope (external systems and interfaces).
4) Solution strategy (core ideas and approaches).
5) Building block view (structure of source code modularisation).
6) Runtime view (important runtime scenarios).
7) Deployment view (hardware and infrastructure).
8) Crosscutting concepts.
9) Architectural decisions.
10) Quality requirements.
11) Risk and technical depth (known problems and risks).
12) Glossary (ubiquitous language).

Architectural assessment - structured evaluation of software architecture to determine
whether it meets intended business goals, quality attributes, and technical requirements.
* Requirements gathering.
  * Functional requirements.
  * Non-functional requirements.
  * Quality atribute scenarios.
  * Constraints.
  * Architectural concerns risks.
* Design architecture.
  * Attribute Driven Design.
  * C4 Architectural development method.
  * ATAM.
* Analys.
  * Proof of concept.
  * Advanced prototyping.
  * ATAM Risk analysis.

Architectural assessment template:
* Introduction.
* Executive summary.
* Architectural drivers.
* Architecture overview.
* Architecture analysis.
* Recomendations.
* Improvement roadmap.

Architectural Decision View - one of the architecture viewpoints (in ISO/IEC/IEEE 42010 or in TOGAF practice)
that describes the key architectural decisions made during system design.
* Intent.
* Context.
* Representation (visually (decision trees, decision maps) or textually; logical, process, deployment, information).
* Element catalog (decisions themselves: name, description, etc.).
* Interface.
* Behavior.
* Variability (options and alternatives).
* Reasoning.

Architecture Roadmaps:
Iteration model - smal iterations, but each has some business value.
Priority model - square with 4 sections with higher & lower priorities.
Portfolio model - one of:
+------------------+---------------------------+----------------+
|                  | Tactical (small & quick)  | Strategic      |
+------------------+---------------------------+----------------+
| Localized scope  | Non-decisive battles  40% | Blockades  20% |
+------------------+---------------------------+----------------+
| Enterprise scope | Decisive battles      35% | Invasions  50% |
+------------------+---------------------------+----------------+

#### ATAM

ATAM (Architecture Tradeoff Analysis Method) - structured method to evaluate software architecture
by analyzing how well it satisfies quality attributes and the tradeoffs between them.

It helps to answer:
* Will this architecture scale?
* Will it be secure enough?
* Will it be maintainable?
* What risks does it carry?
* What are tradeoffs?

ATAM process steps:
1. Present concept of ATAM to stakeholders.
2. Present business drivers (everyone presents and evaluates business drivers).
3. Present high-level architecture.
4. Identify architectural approaches.
5. Generate quality attribute utility tree (core business/technical requirements,
   and map them to appropriate architectural property).
6. Analyze architectural approaches (each scenario).
7. Brainstorm and prioritize scenarios.
8. Analyze architectural approaches (perform step 6 again with the added knowledge).
9. Present results (provide all documentation).

ATAM drawbacks:
* Expensive and time-consuming.
* Not continuous (it assumes requirements are complete and all scenarios are known up front,
  and assumes one time process with no changes).
* Stakehoders busy, hard to collect them all together, and get all requirements.
* Heavy for Microservices (it originally designed for enterprise SOA).

#### Architecture characteristics

Identify no more than 7 driving characteristics.
Pick the top 3 characteristics (in any order).
Implicit characteristics can become driving characteristics if they are critical concerns.
Add 7 additional characteristics that weren't considered as important, as list of others considered.

**Performance** - amount of time it takes for system to process a business request.
**Responsiveness** - amount of time it takes to get a response to the user.
**Availability** - amount of uptime of system, usually measured in 9's (e.g., 99.9%).
**Fault tolerance** - when fatal errors occur, other parts of system continue to function.
**Scalability** - when number of users/requests increase, responsiveness/performance/errors remain constant.
**Elasticity** - system is able to expand and respond quickly to unexpected/extreme loads (from 20 to 2M users).
**Data integrity** - data across system is correct and there is no data loss in system.
**Data consistency** - data across system is in sync and consistent across databases and tables.

**Adaptability** - how easy system can adapt to changes in environment and functionality.
**Concurrency** - ability of system to process simultaneous requests, usualy in order in which they were received.
**Extensibility** - how easy system can be extended with additional features and functionality.
**Deployability** - level of effort needed for release, the frequency in which releases occur, and the overall risk of deployment.
**Testability** - how ease to test and what is the test coverage.
**Abstraction** - level at which parts of system are isolated from other parts.
**Configurability** - ability of system to support multiple configurations, including custom on-demand.
**Recoverability** - ability of system to start where it left off in the event of system crash.

Implicit:
**security** - ability of system to restrict access to sensitive information or functionality.
**maintainability** - level of effort required to locate and apply changes to system.
**observability** - ability of system to make available and stream metrics (performance, uptime, response times, etc.).
**feasibility** - taking into account timeframes/budgets/developer skills when making architectural choices.

All:
* Accessibility.
* Accountability.
* Accuracy.
* Adaptability.
* Administrability.
* Affordability.
* Agility.
* Auditability.
* Autonomy.
* Availability.
* Compatibility.
* Composability.
* Configurability.
* Correctness.
* Credibility.
* Customizability.
* Debugability.
* Degradability.
* Determinability.
* Demonstrability.
* Dependability.
* Deployability.
* Discoverability.
* Distributability.
* Durability.
* Effectiveness.
* Efficiency.
* Evolvability.
* Extensibility.
* Failure transparency.
* Fault-tolerance.
* Fidelity.
* Flexibility.
* Inspectability.
* Installability.
* Integrity.
* Interchangeability.
* Interoperability.
* Learnability.
* Maintainability.
* Manageability.
* Mobility.
* Modifiability.
* Modularity.
* Operability.
* Orthogonality.
* Portability.
* Precision.
* Predictability.
* Process capabilities.
* Producibility.
* Provability.
* Recoverability.
* Relevance.
* Reliability.
* Repeatability.
* Reproducibility.
* Resilience.
* Responsiveness.
* Reusability.
* Robustness.
* Safety.
* Scalability.
* Seamlessness.
* Self-sustainability.
* Serviceability.
* Supportability.
* Securability.
* Simplicity.
* Stability.
* Standards compliance.
* Survivability.
* Sustainability.
* Tailorability.
* Testability.
* Timeliness.
* Traceability.
* Transparency.
* Ubiquity.
* Understandability.
* Upgradability.
* Usability.
