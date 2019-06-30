UDP - User Datagram Protocol
-

header:                   8 bytes.
connection:               connectionless.
in order packet delivery: no.
retransmission:           no.
flow control:             no.

UDP is lightweight, connectionless, unreliable protocol,
which doesn't care about errors.
UDP is RTP (Real-time Transport Protocol).

Although UDP provides integrity verification (via checksum) of the header and payload,
it provides no guarantees to the upper layer protocol for message delivery
and the UDP layer retains no state of UDP messages once sent.

UDP Header:
* Source port number
* Destination port number
* Length
* Checksum
