WebSocket
-

WebSocket is a computer communications protocol,
providing full-duplex (allows communication in both directions) communication channels over a single TCP connection.

WebSocket handshake uses the `HTTP Upgrade header` to change from the HTTP protocol to the WebSocket protocol.
WebSocket connection request uses `Upgrade: WebSocket` and `Connection: Upgrade`.
Server replies with the same `Upgrade: WebSocket` and `Connection: Upgrade` headers and completes the handshake.

WebSocket enables streams of messages on top of TCP.

Events:

* onopen
* onmessage
* onerror
* onclose

API:

* send
* close
