Linux
-

````sh
# nginx
telnet localhost 8080

# php-fpm
telnet localhost 9000

# websocket
telnet 0.0.0.0 12345
````

````
# HTTPS keys:
openssl genrsa 1024 > private.key
openssl req -new -key private.key -out cert.csr
openssl x509 -req -in cert.csr -signkey private.key -out certificate.pem
````

sudo dpkg -i {name}

On Linux systems, a kernel configuration parameter called `net.core.somaxconn`
provides an upper limit on the value of the backlog parameter passed to the listen function
that is used to create the servers listening socket.
If the backlog argument is greater than the value in /proc/sys/net/core/somaxconn,
then it is silently truncated to that value.
The default value in this file is 128.

````sh
/dev/null                               # stream, hide output
/dev/stderr                             # stream 2
/dev/stdin                              # stream 0
/dev/stdout                             # stream 1
/etc                                    # system configuration directory
/etc/os-release                         # about linux
/etc/pam.d/common-session               # ?
/etc/pam.d/common-session-noniteractive # ?
/etc/security/limits.conf               # ?
/etc/sysctl.conf                        #
/etc/ttys                               # logged in users
/opt                                    # optional software directory
/proc/cpuinfo                           # info about cpu
/proc/sys/net/core/somaxconn
/proc/sys/net/ipv4/tcp_keepalive_time # current tcp_keepalive_time value
/sys/devices/system/cpu/cpu[0-7]
````

````sh
cat /proc/sys/net/core/somaxconn
sysctl -n net.core.somaxconn
# on server must be at least 1024
````

````sh
cat /etc/passwd
sar:x:205:105:Stephen Rago:/home/sar:/bin/ksh

colon-separated fields:
sar          - the login name,
x            - encrypted password,
205          - numeric user ID
105          - numeric group ID
Stephen Rago - a comment field
/home/sar    - home directory
/bin/ksh     - and shell program
````

Each directory contains subdirectories `.` and `..`.
<br>`dot` refers to the current directory,
<br>`dot-dot` refers to the parent directory,

````
cc  - the C compiler,
gcc - the GNU C compiler,
````

#### keys

````
Ctrl-k # delete rest of line
Ctrl-u # delete to start of line
Ctrl-w # delete word
````
