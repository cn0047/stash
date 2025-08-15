Software Development Methodologies
-

The ideal task size is not bigger than 1 day work.

Two-pizza teams - where no team should be so big that it could not be fed with two pizzas.

<img src="https://gist.githubusercontent.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/21a54f296bd9ae086dbfe767a92c8638c54b68ad/peopleConnections.png" width="70%" />

WSJF - Weighted Shortest Job First.

Stakeholder - business/company decision-making person/group/organization.
Shareholder - owner of company's outstanding shares.

Strengths:
* Fast learner.
* Detail oriented.
* Passionate & committed.
* Disciplined and focused.
* Collaborative.
* Flexible & versatile.
* Creative & innovative.
* Loyal & trustworthy.
* Resourceful.
* Commercially aware.

Weakness:
* Detail oriented.
* Difficult to ask other people to help.
* Not good at giving people feedback.
* Hard saying no to people.
* Sometimes too sensitive.
* Get stressed if project runs over deadline.

Software Architect:
* Designing the system structure: components, modules, subsystems.
* Defining technical standards.
* Ensuring system quality attributes: non-functional requirements.
* Collaborating with development teams.

Solution Architect:
* Analyzing business requirements (with stakeholders & BA).
* Designing end-to-end solutions.
* Managing trade-offs: balance conflicting requirements and find optimal solutions.
* Collaborating with multiple teams.

ADR (Architectural Decision Record) - document about why and how Architectural Decision was made.

#### TL

* Avoid conflicts.
* Let offer new ideas.
* Ask do people need help.
* Help feel okay to admit being wrong.
* Know team: skils, strengths, goals, interests.
* Bit of delegating & directing and more supporting & coaching.

TL must be:
* Like a teacher and have authority to teach team.
* Like a coach.

#### KPI (Key Performance Indicators)

* Cycle time (from zero - to prod).
* Burdown.
* Velocity.
* PRs count.

#### SDLC (Software development life-cycle)

* Plan (collect requirements from stakeholders).
* Design (analyze requirements and identify the best solutions).
* Implement (coding).
* Test (automation and manual testing).
* Deploy (production).
* Maintain.

#### Waterfall

Originated in the manufacturing and construction industries.
If the fact changes - it's impossible or at least prohibitively expensive to adapt to changes.

Phases:
````
conception -> initiation -> analysis -> design -> construction -> testing -> deployment -> maintenance
````

conception, initiation = 20–40% of the time invested
coding = 30–40%
rest = rest

#### Agile

The Agile Manifesto is based on 12 principles:
1. Customer satisfaction by rapid delivery of useful software.
2. **Welcome changing requirements, even late in development**.
3. Working software is delivered frequently (weeks rather than months).
4. Close, **daily cooperation between business people and developers**.
5. **Projects are built around motivated individuals**, who should be trusted.
6. *Face-to-face conversation* is the best form of communication (co-location).
7. Working software is the principal measure of progress.
8. Sustainable development, able to maintain a constant pace.
9. Continuous attention to technical excellence and good design.
10. Simplicity—the art of maximizing the amount of work not done—is essential.
11. **Self-organizing teams**.
11. **Regular adaptation to changing circumstance**.

Agile software development methods:
* Adaptive software development (ASD).
* Agile modeling.
* Agile Unified Process (AUP).
* Crystal Clear Methods.
* Disciplined agile delivery.
* Dynamic systems development method (DSDM).
* **Extreme programming** (XP).
* Feature-driven development (FDD).
* Lean software development.
* **Kanban** (development).
* **Scrum**.
* Scrum ban.

Agile practices:
* Acceptance test-driven development (**ATDD**).
* Agile modeling.
* Backlogs (Product and Sprint).
* Behavior-driven development (**BDD**).
* Cross-functional team.
* Continuous integration (**CI**).
* Domain-driven design (**DDD**).
* Information radiators (scrum board, task board, visual management board, burndown chart).
* Iterative and incremental development (**IID**).
* **Pair programming**.
* Planning poker.
* Refactoring.
* **Scrum events**.
* Test-driven development (**TDD**).
* Agile testing.
* Timeboxing.
* Use case.
* User story.
* Story-driven modeling.
* Retrospective (good/bad/continue to do).
* **Velocity tracking**.

#### Extreme Programming

Core Practices:
* Simple Design.
* Test-Driven Development.
* Small Releases.
* Continuous Integration.
* Pair Programming.
* Collective Code Ownership.
* Metaphor (common vision of how the program works).
* Whole Team.
* Coding Standard.
* Design Improvement (Refactoring).
* Planning Game (Release Planning, Iteration Planning).
* Customer Tests.
* Sustainable Pace.

#### Kanban

Method for visualizing the flow of work,
in order to balance demand with available capacity and spot bottlenecks.

Kanban focuses on the customer and work which meets their needs.

General practices:
* Visualization.
* Limiting work in progress.
* Flow management.
* Making policies explicit.
* Using feedback loops.
* Collaborative or experimental evolution.

#### Crystal Clear Methods

* Frequent delivery.
* Reflective improvement.
* Close or osmotic communication.
  Osmotic communication - the team must be in the same room for this to work.
* Personal safety.
  People in team polite and calm anyone can ask any stupid question.
* Focus:
  1. Two-hour period where the developer is to have no interruptions, meetings, long questions, phone call, etc.
  2. Definition of goals should be clear.
* Easy access to expert users.
* Technical environment with automated tests, configuration management, and frequent integration.

#### Scrum

Scrum is an iterative and incremental agile software development framework
for managing product development.

Roles:
* Product Owner: final authority, prioritize backlog.
* Scrum Master: facilitator, negotiator, responsible for guiding team, removes impediments.
* Development Team.

Events:
* Sprint planning (planning poker).
* Grooming (backlog refinement).
* Daily scrum.
* Demo.
* Sprint review and retrospective (what to start/stop/continue to do).

Artifacts:
* Product backlog.
* Management (prioritize product backlog items).
* Sprint backlog.
* Product increment (product backlog items completed during a sprint).
* Extensions:
  * Sprint burn-down chart.
  * Release burn-up chart.
  * Definition of done (DoD).
  * Velocity.

**Velocity** is a measure of the amount of work a team can tackle
during a single sprint and is the key metric in Scrum.
<br>Velocity is equivalent to a specification of an object's
speed and direction of motion (e.g. 60 km/h to the north).

A **story point** is an abstract measure of effort required to implement a user story.

Limitations:
Scrum works less well in the following circumstances:
* Teams whose members are geographically dispersed or part-time.
* Teams whose members have very specialized skills.
* Products with many external dependencies.
* Products that are mature or legacy or with regulated quality control.

User story.
Epic.
Burndown chart.
...

Teams that track "velocity" and "story points" treat development as if it’s linear labor.
Software development is not like moving a pile of stones:
![img](https://cdn-images-1.medium.com/max/2000/1*aTJOF6uQ-jlNuCbCj4BQCg.png)

Work that requires problem solving is more like a **hill**.
There’s an uphill phase where you figure out what you’re doing.
Then when you get to the top you can see down the other side and what it’ll take to finish:
![img](https://cdn-images-1.medium.com/max/2000/1*xV-g3zRDo6Zuu0QfTBGtNQ.png)

The most important question for a team isn’t "what is left?"
but "what is unknown?" Can you see the edges?
Correct question - is **"where is that on the hill?"**

#### Adaptive software development

Continuous learning and adaptation to the emergent state of the project.
The characteristics of an ASD life cycle are that it is:

* Mission focused.
* Feature based.
* Iterative.
* Timeboxed.
* Risk driven.
* Change tolerant.

#### Rapid-application development

RAD approaches to software development put less emphasis on planning
and more emphasis on process.

Focus on business problems that are critical to end users rather than technical problems.

Provides a flexible adaptable process.
The ideal is to be able to adapt quickly to both problems and opportunities.
(There is an inevitable trade-off between flexibility and control, more of one means less of the other).

Good for small and medium-sized projects.
