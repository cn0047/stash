WebSocket
-

[tool for ws](https://github.com/zaproxy/zaproxy)

WebSocket is a computer communications protocol,
providing full-duplex (allows communication in both directions) communication channels
over a single TCP connection.

WebSocket handshake uses the `HTTP Upgrade header`
to change from the HTTP protocol to the WebSocket protocol.
WebSocket connection request uses `Upgrade: WebSocket` and `Connection: Upgrade`.
Server replies with the same `Upgrade: WebSocket` and `Connection: Upgrade` headers
and 101 status code and completes the handshake.

WebSocket enables streams of messages on top of TCP.

````sh
ws://
wss:// # TLS
````

Message Types:
* text
* data
* control (ping, pong, close)

Events:
* onopen
* onmessage
* onerror
* onclose

API:
* send
* close

#### Limits

Chrome: 256 total WS Connections & 30 per host.
FireFox: 200 total WS Connections.

#### Security

Check `origin` because hacker can cheat it.

#### Example

Request:
````sh
GET /wss1 HTTP/1.1
Sec-WebSocket-Version: 13
Sec-WebSocket-Key: oGa9MDamyL5gC8sPP1qJzA==
Connection: Upgrade
Upgrade: websocket
Sec-WebSocket-Extensions: permessage-deflate; client_max_window_bits
Host: localhost:5432
````

Response:
````sh
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: bwI9ib/E/FPBJpKWfglHvRJVnDo=

...4.Z.[.?.\.4...(.Y.9.].4....Something from server.
````
