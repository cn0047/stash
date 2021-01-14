Security Attacks
-

* HTTP Desync Attack (Request Smuggling)
* Shellshock

* Regular expression Denial of Service (ReDoS)

#### Server-Site Request Forgery (SSRF)

````php
<?php
$image = fopen($_GET['url'], 'rb');
````
and
````sh
GET /?url=http://localhost/server-status HTTP/1.1
GET /?url=file:///etc/passwd HTTP/1.1
GET /?url=dict://localhost:11211/stat HTTP/1.1
````

FIX:
* Use whitelist for internal services.
* Use only http & https (ports: 80 & 443).

#### Cross-Site Request Forgery (CSRF)

AIM: Force end user to execute unwanted actions (submitting a malicious request)
on a web application in which they're currently authenticated.
Perform action on the victim's behalf.

````
<a href="http://bank.com/transfer.do?acct=MARIA&amount=100000">View my Pictures!</a>

<form action="http://bank.com/transfer.do" method="POST">
<input type="hidden" name="acct" value="MARIA"/>
<input type="hidden" name="amount" value="100000"/>
<input type="submit" value="View my pictures"/>
</form>
````

FIX: CORS (if browser support CORS) or CSRF token.

#### Cross-Site Scripting (XSS)

Occur when:
data enters through an untrusted source;
dynamic content validation.

Categories:
* stored (database)
* reflected (response includes some malicious input)
* DOM - malicious data does not touch the web server
* self - user has own xss on own social page or something like that

````sh
w.Header().Set("X-XSS-Protection", "0")

<b onmouseover=alert('Wufff!')>click me!</b>
<img src="http://url.to.file.which/not.exist" onerror=alert(document.cookie);>

<? php
print "Not found: " . urldecode($_SERVER["REQUEST_URI"]);
// http://testsite.test/<script>alert("TEST");</script>
?>
````
````
<noscript><p title="</noscript><img src=x onerror=alert(1)>">
````

FIX:
Filter input escape output.
Use `.innerText` instead of `.innerHtml`
The use of `.innerText` will prevent most XSS problems as it will automatically encode the text.

#### A1 Injection

SQL Injection and whatnot.

#### Directory (path) traversal attack

AKA: ../ (dot dot slash) attack.
AIM: gain unauthorized access to the file system.

It is exploiting insufficient security validation / sanitization of user-supplied input file names.

````php
<?php
$template = 'red.php';
if (isset($_COOKIE['TEMPLATE']))
   $template = $_COOKIE['TEMPLATE'];
include ("/home/users/phpguru/templates/" . $template);
````
````
Cookie: TEMPLATE=../../../../../../../../../etc/passwd
````

FIX: Query string is usually URI decoded before use.

#### ~~Clickjacking (UI redress attack)~~

The hacker can only send a single click.

For example, imagine an attacker who builds a web site that has a button on it that says "click here for a free iPod".
However, on top of that web page, the attacker has loaded an iframe with your mail account,
and lined up exactly the "delete all messages" button directly on top of the "free iPod" button.

FIX: Headers `X-Frame-Options: DENY, X-Frame Options: DENY, X-Frame Options: SAMEORIGIN`.

#### Insecure Deserialization

Never deserialize untrusted data
(crossed trust doundary, couldn't have been modified).

#### ~~JSON hijacking~~

In early versions of browsers the JSON file could be loaded as a normal script.

#### ~~XML External Entity (XXE)~~

Is a type of attack against an application that parses XML input,
occurs when XML input containing a reference to an external entity
is processed by a weakly configured XML parser.

#### DDos (Denial-of-service)

A distributed DDoS is a cyber-attack where the perpetrator uses more than one unique IP address,
often thousands of them.

Distributed autoscale systems may try to cope with DDoS.
