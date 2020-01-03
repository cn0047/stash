#!/usr/bin/expect -f

spawn telnet localhost 8080
expect "Escape character is"
send -- "GET /x=1 HTTP/1.1\n"
send -- "Host: localhost:8080\n\r"
expect eof
