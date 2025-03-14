Microservices
-

[patterns](https://microservices.io/patterns/)
[consul](https://www.consul.io/)
[service discovery](https://github.com/etcd-io/etcd)

Microservices is a variant of the SOA
architectural style that structures an application as a collection of loosely coupled services.
It is *modern* interpretation of SOA used to build distributed software systems.
Microservices easier to scale than SOA (because you have scale only 1 service not bunch).

Microservices are independently deployable services modeled around a business domain.

Philosophy: Do one thing and do it well.
Microservice must be stateless.
Microservice must be resilient. And respond to changes quickly.
Microservice must have:
* High cohesion & loose coupling.
* Lightweight communication mechanism.
* Independent data storage.

Microservices-based architectures enable Continuous Delivery and Deployment.

Microservices is known as a "share-nothing" architecture
or at least "share-as-little-as-possible".

Synchronous communication:
* Request/Response.
* HTTP.
* RPC.
* REST.

Asynchronous communication:
* Event based.
* Message queue protocol.

Orchestration - when one microservice supervise and manage else microservices.
<br>Choreography - when microservice independent, and only have to know about message-queue.

The less well you understand domain,
the harder it will be for you to find proper bounded contexts for your services.
In case you don't have clear bounded contexts - do refactoring
with purpose to split stuff into separated modules (same approach for database).
Greenfield development is also quite challenging.
So again, consider starting monolithic first and break things out when you’re stable.

Many of the challenges you’re going to face with microservices get worse with scale.

* Model around business concepts
  (use bounded contexts).
* Adopt a culture of automation
  (automated testing, deploy the same way everywhere, continuous delivery).
* Hide internal implementation details
  (modeling bounded contexts, services should also hide their databases, consider using REST).
* Decentralize all the things
  (teams own their services, align teams to the organization, prefer choreography over orchestration).
* Independently deployable
  (coexist versioned endpoints, one-service-per-host).
* Isolate failure
  (expect failure will occur anywhere and everywhere).
* Highly observable
  (aggregate your logs, aggregate your stats).

In real world app it's *ok* to have **HYBRID-Microservices** architecture.

#### Advantages

* Possibility to use different languages.
* Isolated problems.
* Independent deploy.
* Scalability/resilience.
* Possibility to build/change system faster.
* Responsibilities are clearly defined.
* Easier to oversee and understand.

#### Disadvantages

* Networks are unreliable and slow.
* Required reliable CI/CD.
* Architecture has to be well-thought through from the beginning.
* Duplicate code/data.
* No foreign keys for decoupled DBs.
* Too many programing languages.
* Making components work together.
* Requires more effort in communication.
* Harder to do integration tests.
* Difficult to monitor the whole system.
* Debugging production issues may be hard.
* Logging to one place is challenging.
* Having more and more microservices makes the whole system more complex and harder to oversee the whole operation.
* Hard to see the whole usage graph.
* Github for configs, db migrations, etc.

#### Have to have

* CI/CD.
* Logging/Monitoring/Telemetry.
* Auth (tokens, headers, etc).
* Config (env, other services, etc).
* Routing (healthcheck, discovery).
* Retry/Timeout/Circuit Breaker implementations.
* LB (to scale).

#### Monolith to microservices

Boundary context might be defined by:
* Business feature or UI composition (widget).
* Data.
* Process.

Defining services by applying the decompose by:
* Business capability pattern (courier, restaurant, order management, etc.).
* Sub-domain pattern (DDD bounded context + SRP, CCP).

Business logic organization patterns:
* Transaction script pattern.
  Organize business logic as collection of procedural transaction scripts, one for each type of request.
* Domain model pattern.
  Organize business logic as object model consisting of classes that have state and behavior.
* Domain event.
  An aggregate publishes domain event when it’s created or undergoes some other significant change.

Key questions:
* What are you hoping to achieve?
* Have you considered alternatives to using microservices?
* How will you know if the transition is working?

When might microservices be a bad idea:
* Unclear domain.
* Startup.
* Customer-installed and managed software.
* Not having a good reason.

Pay attention to:
* Shifting structures.
* Reorganizing teams.
* How will you know if the transition is working?

Migration patterns:
* Strangler fig application (HTTP reverse proxy), [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.StranglerFigApplication.png).
* UI composition (page composition, widget composition, micro frontends), [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.UIComposition.png).
* Branch by abstraction, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.BranchByAbstraction.png).
* Parallel run (comparing credit derivative pricing), [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.ParallelRun.png).
* Decorating collaborator, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.DecoratingCollaborator.png).
* Change data capture, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.MigrationPattern.ChangeDataCapture.png).

Decomposing database:
* Shared database, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.SharedDatabase.png).
* Database view, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.DatabaseView.png).
* Database wrapping service, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.DatabaseWrappingService.png).
* Aggregate exposing monolith, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.AggregateExposingMonolith.png).
* Change data ownership, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.ChangeDataOwnership.png).
* Synchronize data in application, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.SynchronizeDataInApplication.png).
* Tracer write, [see](https://gist.github.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/6bca609ed5dae87c29e867db082cb35db8a84b29/microservices.DecomposingDataBase.TracerWrite.png).

Splitting the database:
* Split the database first (repository per bounded context, database per bounded context).
* Split the code first (monolith as data access layer, multischema storage).
* Split database and code together.

#### Info

* Avoid shared database because changes in schema in 1 microservice
leads to redeploy few microservices.

* Avoid client libraries (SDK) (consumer of your service requires client library),
because when you change your service your counsumer must change client library,
also you force user to use specific technology platform.
Solution: use interface.

* Avoid shared libraries because in case of bug you have to re-deploy all microservices
which are using this library. Maybe this library must be separated microservice?

* If two microservices have something common
(code, db, etc) - maybe you have to separate it into another microservice.

* For transactions use BASE transactions
(basic availability, soft state, and eventual consistency):
  1. Two-Phase Commit Protocol.
  2. Eventual consistency (put job in queue and eventually all will be done).

* Avoid **Nanoservice** - anti-pattern,
is a service whose overhead (communications, maintenance etc.) outweighs its utility.

* To bind requests to different microservices -
have to use **Correlation ID** (or **Trace ID**) for each service call
(and provide parent id to recreate whole picture).

* Use timeouts (or circuit breaker) in communication between microservices.

* Use an **exponential backoff** algorithm retries requests exponentially,
increasing the waiting time between retries up to a maximum backoff time.

* Use separate Authentication & Authorization service for auth,
which is connected with API-Gateway (or part of it).

* If you need to deploy microservice and for that you need to deploy another one - you're not
doing microservices, you're doing another kind of services but not microservices.
Microservices must be independently deployable.

* Size of a service is mostly unimportant.
