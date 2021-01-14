Swagger
-

[editor](http://editor.swagger.io/)

Swagger is a set of open-source tools built around the OpenAPI Specification
that can help you design, build, document and consume REST APIs.

OpenAPI Specification (formerly Swagger Specification) is an API description format for REST APIs.

`$ref: 'reference to definition'`.
Use `components` for next re-use with `$ref`.

Parameter Types:
* path - `/users/{id}`
* query - `/users?role=admin`
* header - `X-MyHeader: Value`
* cookie (in Cookie header) - `Cookie: debug=0; csrftoken=BUSe35dohU3O1MZvDCU`

Parameter Serialization - translating data structures or object state into a format
that can be transmitted and reconstructed later.

Data Types:
* string (this includes dates and files)
* number (float, double)
* integer (int32, int64)
* boolean
* array
* object

No null type, the nullable attribute is used.
Not data type: Enum, Dictionaries, HashMaps and Associative Arrays. Use object instead.

Mixed Types:
````
oneOf:
  - type: string
  - type: integer
````

[String Formats](https://swagger.io/docs/specification/data-models/data-types/).

`oneOf`, `anyOf`, `allOf`, `not`.

[Supported JSON Schema Keywords](https://swagger.io/docs/specification/data-models/keywords/).

#### YAML

````yaml
swagger: '2.0'
info:
  title: MyApi
  description: MyRestApi
basePath: /api/v1
paths:
  /systems/{systemId}/events:
    get:
      parameters:
        - $ref: '#/parameters/systemId'
        - $ref: '#/parameters/limit'
        - $ref: '#/parameters/offset'
      responses:
        200:
          description: Successfully retrieved events
          schema:
            type: object
            properties:
              items:
                type: array
                items:
                  $ref: '#/definitions/Event'
              count:
                type: number
        404:
          description: System not found
          schema:
            $ref: '#/definitions/Error'
````

#### NodeJS

````javascript
/**
  * @swagger
  * /v1/users:
  *   put:
  *     summary: Creates a new user
  *     description:
  *       "Required roles: `admin`"
  *     tags:
  *       - Users
  *     parameters:
  *       - name: body
  *         in: body
  *         required: true
  *         schema:
  *           type: object
  *           required:
  *             - username
  *             - password
  *           properties:
  *             username:
  *               type: string
  *             password:
  *               type: password
  *           example: {
  *             "username": "someUser",
  *             "password": "somePassword"
  *           }
  *     responses:
  *       200:
  *         schema:
  *           type: object
  *           properties:
  *             id:
  *               type: integer
  *             username:
  *               type: string
  *         examples:
  *           application/json: {
  *             "id": 1,
  *             "username": "someuser"
  *           }
  *       409:
  *         description: When the username is already in use
  */
````
