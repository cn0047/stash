UDP - User Datagram Protocol
-

````
header:                   8 bytes.
connection:               connectionless.
in order packet delivery: no.
retransmission:           no.
flow control:             no.
````

![example](https://gist.github.com/cn007b/1859adf8ee58818fb19bd4ec2e9ca78f/raw/23279fd5ca4dc9dec2e3e0320dd51101ff835072/udp.jpeg)

UDP is lightweight, connectionless, unreliable protocol,
which doesn't care about errors.
UDP is RTP (Real-time Transport Protocol).
UDP is good for broadcast.

Although UDP provides integrity verification (via checksum) of the header and payload,
it provides no guarantees to the upper layer protocol for message delivery
and the UDP layer retains no state of UDP messages once sent.

UDP Header:
* Source port number
* Destination port number
* Length
* Checksum
