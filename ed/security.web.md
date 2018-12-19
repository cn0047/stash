Security - Web Application Security Risks
-

To store password in db add salt to password and hash it.

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

#### ~~JSON hijacking~~

In early versions of browsers the JSON file could be loaded as a normal script.

#### Using Components with Known Vulnerabilities
