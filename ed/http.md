HTTP (Hypertext Transfer Protocol)
-
2.0
1.1

HTTP 1.1 was standardized in 1997.

#### Request

````
[method] [URL] [version]
[headers]
[body]
````

URL:

Length up to 2048 chars.

Headers:

* Authorization
* Accept
* Accept-Charset
* Accept-Encoding
* Accept-Language
* Cookie
* Content-Type
* Referer
* User-Agent
* If-Modified-Since

````
GET /api/collection
Accept: application/x-collection+yaml

Content-Type: application/json

Authorization: Basic YWRtaW46cGFzc3dvcmQ=
````

#### Internet media type

[source](http://en.wikipedia.org/wiki/Internet_media_type#List_of_common_media_types5)

The currently registered top-level type names are:
* application
  ("application/javascript", "application/json", "application/pdf", "application/soap+xml", "application/x-www-form-urlencoded")
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
* Cache-Control (private, public, no-cache)
* Last-Modified
* Expires
* ETag
* X-Powered-By
* X-Frame-Options
* Content-Security-Policy (prevent xss, clickjacking and other code injection attacks)

````
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

A CORS preflight request is a CORS request that checks to see if the CORS protocol is understood.
It is an OPTIONS request using two HTTP request headers: `Access-Control-Request-Method` and `Access-Control-Request-Headers`
and the `Origin` header.

````
OPTIONS /resource/foo 
Access-Control-Request-Method: DELETE 
Access-Control-Request-Headers: origin, x-requested-with
Origin: https://foo.bar.org
````

#### Codes

1xx Informational

* 100 **Continue**
* 101 Switching Protocols
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
* 303 See Other (since HTTP/1.1)
* 304 **Not Modified**
* 305 Use Proxy (since HTTP/1.1)
* 306 Switch Proxy
* 307 Temporary (Internal) Redirect (since HTTP/1.1)
* 308 Permanent Redirect (approved as experimental RFC)[12]

4xx Client Error

* 400 **Bad Request**
* 401 **Unauthorized**
* 402 Payment Required
* 403 **Forbidden**
* 404 **Not Found**
* 405 **Method Not Allowed**
* 406 Not Acceptable
* 407 Proxy Authentication Required
* 408 Request Timeout
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
* 422 Unprocessable Entity (WebDAV; RFC 4918)
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
* 520 Origin Error (Cloudflare)
* 522 Connection timed out
* 523 Proxy Declined Request (Cloudflare)
* 524 A timeout occurred (Cloudflare)
* 598 Network read timeout error (Unknown)
* 599 Network connect timeout error (Unknown)

## HTTP 2

* Data compression (binary)
* Server Push
* Multiplexing
* Pipelining

## HTTPS

https://hstspreload.org/ for add your site to preload.
http://www.badssl.com/ helps to find problems with https.
http://www.httpvshttps.com/

The main point of HTTPS is
authentication of the visited website and protection of the privacy and integrity of the exchanged data.
HTTPS creates a secure channel over an insecure network.

Solves problems with:

* Man-In-The-Middle.
* Phishing
* DNS hijacking

<br>CA - Certification Authority.
<br>SSL - Secure Sockets Layer (SSL 3.0 is vulnerable).
<br>TLS - Transport Layer Security, is modern implementation of SSL.

HTTP Strict Transport Security:
`strict-transport-security` header in response with `max-age` as value
says to browser to reflect to 307 redirect and faster perform secure (not insecure) request.

<meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests">
