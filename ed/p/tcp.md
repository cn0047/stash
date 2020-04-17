TCP - Transmission Control Protocol
-

````
header:                   20 bytes.
connection:               yes (three-way handshake).
in order packet delivery: yes.
retransmission:           yes (for unacknowledged packets).
flow control:             limits the rate a sender transfers data to guarantee reliable delivery.
````

![example](https://gist.githubusercontent.com/cn007b/1859adf8ee58818fb19bd4ec2e9ca78f/raw/23279fd5ca4dc9dec2e3e0320dd51101ff835072/tcp.jpeg)

TCP is reliable, connection oriented protocol with error recovery
and with ordered data delivery.

Connection establishment - TCP handshake (three-way) (requires 3 messages):
1. Clinet sends `SYN (Synchronize)`.
2. Server replies with `SYN-ACK (Synchronize-Acknowledged)`.
3. Client replies with `ACK (Connection Acknowledged)`.

Connection termination  (four-way handshake):
1. Server sends `FIN-ACK`.
2. Clinet replies with `ACK`.
3. Clinet replies with `FIN/ACK`.
4. Server replies with `ACK`.

Drop connection:
1. Server sends `RST (Reset)`.

Segments - the individual units of data transmission that a message is divided into.
Segment is then encapsulated into an Internet Protocol (IP) datagram, and exchanged with peers.
Segment consists of a segment header and a data section.

TCP Header:
* Source port
* Destination port
* Sequence number
* Acknowledgment number
* Data offset (size of the TCP header)
* Reserved
* Flags (aka Control bits)
* Window size
* Checksum (for error-checking of the header)
* Urgent pointer
* Options
* Padding

TCP ports:
* FTP - 20 and 21
* SSH - 22
* TELNET - 23
* SMTP - 25
* HTTP - 80
* HTTP over SSL/TLS - 443
