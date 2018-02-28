Architectures
-

## Back End

#### Simple

````
└── src
    ├── 🗂 static     # Img, css, etc.
    ├── 🗂 views      # Templates.
    ├── 🗂 dao        # DAO objects.
    └── app.js        # Init. BL.
````

````
log
└── src
    ├── 🗂 configs     # All app configs.
    ├── 🗂 middlewares # CORS. X-Powered-By.
    ├── 🗂 routes      # HTTP controllers. BL.
    ├── 🗂 sockets     # WS controllers. BL.
    ├── 🗂 views       # Templates.
    └── app.js         # Init.
````

````
http
    controller
    middleware
    request
    response

````

#### DDD

````
wall
└── src
    ├── 🗂 app                           #
    │   ├── 🗂 config                    #
    │   ├── 🗂 implementation            #
    │   │   ├── 🗂 laravel               #
    │   │   ├── 🗂 phalcon               #
    │   │   ├── 🗂 plainphp              #
    │   │   └── 🗂 symfony               #
    │   ├── 🗂 kernel                    #
    │   │   ├── 🗂 Exception             #
    │   │   └── Di.php                   #
    │   ├── 🗂 migrations                #
    │   └── var                          #
    ├── 🗂 bin                           #
    ├── 🗂 ddd                           #
    │   └── Wall                         #
    │       ├── 🗂 Application           #
    │       │   ├── 🗂 Exception         #
    │       │   ├── 🗂 Service           #
    │       │   └── 🗂 VO                #
    │       ├── Domain                   #
    │       │   ├── 🗂 Model             #
    │       │   └── 🗂 Service           #
    │       └── Infrastructure           #
    │           ├── 🗂 FullTextSearching #
    │           │   └── 🗂 ElasticSearch #
    │           ├── 🗂 Logging           #
    │           ├── 🗂 Messaging         #
    │           └── 🗂 Persistence       #
    │               ├── 🗂 CSV           #
    │               ├── 🗂 Mongo         #
    │               └── 🗂 MySql         #
    └── 🗂 web                           #
        ├── 🗂 css                       #
        ├── 🗂 html                      #
        │   └── implementation           #
        │       ├── 🗂 jquery            #
        │       └── 🗂 react             #
        ├── 🗂 js                        #
        │   └── implementation           #
        │       ├── 🗂 jquery            #
        │       └── 🗂 react             #
        ├── 🗂 phalcon                   #
        └── 🗂 plainphp                  #
````
