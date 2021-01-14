JSON API
-
v1.0

[jsonapi](http://jsonapi.org/)

JSON API is a format that works with HTTP.
It delineates how clients should request data from a server, and how the server should respond.

JSON API requires use of the JSON API media type (`application/vnd.api+json`) for exchanging data.

#### Document Structure

A JSON `object` MUST be at the root of every JSON API request and response containing data:

* `data`: the document’s "primary data"
* `errors`: an array of error objects
* `meta`: a meta object that contains non-standard meta-information.

The members data and errors MUST NOT coexist in the same document.

#### Resource Object

A resource object MUST contain at least the following top-level members:

* `id`
* `type`

#### Compound Documents

To reduce the number of HTTP requests, servers MAY allow responses that include related resources
along with the requested primary resources. Such responses are called “compound documents”.

In a compound document, all included resources MUST be represented as an array of resource objects
in a top-level `included` member.

#### Pagination

Pagination links MUST appear in the `links` object that corresponds to a collection.

The following keys MUST be used for pagination links:

* `first`: the first page of data
* `last`: the last page of data
* `prev`: the previous page of data
* `next`: the next page of data

#### Fetching Data

Multiple related resources can be requested in a comma-separated list:

````sh
GET /articles/1?include=author,comments.author HTTP/1.1
Accept: application/vnd.api+json
````

A client MAY request that an endpoint return only specific fields:

````sh
GET /articles?include=author&fields[articles]=title,body&fields[people]=name HTTP/1.1
Accept: application/vnd.api+json
````

#### Creating Resources

`POST` method.

The response SHOULD include a `Location` header identifying the location of the newly created resource.

The response MUST also include a document that contains the primary resource created.

#### Updating Resources

`PATCH` method.

#### Asynchronous Processing

````sh
POST /photos HTTP/1.1
````

The request SHOULD return a status 202 Accepted with a link in the Content-Location header.

````sh
HTTP/1.1 202 Accepted
Content-Type: application/vnd.api+json
Content-Location: https://example.com/photos/queue-jobs/5234

{
  "data": {
    "type": "queue-jobs",
    "id": "5234",
    "attributes": {
      "status": "Pending request, waiting other process"
    },
    "links": {
      "self": "/photos/queue-jobs/5234"
    }
  }
}
````

To check the status of the job process:

````sh
GET /photos/queue-jobs/5234 HTTP/1.1
Accept: application/vnd.api+json
````

Requests for still-pending jobs SHOULD return a status 200 OK,
optionally, the server can return a `Retry-After` header.

#### Relationships

Relationships may be to-one or to-many.
`null` for empty to-one relationships.
an empty array (`[]`) for empty to-many relationships.

`links, data, meta`

````js
{
  "type": "articles",
  "id": "1",
  "attributes": {
    "title": "Rails is Omakase"
  },
  "relationships": {
    "author": {
      "links": {
        "self": "http://example.com/articles/1/relationships/author",
        "related": "http://example.com/articles/1/author"
      },
      "data": { "type": "people", "id": "9" }
    }
  },
  "links": {
    "self": "http://example.com/articles/1"
  }
}
````
