Networking
-

?DHCP Server

````
| Client         | Name Server     | Root NS   | .info NS      | srv.info NS    |
|                | ns1.isp.com     |           |               |                |
+----------------+-----------------+-----------+---------------+----------------+
| → www.srv.info |                 |           |               |                |
|                | → www.srv.info? |           |               |                |
|                |                 | → *.info? |               |                |
|                |                 | ←         |               |                |
|                | →               | →         | → *.srv.info? |                |
|                |                 | ←         | ←             |                |
|                | →               | →         | →             | → resp with ip |
|                |                 |           |               | to ns1.isp.com |
|                | ←               | ←         | ←             | ←              |
| ←              |                 |           |               |                |
````
