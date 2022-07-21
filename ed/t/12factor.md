The Twelve-Factor App
-

1. Codebase:
There is **only one codebase per app**.
If there are multiple codebases, it’s not an app – it’s a distributed system.
Each component in a distributed system is an app.
Multiple apps sharing the same code is a violation of 12F.

2. Dependencies:
**Explicitly** declare and isolate dependencies.

3. Config:
Apps sometimes store config as constants in the code - this is a violation of 12F.
Store **config in the environment**.
Env vars are easy to change between deploys without changing any code.
It's language and OS agnostic standard.

4. Backing:
A backing service is any service the app consumes over the network
as part of its normal operation: db, mq systems, caching systems, etc.
and 3d-parties: newrelic, aws s3, twitter, google, etc.
**Code must makes no distinction between local and third party services**.
To the app, both are attached resources.
Resources can be attached and detached.

5. Build, release, run:
Strictly separate build and run stages.
**Separate: build, release and run stages**.
Every release should always have a unique release ID.

6. Processes:
Execute the app as one (or more) **stateless processes**.
12F processes are stateless and share-nothing.
Any data that needs to persist must be stored in a stateful backing service.
Sticky sessions are a violation of 12F and should never be used.

7. Port binding:
Services should make themselves available to other services by **specified ports**.

8. Concurrency:
In 12F processes are a first class citizen.
App must be able to span multiple processes running on multiple physical machines.
The process model truly shines when it comes time to scale out.

9. Disposability:
Maximize robustness with **fast startup and graceful shutdown**.
Processes can be started or stopped at any time.
Processes shut down gracefully when they receive a SIGTERM signal.

10. Dev/Prod parity:
All environments must be as **similar** as possible.

11. Logs:
Treat logs as event streams.

12. Admin Processes:
Admin code must ship **with application code** to avoid synchronization issues.
All admin tasks should be kept in source control and packaged with the application.
