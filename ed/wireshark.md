Wireshark
-

Capture filter:

`tcp and port  443`

Filter:

`frame contains pluralsight.com`

````sh
env SSLKEYLOGFILE='/path/to/sslkey.log'
# and specify this file in Preferences -> Protocols -> SSL
````
