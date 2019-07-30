Security
-

[map](http://map.norsecorp.com/)
[xss](https://www.openbugbounty.org/)
[vulnerabilities](http://cve.mitre.org/) [and](https://cwe.mitre.org/)
[twitter](https://twitter.com/MisterRobot245/following/)
[check dns](https://dnsspy.io/)
[check site](https://observatory.mozilla.org/)
[check headers](https://securityheaders.com)
[security platform](https://www.hackerone.com/)

* Using Components with Known Vulnerabilities
* Man-In-The-Middle
* Phishing
* DNS hijacking

## WEB

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

#### Browser security

````sh
Public-Key-Pins: pin-sha256="<pin-value>"; max-age=<expire-time>; includeSubDomains; report-uri="<uri>"
````

HTTP Strict-Transport-Security (HSTS):
````sh
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
````sh
Content-Security-Policy: upgrade-insecure-requests;
````

````sh
Content-Security-Policy-Report-Only: default-src https:; report-uri /csp-violation-report-endpoint/
````
