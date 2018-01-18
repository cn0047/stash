Wireshark
-

Capture filter:

`tcp and port  443`

Filter:

`frame contains pluralsight.com`

````
env SSLKEYLOGFILE='/path/to/sslkey.log'
# and specify this file in Preferences -> Protocols -> SSL
````
