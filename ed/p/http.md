HTTP (Hypertext Transfer Protocol)
-
2.0
1.1

HTTP 1.1 was standardized in 1997.

HTTP Long Polling - the client polls the server requesting new information.
The server holds the request open until new data is available.
Once available, the server responds and sends the new information.
When the client receives the new information, it immediately sends another request,
and the operation is repeated.

#### Request

````
[method] [URL] [version]
[headers]
[body]
````

Methods:
* CONNECT - client asks an HTTP Proxy server to tunnel the TCP connection.
* HEAD
* GET
* ...

URL:

Length up to 2048 chars.

````
[scheme:][//[userinfo@]host][/]path[?query][#fragment]

┌─ URL ──────────────┐
protocol + host + path + all after question mark
│          └─ URN ─────────────────────────────┤
└─ URI ────────────────────────────────────────┘

URI - Uniform Resource Identifier (protocol + URN).
URL - Uniform Resource Locator (protocol + host + path).
URN - Uniform Resource Name (host + path + all after question mark).
````

Headers:

* Accept
* Accept-Charset
* Accept-Encoding
* Accept-Language
* Authorization
* Content-Type
* Cookie
* If-Modified-Since
* Referer
* Transfer-Encoding - form of encoding used to safely transfer the body payload.
* User-Agent
* X-Forwarded-Proto - identifying the protocol.
* X-Forwarded-Port

````sh
GET /api/collection

Accept: application/x-collection+yaml
Authorization: Basic YWRtaW46cGFzc3dvcmQ=
Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9
Authorization: SHARED-SECRET eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9
Authorization: OAuth oauth_consumer_key="", oauth_nonce="", oauth_signature="", oauth_signature_method="HMAC-SHA1", oauth_timestamp="1471672391", oauth_token="", oauth_version="1.0"
Content-Type: application/json
Transfer-Encoding: chunked|compress|deflate|gzip|identity
X-Forwarded-Proto: https
````

Request with file:
````sh
curl -X POST http://$h/labs/fastly/errors \
  -H "Content-Type: multipart/form-data" -F "file=@/Users/k/Downloads/x.file" -F "msg=MyFile"

Host: localhost:8207
Accept: */*
Content-Length: 325
Content-Type: multipart/form-data; boundary=------------------------bb5595ef8190b841
Expect: 100-continue
User-Agent: curl/7.54.0

--------------------------bb5595ef8190b841
Content-Disposition: form-data; name="file"; filename="x.file"
Content-Type: application/octet-stream

>>> This is file content <<<

--------------------------bb5595ef8190b841
Content-Disposition: form-data; name="msg"

MyFile
--------------------------bb5595ef8190b841--
````

#### Internet media type

[source](http://en.wikipedia.org/wiki/Internet_media_type#List_of_common_media_types5)

The currently registered top-level type names are:
* application:
  ````
  application/javascript
  application/json
  application/octet-stream
  application/pdf
  application/soap+xml
  application/x-www-form-urlencoded
  application/zip
  ````
* audio ("audio/mp4", "audio/mpeg")
* example
* image ("image/png")
* message
* model
* multipart
* text ("text/plain", "text/html", "text/csv", "text/rtf")
* video ("video/avi", "ideo/mpeg")

Currently the following trees are created:
* standard ("application/xhtml+xml", "image/png")
* vendor ("application/vnd.ms-excel", "application/vnd.oasis.opendocument.text")
* personal or vanity
* unregistered "x."

#### Connections

* parallel (in the past, browsers have used multiple TCP connections to issue parallel requests)
* persistent (HTTP 2 connections are persistent, only one connection per origin is required)
* pipeline (multiple HTTP requests are sent on a single TCP connection)

#### Response

````
[version] [status] [reason]
[headers]
[body]
````

Headers:

* Access-Control-Allow-Origin
* Connection (close)
* Location (path for 302 HTTP code)
* Set-Cookie
* Content-Type
* Content-Length
* Cache-Control (private (client's browser), public|max-age|s-maxage (public proxy servers), no-cache)
* Last-Modified
* Expires
* ETag
* X-Powered-By
* X-Frame-Options
* Content-Security-Policy (prevent xss, clickjacking and other code injection attacks)

````sh
Server: nginx/1.10.2

Set-Cookie: AWSELB=3B731DCB1E7B5DC042400ABD5CCF735F3FFEB4F54E46E6B7D...

X-Powered-By: Express

# CORS
Access-Control-Allow-Origin: http://www.example.com
Access-Control-Allow-Origin: *

Content-Type: text/plain; charset=utf-8

# indicate whether or not a browser should be allowed to render a page in a <frame>, <iframe> or <object>.
X-Frame-Options: DENY
````

Cache:

````sh
Cache-Control: private, max-age=0, no-cache

Cache-Control: private, max-age=60
Expires: Thu, 31 Dec 2037 23:55:55 GMT
````

ETag Cache:

````sh
# Request file -> Response:
ETag: "6d82cbb050ddc7fa9cbb659014546e59"

# Next Request file:
If-None-Match: "6d82cbb050ddc7fa9cbb659014546e59"
# if ETag value same -> Response: 304 Not Modified
````

#### CORS (Cross-Origin Resource Sharing)

A CORS preflight request is a CORS request that checks to see if the CORS protocol is understood.
It is an OPTIONS request:
````sh
OPTIONS /resource/foo
Access-Control-Request-Method: DELETE
Access-Control-Request-Headers: origin, x-requested-with
Origin: https://foo.bar.org
````
Response:
````sh
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: OPTIONS, GET, POST, PUT, DELETE
Access-Control-Allow-Headers: *
````

#### Codes

1xx Informational

* 100 **Continue**
* 101 **Switching Protocols**
* 102 **Processing** (WebDAV; RFC 2518)
* 103 Checkpoint (draft POST PUT)
* 105 Name Not Resolved

2xx Success

* 200 **OK**
* 201 **Created**
* 202 **Accepted**
* 203 Non-Authoritative Information (since HTTP/1.1)
* 204 **No Content**
* 205 Reset Content
* 206 Partial Content
* 207 Multi-Status (WebDAV; RFC 4918)
* 208 Already Reported (WebDAV; RFC 5842)
* 226 IM Used (RFC 3229)

3xx Redirection

* 300 Multiple Choices
* 301 **Moved Permanently** (http -> https)
* 302 **Found / Moved Temporary** (location header)
* 303 See Other (since HTTP/1.1) (reidrect in golang)
* 304 **Not Modified**
* 305 Use Proxy (since HTTP/1.1)
* 306 Switch Proxy
* 307 **Temporary Redirect** (Internal) (strict-transport-security) (since HTTP/1.1)
* 308 Permanent Redirect (approved as experimental RFC)[12]

4xx Client Error

* 400 **Bad Request**
* 401 **Unauthorized**
* 402 **Payment Required**
* 403 **Forbidden**
* 404 **Not Found**
* 405 **Method Not Allowed**
* 406 Not Acceptable
* 407 Proxy Authentication Required
* 408 **Request Timeout**
* 409 **Conflict**
* 410 Gone
* 411 Length Required
* 412 Precondition Failed
* 413 Request Entity Too Large
* 414 Request-URI Too Long
* 415 Unsupported Media Type
* 416 Requested Range Not Satisfiable
* 417 Expectation Failed
* 418 *I'm a teapot* (RFC 2324)
* 419 Authentication Timeout (not in RFC 2616)
* 420 Enhance Your Calm (Twitter)
* 420 Method Failure (Spring Framework)
* 422 **Unprocessable Entity** (WebDAV; RFC 4918)
* 423 Locked (WebDAV; RFC 4918)
* 424 Failed Dependency (WebDAV; RFC 4918)
* 424 Method Failure (WebDAV)[14]
* 425 Unordered Collection (Internet draft)
* 426 Upgrade Required (RFC 2817)
* 428 Precondition Required (RFC 6585)
* 429 Too Many Requests (RFC 6585)
* 431 Request Header Fields Too Large (RFC 6585)
* 434 Requested host unavailable.
* 440 Login Timeout (Microsoft)
* 444 No Response (Nginx)
* 449 Retry With (Microsoft)
* 450 Blocked by Windows Parental Controls (Microsoft)
* 451 Redirect (Microsoft)
* 451 Unavailable For Legal Reasons (Internet draft)
* 456 Unrecoverable Error
* 494 Request Header Too Large (Nginx)
* 495 Cert Error (Nginx)
* 496 No Cert (Nginx)
* 497 HTTP to HTTPS (Nginx)
* 499 Client Closed Request (Nginx)

5xx Server Error

* 500 **Internal Server Error**
* 501 **Not Implemented**
* 502 **Bad Gateway**
* 503 **Service Unavailable**
* 504 **Gateway Timeout**
* 505 HTTP Version Not Supported
* 506 Variant Also Negotiates (RFC 2295)
* 507 Insufficient Storage (WebDAV; RFC 4918)
* 508 Loop Detected (WebDAV; RFC 5842)
* 509 Bandwidth Limit Exceeded (Apache bw/limited extension)
* 510 Not Extended (RFC 2774)
* 511 Network Authentication Required (RFC 6585)
* 520 Unknown Error
* 520 Origin Error (Cloudflare)
* 522 Connection timed out
* 523 Proxy Declined Request (Cloudflare)
* 524 A timeout occurred (Cloudflare)
* 598 Network read timeout error (Unknown)
* 599 Network connect timeout error (Unknown)
