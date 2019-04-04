Networking
-

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

IP - Internet Protocol.
DHCP - Dynamic Host Configuration Protocol.

````sh
nslookup github.com

0.0.0.0/0 # quad-zero route, /0 - subnet mask, which specifies all networks
::/0      # represents ipv6
````

Network address translation (NAT) - is a method of remapping one IP address space into another
by modifying network address information in the IP header.

## Virtual Machine

Virtual Machine may have max 4096 ports.
