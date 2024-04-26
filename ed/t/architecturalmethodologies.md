Architectural Methodologies
-
[stakeholder management](https://pubs.opengroup.org/architecture/togaf9-doc/arch/chap21.html)

Enterprise Architecture:
* Discipline.
* Process (create -> evolve -> change -> manage architecture).
* Work products (roadmap, org chart, matrices, etc.).

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
* business requirements.
* stakeholder requirements.
* transition requirements (migration).

Solution requirements:
* Functional requirements.
* Non-functional requirements.

Architect Focus:
* business drivers, goals, objectives;
* stakeholders, business processes, requirements;

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

Architectural view model:
* Logical view - functionality that system provides to end-users.
* Process view - run-time behavior of system (concurrency, performance, scalability, etc.).
* Development view - development view (implementation view: packages, components, etc.).
* Physical view - system engineer's point of view (instances topology).
* Scenarios - use cases.

Architectural plan:
1) Introduction and goals (fundamental requirements).
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
