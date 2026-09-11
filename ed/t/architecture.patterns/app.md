Architectures
-

````sh
coordinator
dispatcher
processor
worker

app    #
ddd    # DDD
test   # all tests
bin    # binary files, reserved in golang
src    # source code; reserved in golang
pkg    # reserved in golang
static # images, icons, etc.
dist   # reserved in js
````

The actor model adopts the philosophy that everything is an actor.
Like everything is an object philosophy in object-oriented programming.

## Load Balancer

* Service Discovery.
  When new service starts it must inform LB that it's up and running and what it can do.

* Service Termination.
  Shutdown gracefully - service informs LB that it's gonna stop,
  so LB aware that service no more available.

* Heartbeat Checks.
  Needed to check service's health: LB asks service: "are you here",
  if all ok - service responds: "I'm ok".

Load balancers are generally grouped into two categories:
Layer 7 - Application LB
          (based on application specific data such as HTTP headers, cookies, etc).
Layer 4 - Transport Layer LB.

Methods:
* Least Connection - use the service with the fewest active connections.
* Round Robin - assigns connection to the first service in the list,
  and then moves that service to the bottom of the list.
* Weighted Round Robin.
* Least Response Time - use service with the fewest active connections
  and the lowest average response time.
* Least Bandwidth - use service that is currently serving the least amount of traffic (Mbps).
* URL param or RDP-cookie.

Communication types:
* Unicast - one-to-one.
* Broadcast - one-to-all.
* Multicast - one-to-few.

## Back End

URL routing should be decoupled from the controllers
as such it is possible to swap and replace them easily or update them all at once.
Moreover separated routes easier to configure in terms of middlewares and controllers,
so easier to keep track of middlewares usage and update them all at once.

Controllers must be separated from routing
because you can use controllers not only for HTTP but also for WebSockets etc.

Middlewares must be separated form routing and controllers
with purpose to use them not only for HTTP but also for WebSockets etc.

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

#### Clean Architecture

* Entities.
* Use Cases (Business Logic).
* Controller.
* Framework & Driver.

#### Simple

````
simpleapp
└── 🗂 src
    ├── 🗂 static   # Img, css, etc.
    ├── 🗂 views    # Templates.
    ├── 🗂 dao      # DAO objects.
    └── app.js      # Init. BL.
````

````
log
└── 🗂 src
    ├── 🗂 configs       # All app configs.
    ├── 🗂 middlewares   # CORS, X-Powered-By, etc.
    ├── 🗂 routes        # HTTP controllers. BL.
    ├── 🗂 sockets       # WS controllers. BL.
    ├── 🗂 views         # Templates.
    └── app.js           # Init.
````

````
sandbox-log
└── 🗂 src
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
└── 🗂 src
    └── 🗂 go-app
        ├── 🗂 .gae
        │   ├── app.yaml
        │   └── main.go
        ├── 🗂 common
        ├── 🗂 config
        ├── 🗂 controller
        ├── 🗂 route
        └── 🗂 service
````

````
monitoring.v2
└── 🗂 src
    └── 🗂 go-app
        ├── 🗂 .gae
        ├── 🗂 app
        │   ├── 🗂 config
        │   │   └── 🗂 taxonomy
        │   ├── 🗂 errors
        │   │   ├── 🗂 AppError
        │   │   ├── 🗂 BLError
        │   │   └── 🗂 InvalidVOError
        │   ├── 🗂 routes
        │   └── 🗂 vo
        ├── 🗂 controller
        │   ├── 🗂 ah
        │   ├── 🗂 api
        │   ├── 🗂 cron
        │   ├── 🗂 home
        │   └── 🗂 worker
        ├── 🗂 middleware
        └── 🗂 service
            ├── 🗂 internal
            │   ├── 🗂 cache
            │   ├── 🗂 datastore
            │   └── 🗂 vo
            ├── 🗂 chart
            ├── 🗂 measurement
            ├── 🗂 ping
            ├── 🗂 project
            ├── 🗂 queue
            ├── 🗂 renderer
            └── 🗂 validator
````

````
monitoring.v3
└── 🗂 src
    └── 🗂 go-app
        ├── 🗂 .gae
        ├── 🗂 app
        │   ├── 🗂 config
        │   ├── 🗂 routes
        │   ├── 🗂 middleware
        │   ├── 🗂 controller
        │   ├── 🗂 vo             # optional here
        │   ├── 🗂 errors         # optional here
        │   └── 🗂 taxonomy       # optional here
        └── 🗂 service.v1
            ├── 🗂 internal
            │   ├── 🗂 cache
            │   ├── 🗂 datastore
            │   └── 🗂 vo         # ‼️ private VOs
            ├── 🗂 chart
            ├── 🗂 measurement
            ├── 🗂 ping
            ├── 🗂 project
            ├── 🗂 queue
            ├── 🗂 renderer
            └── 🗂 validator
````

#### DDD

````
wall
└── 🗂 src
    ├── 🗂 app ⓵   # PRESENTATION LAYER + Stuff common for all PHP and JavaScript frameworks + PHP framework itself.
    ├── 🗂 bin ⓶   # All binary files must be hosted here (artisan, console, migration run command, etc).
    ├── 🗂 ddd ⓷   # All stuff related to DDD.
    └── 🗂 web ⓸   # USER INTERFACE LAYER (public stuff).
````

````
wall
└── 🗂 src
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
        ├── 🗂 migrations       # Framework agnostic DB migrations (actual migration instructions).
        └── 🗂 var              # Cache, logs, etc.
````

````
wall
└── 🗂 src
    └── 🗂 ddd ⓷
        └── 🗂 Wall
            ├── 🗂 Application             # 🔰 APPLICATION DDD LAYER.
            │   │                          # Any PHP implementation can work only with this layer.
            │   ├── 🗂 Exception
            │   ├── 🗂 Service
            │   └── 🗂 VO                  # Any request must be represented by VO.
            ├── 🗂 Domain                  # 🔰 DOMAIN DDD LAYER.
            │   ├── 🗂 Model
            │   └── 🗂 Service
            └── 🗂 Infrastructure          # 🔰 INFRASTRUCTURE DDD LAYER.
                ├── 🗂 FullTextSearch
                │   └── 🗂 ElasticSearch
                ├── 🗂 Logging
                └── 🗂 Persistence         # Implements all domain interfaces and returns canonical DTOs as result.
                    ├── 🗂 MongoDB
                    └── 🗂 MySql
````

````
wall
└── 🗂 src
    └── 🗂 web ⓸
        ├── 🗂 css
        │
        ├── 🗂 html
        │   └── 🗂 implementation
        │       ├── 🗂 jquery    # Index page for SPA based on jQuery.
        │       └── 🗂 react     # Index page for SPA based on ReactJS.
        │
        ├── 🗂 js                # FRONTEND.
        │   └── 🗂 implementation
        │       ├── 🗂 jquery    # jQuery scripts.
        │       └── 🗂 react     # ReactJS components, etc.
        │
        ├── 🗂 laravel           # Laravel entry point.
        ├── 🗂 phalcon           # Phalcon entry point.
        ├── 🗂 plainphp          # PlainPHP entry point.
        └── 🗂 symfony           # Symfony entry point.
````

#### Hexagonal

````
prj
├── 🗂 cmd
│   └── main.app
├── 🗂 adapters
└── 🗂 core
    ├── 🗂 ports
    ├── 🗂 usecases
    └── 🗂 domain
````

#### Tests

````
prj
├── 🗂 src
└── 🗂 test
    ├── 🗂 functional
    │   └── 🗂 jmeter
    ├── 🗂 integration
    └── 🗂 unit
        ├── 🗂 fixture
        ├── 🗂 mock
        ├── 🗂 stub
        └── 🗂 prj
````

#### Go

1. Root package is for domain types.
2. Group subpackages by dependency.
3. Use a shared mock subpackage.
4. Main package ties together dependencies.

````
root
├── 🗂 domain
├── 🗂 infrastructure
└── myapp.go

root
├── 🗂 app
│   ├── 🗂 domain
│   └── 🗂 infrastructure
├── 🗂 cmd
└── 🗂 ops # devops

root
├── 🗂 ops
├── 🗂 cmd
│   └── 🗂 app-name
│       └── main.go
├── 🗂 app
│   ├── 🗂 svc1
│   ├── 🗂 svc2
│   └── ...
└── go.mod
````
