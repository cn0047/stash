Layers (AKA multitier or 3-tier or n-tier)
-

Layered architecture - monolithic architectural pattern that organizes an application into layers.

* Presentation layer (UI layer (presentation)).
* Application layer (service layer).
* Business logic layer (domain layer).
* Data access layer (persistence (data source) layer).

Classic 3-tier: Presentation -> Business -> Data.

Tier - physical separation: client–server (2-tier system) and the separation is physical.
Layer - logical separation: you don’t have to run layers on different machines.

Disadvantages:
* To add field to display on UI - have to add it to every layer.
* Extra layers can harm performance.

Usage: desktop app, E-commerce, web app.

**Functional layout** - group code by it’s functional type: controllers, models, etc.
But: names are weird (controller.UserController, service.UserService, ...).

**Module layout** - group code by it’s module: user, account, etc.
But: names are weird (users.User, accounts.Controller, should interact with our `users.Controller` in `go` (-> critical err), ...).
