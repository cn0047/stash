Architectures
-

coordinator
dispatcher
processor
worker

## Load Balancer

* Service Discovery.
  When new service starts it must inform LB that it's up and running and what it can do.

* Service Termination.
  Shutdown gracefully - service informs LB that it's gonna stop,
  so LB aware that service no more available.

* Heartbeat Checks.
  Needed to check service's health: LB asks service: "are you here",
  if all ok - service responds: "I'm ok".

## Back End

URL routing should be decoupled from the controllers
as such it is possible to swap and replace them easily.

`model` layer (directory) is ambiguous because:
* it's unclear what is model in reach app (mysql, mongo, redis, elasticsearch, etc)?
* which model should contain stuff common for 2 models?
* where to place infrastructural stuff (doctrine annotations, etc)?
Example: [one](https://monosnap.com/file/xa6jUhfGdo6fCFDeHDLjZ82GU9seJ6)

Q: Why do you need `service` for simple MVC CRUD project?
A: To add cache; To select DB (master, slave);

Q: Why do you need SPA for admin?
A: Avoid data loss for expired session when handling submit POST form.

Functional organization:
Create directories: Customers, Products, Vendors, ... instead of Controllers, Models, Views.
Pros: easy to navigate.
Cons: lose framework convention.

#### Functional layout

Group code by it’s functional type: controllers, models, etc.

But:
* names are atrocious (controller.UserController, service.UserService, ...).

#### Module layout

Group code by it’s module: user, account, etc.

But:
* terrible names like: (users.User, ...)
  or `accounts.Controller` needs to interact with our `users.Controller` in `go` (-> critical err).

#### Clean Architecture

* Entities
* Use Cases (Business Logic)
* Controller
* Framework & Driver

#### Monolith

Not so awful in case of server-side-rendering.

But:
* [10K SLOC](https://en.wikipedia.org/wiki/Source_lines_of_code).
* Extremely difficult to navigate the code and isolate your code.

#### Simple

````
mongomart
└── src
    ├── 🗂 static   # Img, css, etc.
    ├── 🗂 views    # Templates.
    ├── 🗂 dao      # DAO objects.
    └── app.js      # Init. BL.
````

````
log
└── src
    ├── 🗂 configs       # All app configs.
    ├── 🗂 middlewares   # CORS. X-Powered-By.
    ├── 🗂 routes        # HTTP controllers. BL.
    ├── 🗂 sockets       # WS controllers. BL.
    ├── 🗂 views         # Templates.
    └── app.js           # Init.
````

````
sandbox-log
└── src
    ├── 🗂 di
    │   ├── config.go
    │   └── container.go
    ├── 🗂 http
    │   ├── 🗂 controller
    │   ├── 🗂 middleware
    │   ├── 🗂 request
    │   └── 🗂 response
    ├── 🗂 service
    │   ├── 🗂 v1
    │   └── 🗂 v2
    └── main.go
````

````
monitoring.v1
└── src
    └── go-app
        ├── .gae
        │   ├── app.yaml
        │   └── main.go
        ├── common
        ├── config
        ├── controller
        ├── route
        └── service
````

````
monitoring.v3
└── src
    └── go-app
        ├── .gae
        ├── app
        │   ├── config
        │   │   └── taxonomy
        │   ├── errors
        │   │   ├── AppError
        │   │   ├── BLError
        │   │   └── InvalidVOError
        │   ├── routes
        │   └── vo
        ├── controller
        │   ├── ah
        │   ├── api
        │   ├── cron
        │   ├── home
        │   └── worker
        ├── middleware
        └── service
            ├── chart
            ├── internal
            │   ├── cache
            │   ├── datastore
            │   └── vo
            ├── measurement
            ├── ping
            ├── project
            ├── queue
            ├── renderer
            └── validator
````

````
monitoring.next
└── src
    └── go-app
        ├── .gae
        ├── app
        │   ├── config
        │   ├── taxonomy
        │   ├── errors
        │   ├── routes
        │   ├── middleware
        │   ├── controller
        │   └── vo
        └── service
            ├── internal
            │   ├── cache
            │   ├── datastore
            │   └── vo
            ├── chart
            ├── measurement
            ├── ping
            ├── project
            ├── queue
            ├── renderer
            └── validator
````

#### DDD

````
wall
└── src
    ├── 🗂 app ⓵   # PRESENTATION LAYER + Stuff common for all PHP and JavaScript frameworks + PHP frameworks.
    ├── 🗂 bin ⓶   # All binary files must be hosted here (artisan, console, migration, etc).
    ├── 🗂 ddd ⓷   # All stuff related to DDD.
    └── 🗂 web ⓸   # USER INTERFACE LAYER (public stuff).
````

````
wall
└── src
    └── 🗂 app ⓵
        ├── 🗂 config           # All project's configs.
        │                       # Any particular PHP implementation must use these configs.
        ├── 🗂 implementation   # PRESENTATION LAYER.
        │   ├── 🗂 laravel
        │   ├── 🗂 phalcon
        │   ├── 🗂 plainphp
        │   └── 🗂 symfony
        ├── 🗂 kernel
        │   ├── 🗂 Exception    # Kernel exceptions.
        │   └── Di.php          # Simple DIC container.
        │                       # One for any PHP framework implementation (with purpose to support DRY).
        │                       # This DIC also performs common stuff like init bridges, init facades
        │                       # with custom logic which is common for all PHP implementations.
        ├── 🗂 migrations       # Framework agnostic DB migrations.
        └── var                 # Cache, logs, etc.
````

````
wall
└── src
    └── 🗂 ddd ⓷
        └── Wall
            ├── 🗂 Application             # 🔰 APPLICATION DDD LAYER.
            │   │                          # Any PHP implementation can work only with this layer.
            │   ├── 🗂 Exception
            │   ├── 🗂 Service
            │   └── 🗂 VO                  # Any request must be represented by VO.
            ├── Domain                     # 🔰 DOMAIN DDD LAYER.
            │   ├── 🗂 Model
            │   └── 🗂 Service
            └── Infrastructure             # 🔰 INFRASTRUCTURE DDD LAYER.
                ├── 🗂 FullTextSearching
                │   └── 🗂 ElasticSearch
                ├── 🗂 Logging
                └── 🗂 Persistence         # Implements all domain interfaces and returns canonical DTOs as result.
                    ├── 🗂 MongoDB
                    └── 🗂 MySql
````

````
wall
└── src
    └── 🗂 web ⓸
        ├── 🗂 css
        │
        ├── 🗂 html
        │   └── implementation
        │       ├── 🗂 jquery    # Index page for SPA based on jQuery.
        │       └── 🗂 react     # Index page for SPA based on ReactJS.
        │
        ├── 🗂 js                # FRONTEND.
        │   └── implementation
        │       ├── 🗂 jquery    # jQuery scripts.
        │       └── 🗂 react     # ReactJS components, etc.
        │
        ├── 🗂 laravel           # Laravel entry point.
        ├── 🗂 phalcon           # Phalcon entry point.
        ├── 🗂 plainphp          # PlainPHP entry point.
        └── 🗂 symfony           # Symfony entry point.
````

#### Tests

````
prj
├── src
└── test
    ├── functional
    │   └── jmeter
    ├── integration
    └── unit
        ├── fixture
        ├── mock
        ├── stub
        └── prj
````

#### Go

1. Root package is for domain types.
2. Group subpackages by dependency.
3. Use a shared mock subpackage.
4. Main package ties together dependencies.

````
root
├── domain
├── infrastructure
└── myapp.go
````
