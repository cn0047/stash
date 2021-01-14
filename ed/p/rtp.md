RTP - Real-time Transport Protocol
-

RTP - protocol for delivering audio and video over IP networks.
RTP typically runs over UDP.
RTP is used in conjunction with the RTCP.
While RTP carries the media streams (audio and video),
RTCP is used to monitor transmission statistics and quality.

`RTP packets` are created at the application layer and handed to the transport layer for delivery.
Each unit of RTP media data created by an application begins with the `RTP packet header`.
