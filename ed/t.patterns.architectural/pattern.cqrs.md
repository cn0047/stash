CQRS - Command Query Responsibility Segregation
-

Type 1 - Single database CQRS.

Type 2 - Two database CQRS:

* Commands use write DB
* Queries use read DB

Type 3 - **Event Sourcing**:

Store events -> Replay events -> Modify entity -> Store new event -> Update read DB.

Pros:

* Roint-in-time reconstruction.
* Multiple read DBs.
* Rebuild PROD DB.
