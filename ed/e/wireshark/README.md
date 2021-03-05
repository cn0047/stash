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

````sh
frame contains pluralsight.com
````

HTTPS:

````sh
# 1
SSLKEYLOGFILE=/tmp/https.log open /Applications/Firefox.app

# 2
# Go to wireshark's Preferences -> Protocols -> TLS, and set Pre-Master-Secret log file.
````
