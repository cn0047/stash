Networking
-

?DHCP Server

````
| Client         | Name Server     | Root NS   | .info NS      | srv.info NS      |
|                | ns1.isp.com     |           |               |                  |
+----------------+-----------------+-----------+---------------+------------------+
| www.srv.info ---->                                                              |
|                  www.srv.info? ---->                                            |
|                                    *.info? ---->                                |
|                                                *.srv.info? ---->                |
|                                                <-------------- resp with ip     |
|                                                                to ns1.isp.com   |
|                                    <------------                                |
|                  <------------------                                            |
| <-----------------                                                              |
````

````
nslookup github.com
````

Network address translation (NAT) - is a method of remapping one IP address space into another
by modifying network address information in the IP header.

## Virtual Machine

Virtual Machine may have max 4096 ports.
