Ports Test Task
-

## Run application

To run application use next command: `make run`.

Now you can run next commands:

````sh
# Check application health:
curl -i "http://localhost:8080/health"

# Create new port:
curl -i -X POST "http://localhost:8080/v1/ports" -d '{
    "id": "UAODS",
    "name": "Odessa",
    "city": "Odessa",
    "country": "Ukraine",
    "alias": ["ua_ods"],
    "regions": ["Europe"],
    "coordinates": [30.7233095, 46.482526],
    "province": "Odessa Oblast",
    "timezone": "Europe/Kiev",
    "unlocs": ["UAODS"],
    "code": "46275"
}'

# Update port:
curl -i -X PUT "http://localhost:8080/v1/ports/UAODS" -d '{
    "id": "UAODS",
    "name": "Odessa",
    "city": "Odessa",
    "country": "Ukraine",
    "alias": ["ua_ods"],
    "regions": ["Europe"],
    "coordinates": [30.7233095, 46.482526],
    "province": "Odessa Oblast",
    "timezone": "Europe/Kiev",
    "unlocs": ["UAODS"],
    "code": "46275"
}'
````

## Tests

To run tests please use next command: `make test`.

Unit test example you can find here: config/config_test.go
Integration test example (with testing business logic) you can find here: app/handler_test.go

## Project structure

DDD layers represented in this application in next way:

````
# Application layer:
app/app.go
app/handler.go
app/router.go

# Domain layer:
app/ports

# Infrastructure layer:
app/storage
````

## Things to discuss

Project has bunch of thing to discuss, each of them marked with `@TBD`.

Also, would be nice to discuss `in memory database` bit more,
because it affected many things in application.

And also worth to mention that `payload` package named in this way intentionally,
this may turn out to be slippery slope for me,
but probably better to have conversation here and explain all thoughts.
