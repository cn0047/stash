RTCP - Real-Time Transport Control Protocol
-

RTCP provides feedback on the quality of service (QoS) in media distribution such as:
transmitted octet and packet counts, packet loss, packet delay variation,
and round-trip delay time to participants in a streaming multimedia session.

Typically RTP will be sent on an even-numbered UDP port,
and does not provide any flow encryption or authentication methods.

Message types:
* Sender report.
* Receiver report.
* Source description.
* Goodbye.
* Application-specific message.
