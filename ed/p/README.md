Networking
-

![tcp vs udp](https://gist.github.com/cn007b/1859adf8ee58818fb19bd4ec2e9ca78f/raw/23279fd5ca4dc9dec2e3e0320dd51101ff835072/tcpAndUdp.jpeg)

<br>BGP       - Border Gateway Protocol.
<br>CIP       - Common Industrial Protocol (control, safety, synchronization, motion, configuration & information).
<br>DHCP      - Dynamic Host Configuration Protocol (DHCP server can assign dynamic IP address).
<br>STUN      - Session Traversal Utilities for NAT, protocol to discover your public address.
<br>IMAP      - Internet Message Access Protocol.
<br>SFTP      - SSH (also Secure) File Transfer Protocol.
<br>LDAP      - Lightweight Directory Access Protocol.

<br>ISP       - Internet Service Provider.
<br>IP        - Internet Protocol asdress.
<br>MAC       - Media Access Control address.
<br>CDN       - Content Distribution Network.
<br>CIDR      - Classless Inter-Domain Routing.
<br>Cloud DNS - Domain Name System.
<br>ECMP      - Equal-Cost Multi-Path routing.
<br>FQDN      - Fully Qualified Domain Name.
<br>IANA      - Internet Assigned Numbers Authority (assigns addresses to RIR's).
<br>LAN       - Local Area Network.
<br>MPLS      - Multiprotocol Label Switching.
<br>NAT       - Network Address Translation.
<br>RIR       - Regional Internet Registry (assigns space to ISP's).
<br>VPC       - Virtual Private Cloud.
<br>VPN       - Virtual Private Network.
<br>VPS       - Virtual Private Server.
<br>WAN       - Wide Area Network.

802     = LAN technologies.
802.3   = Ethernet.
802.3an = 10Gbps.

Bandwidth - count of lines on high-way.
Latency - speed limit on high-way.

Upstream network traffic - flow of data sent from a client or local network towards a remote server.
Downstream network traffic - flow of data received by a client or local network from a remote server.

The 95th percentile is a widely used mathematical calculation
to evaluate the regular and sustained utilization of a network pipe.
Basically the 95th percentile says that 95% of the time, the usage is below this amount.
So a 95th percentile tells you the value which is greater than or equal to 95% of your data.
High percentiles (95, 99) of response times, also known as tail latencies.

If you take your list of response times and sort it from fastest to slowest,
then the median is the halfway point, for example:
if your median response time is 200 ms, that means half your requests return in less than 200 ms,
and half your requests take longer than that.
The median is also known as the 50th percentile, and sometimes abbreviated as p50.

Proxy:
* forward-proxy: 
  * protects clients online identity.
  * bypass browsing restrictions.
  * blocks access to certain content.
* reverse-proxy:
  * protect website (website's IP hidden from web).
  * LB.
  * caching.
  * handle SSL encription.

````sh
nslookup github.com

0.0.0.0/0 # quad-zero route, /0 - subnet mask, which specifies all networks
::/0      # represents ipv6

0.0.0.0
0.0.0.0:ipv4_
_eth0:ipv4_
````

Network Address Translation (NAT) - is a method of remapping one IP address space into another
by modifying network address information in the IP header.
One Internet-routable IP address of a NAT gateway can be used for an entire private network.
IP masquerading - technique that hides an entire IP address space,
usually consisting of private IP addresses,
behind a single IP address in another, usually public address space.

A record - Returns 32-bit IPv4 address,
most commonly used to map hostnames to an IP address of the host.

AAAA record - Returns IPv6 address.

CNAME - Canonical name record.
CNAME records must always point to another domain name,
never directly to an IP address.
CNAME record creates an alias for single: `www.example.com.  CNAME  example.com.`

DNAME record - Delegation Name record.
Creates an alias for an entire subtree of the domain name tree.
DNAME record creates an alias for subdomains.

ANAME record - ALIAS Name record.
ANAME records are typically configured to point to another domain,
but when queried by a client, answer with an IP address.
ANAME faster that CNAME.

IoT Protocols: MQTT, CoAP (lightweight HTTP).

## OSI model (Open Systems Interconnection)

<img src="https://gist.github.com/cn007b/384d6938ebef985347b29c15476b55c5/raw/9357d5528a6648a6e02b23bb9c8badc3be40fd40/OSIModel.png" width="70%" />

OSI describes 7 layers that computer systems use to communicate over network.

7. **Application Layer** - High-level APIs (**SMTP**, **HTTP**, **FTP**, etc).
6. Presentation Layer - Translation of data between a networking service and an application
   (encoding/compression/encryption/decryption).
5. Session Layer - managing communication sessions.
4. **Transport Layer** -  reliable transmission of data segments between points on a network (**TCP**, **UDP**).
3. Network Layer - structuring and managing a multi-node network,
   including addressing, **routing** and traffic control (betweep 2 IPs).
2. Data Link Layer - reliable transmission of data **frames** between two nodes.
1. Physical Layer - transmission and reception of **raw bit streams** over a physical medium.

## DNS

Domain name - the human-readable address to access websites on the internet.
Domain name formed by the rules and procedures of the DNS.
Subdomain - prefix added before the main domain and separated by dot.

For `www.example.com`: `www` - subdomain, `example` - second-level domain, `.com` - top-level domain.

DNS (Domain Name System) - fundamental system of internet that translates human-readable domain name into IP address.

DNS records types:
* A record.
* AAAA record.
* CNAME record.
* MX record (mail servers for email delivery).
* NS record (authoritative name servers for the domain).
* TXT Record (provides text-based information for verification or configuration purposes).

TLD - Top Level Domain (.info, .com, .net, .org, ...).
Authoritative Name Server.

DNS Lookup:

Client asks Resolver to get IP by URL,
Resolver checks own cache and asks Root NS,
with Root NS reply asks TLD,
with TLD reply asks Authoritative NS,
and store Authoritative NS reply in cache.
DNS Lookup uses DNS protocol.

````
| Client         | Name Server     | Root NS   | Top-Level D   | Authoritative NS |
|                | Resolver        |           |               |                  |
|                | ns1.isp.com     |           | .info NS      | srv.info NS      |
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

DNS_PROBE_FINISHED_NXDOMAIN

## IP

[IP packets](https://monosnap.com/file/o9SgeRZA2bfEuszfKiiSICSeyRtEcH).

## IP Address

IP Address - 2 in 1: host asdress & network asdress.

Classes:
````
A: Devices - Large networks (0.0.0.0 - 127.0.0.0)
             1.2.3.4
             └┘└────┘
              n host

B: Devices - Medium networks (128.0.0.0 - 191.255.0.0)
             129.2.3.4
             └────┘└──┘
              n     host

C: Devices -  Small networks (192.0.0.0 - 223.255.255.0)
             1**.2.3.4
             └──────┘└┘
              network host

D: Multicast (2 - 4.0.0.0 - 239.255.255.255)

E: Reserved
````

Network Address - opposite to broadcast address -> all host bits are turned off (0000 0000).
Broadcast Address - very last IP in local network -> all host bits are turned on (1111 1111).
255.255.255.255 - is also broadcast address (reserved) for local network.

## Internet Protocol Stack

* Application.
* Transport.
* Network (router).
* Link (switch).
* Physical (hub).

Hub - when node sends data to hub, hub replicates data on all other nodes (waste of bandwidth).
<br>Switch - has switch table between switch ports and network nodes MAC addresses.
<br>Router - passes packet from network to network, like bridge between private network and
public network of internet provider.
A packet is typically forwarded from one router to another router through the networks
that constitute an internetwork until it reaches its destination node.

## Firewall

Firewall - hardware or software device that monitors and controls
incoming and outgoing network traffic.
Firewall establishes a barrier between a trusted internal network and untrusted external network.

Network firewalls filter traffic between two or more networks basing on predetermined security rules.
Host-based firewalls run on host computers and control network traffic in and out of those machines.

1st generation  - packet filters: inspect packets transferred between computers.
2nd generation  - stateful filters: maintain knowledge of specific conversations between endpoints
                  by remembering which port number the two IP addresses are using.
3d  generation  - application layer: it can understand certain applications and protocols
                  like ftp, dns, http...
next-generation - (NGFW) wider or deeper inspection at the application layer.

Types:
* Network layer - it's packet filters.
* Application-layer - like a packet filter but apply filtering rules on a per process basis.
* Proxies - proxy server is a gateway from one network to another.
* Network address translation - not only NAT ↑ but also hide the true address of computer.

## Virtual Machine

Virtual Machine may have max 4096 ports.
