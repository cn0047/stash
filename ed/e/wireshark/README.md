Wireshark
-

Capture filter:

````sh
tcp and port 443

# it's impossible directly capture WebSocket protocols, but
tcp port 8080 # where 8080 - is WS port

host stackoverflow.com
````

Filter:

`frame contains pluralsight.com`

````sh
env SSLKEYLOGFILE='/path/to/sslkey.log'
# and specify this file in Preferences -> Protocols -> SSL
````

HTTPS:
````sh
# 1
SSLKEYLOGFILE=/tmp/https.log open /Applications/Firefox.app

# 2
# Go to wireshark's Preferences -> Protocols -> TLS, and set Pre-Master-Secret log file.
````
