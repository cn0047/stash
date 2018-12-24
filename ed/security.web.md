Security - Web Application Security Risks
-

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
