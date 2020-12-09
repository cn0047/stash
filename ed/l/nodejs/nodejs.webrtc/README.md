WebRTC (Web Real-Time Communications)
-

[docs](https://webrtc.org)
[docs](https://www.w3.org/TR/webrtc/)
[chrome](chrome://webrtc-internals)

ICE  - Interactive Connectivity Establishment, framework to allow browser to connect with peers (uses STUN and/or TURN).
STUN - Session Traversal Utilities for NAT, protocol to discover your public address.
NAT  - Used to give your device a public IP address.
TURN - Traversal Using Relays around NAT, to open connection with a TURN server and relaying all information through that server.
SDP  - Session Description Protocol, data format used to describe connection that shares media between devices.

Signaling - clients exchange metadata to coordinate communication.

WebRTC works via UDP.
All data transferred using WebRTC is encrypted.
Datagram Transport Layer Security (DTLS) used as encryption.
