TLS (Transport Layer Security)
-

Use TLSv1.1 or TLSv1.2 (TLSv1.0 is vulnerable).

TLS termination - use https for public accessible load-balancer,
then decrypting the TLS and passing unencrypted request downstream.

TLS works on top of TCP.

Certificate types: self-signed and wildcard (`*.kint.com`)

#### Data Encryption Protocols Cipher:

* ~~3DES~~ (168 bits) (don't use is vulnerable)
* AES (128 or 256 bits) (*USA government standard encryption protocol now*)
  * GCM
  * CBC
* Chacha20
  * Poly1305

#### Key Encryption Protocols:

* RSA (don't use is vulnerable)
* DH (Diffie-Hellman)
  * Signed with RSA
* ECDH (Elliptical Curve Diffie-Hellman)
  * Signed with ECDSA
  * Signed with RSA

#### Handshake Intergrity

* SHA
* SHA-256
* SHA-384

#### TLS handshake:

1. After TCP handshake, Client sends hello to server.
2. Server responds with hello and certificate
   which contains public key (`ServerEncryptedKey`) back to client.
3. Client verifies certificate against CAs.
4. Client generates symmetric pre-master key and encrypt it with public key from certificate
   and transmits it to the server.
5. Server decrypt pre-master and generates symmetric key.
   Now both client and server have symmetric key (session key).
5. TLS is established and secure communication begins.

Step 3:
Browsers have public keys for all of the major Certificate Authorities,
and it uses this public key to verify that the web server's certificate
was indeed signed by the trusted Certificate Authority.
