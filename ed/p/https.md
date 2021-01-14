HTTPS - HyperText Transfer Protocol Secure 
-

[to add your site to preload](https://hstspreload.org/)
[helps to find problems with https](https://www.badssl.com/)
[http vs https](https://www.httpvshttps.com/)
[test ssl](https://www.ssllabs.com/ssltest/)

The main point of HTTPS is
authentication of the visited website and protection of the privacy and integrity of the exchanged data.
HTTPS creates a secure channel over an insecure network.

Solves problems with:
* Man-In-The-Middle.
* Phishing
* DNS hijacking

<br>CA - Certification Authority.
<br>SSL - Secure Sockets Layer (SSL 3.0 is vulnerable).
<br>TLS - Transport Layer Security, is modern implementation of SSL.

HTTP Strict Transport Security:
`strict-transport-security` header in response with `max-age` as value
says to browser to reflect to 307 redirect and faster perform secure (not insecure) request.

<meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests">
