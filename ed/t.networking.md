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

IP   - Internet Protocol.
DHCP - Dynamic Host Configuration Protocol.
CDN  - Content distribution network.
IMAP - Internet Message Access Protocol.
UDP  - User Datagram Protocol (good for broadcast).

The 95th percentile is a widely used mathematical calculation
to evaluate the regular and sustained utilization of a network pipe.
Basically the 95th percentile says that 95% of the time, the usage is below this amount.
So a 95th percentile tells you the value which is greater than or equal to 95% of your data.

````sh
nslookup github.com

0.0.0.0/0 # quad-zero route, /0 - subnet mask, which specifies all networks
::/0      # represents ipv6

0.0.0.0
0.0.0.0:ipv4_
_eth0:ipv4_
````

Network address translation (NAT) - is a method of remapping one IP address space into another
by modifying network address information in the IP header.

A record - Returns a 32-bit IPv4 address,
most commonly used to map hostnames to an IP address of the host.

CNAME - Canonical Name record.
CNAME records must always point to another domain name,
never directly to an IP address.
CNAME record creates an alias for a single:
````
foo.example.com.  CNAME  bar.example.com.
````

DNAME record - Delegation Name record.
Creates an alias for an entire subtree of the domain name tree.
DNAME record creates an alias for subdomains.

ANAME record - ALIAS Name record.
ANAME records are typically configured to point to another domain,
but when queried by a client, answer with an IP address.
ANAME fater that CNAME.

## Virtual Machine

Virtual Machine may have max 4096 ports.
