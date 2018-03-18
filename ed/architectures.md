Architectures
-

## Back End

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
