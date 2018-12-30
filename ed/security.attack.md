Security
-

#### DDos (Denial-of-service attack)

A distributed DDoS is a cyber-attack where the perpetrator uses more than one unique IP address,
often thousands of them.

Distributed autoscale systems may try to cope with DDoS.

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

FIX: Header `X-Frame-Options: DENY`.
