Security - Web Application Security Risks
-

Remove `X-Powered-By` header.

To store password in db add salt to password and hash it.

Registration form:
when email is taken - don't say it in error, just say "registration in progress"
and send email with message "hey, you already have account here".

Login form:
When password is incorrect - don't say it in error, just say "credentials incorrect",
so you wont leak information that email exsists but password invalid.
And with each next invalid password login try - increase response delay,
so after 10 tries it will take so much time to submit new request,
so it will stop botnet brute force attack.

Web account settings:
To change email address - send email to origin email adress with link
to verify that person who is changing email is owner of previous email.

Reset password:
Send email with message "please re-log-in first" to verify person.

Log off:
Use expiring token in url to perform log-off.

Cookie:
Don't forget about `HttpOnly, secure` flags.

#### A1 Injection

#### Cross-Site Scripting (XSS)

Occur when:
data enters through an untrusted source;
dynamic content validation.

Categories: stored (database) and reflected (response includes some malicious input).

````
<b onmouseover=alert('Wufff!')>click me!</b>
<img src="http://url.to.file.which/not.exist" onerror=alert(document.cookie);>

<? php
print "Not found: " . urldecode($_SERVER["REQUEST_URI"]);
// http://testsite.test/<script>alert("TEST");</script>
?>
````

FIX:

Filter input escape output.

Use `.innerText` instead of `.innerHtml`
The use of `.innerText` will prevent most XSS problems as it will automatically encode the text.

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

#### Insecure Deserialization

Never deserialize untrusted data
(crossed trust doundary, couldn't have been modified).

#### Using Components with Known Vulnerabilities

#### ~~JSON hijacking~~

In early versions of browsers the JSON file could be loaded as a normal script.

#### Browser security

````
Public-Key-Pins: pin-sha256="<pin-value>"; max-age=<expire-time>; includeSubDomains; report-uri="<uri>"
````

HTTP Strict-Transport-Security (HSTS):
````
Strict-Transport-Security: max-age=<expire-time-in-seconds>; includeSubDomains; preload
````
* client should interact over https
* protects against "downgrade attack"
* relies on "trust on first use" (TOFU)

Subresource Integrity (SRI):
````html
<script src="https://example.com/example-framework.js"
    integrity="sha384-oqVuAfXRKap7fdgcCY5uykM6+R9GqQ8K/uxy9rx7HNQlGYl1kPzQho1wx4JwY8wC"
    crossorigin="anonymous"></script>
````

Content-Security-Policy (CSP):
````html
*|none|self

<meta http-equiv="Content-Security-Policy" content="default-src 'self'; img-src https://*; child-src 'none';">
````
Header:
````
Content-Security-Policy: upgrade-insecure-requests;
````

````
Content-Security-Policy-Report-Only: default-src https:; report-uri /csp-violation-report-endpoint/
````

#### ~~Clickjacking (UI redress attack)~~

The hacker can only send a single click.

For example, imagine an attacker who builds a web site that has a button on it that says "click here for a free iPod".
However, on top of that web page, the attacker has loaded an iframe with your mail account,
and lined up exactly the "delete all messages" button directly on top of the "free iPod" button.

FIX: Header `X-Frame-Options: DENY`.
