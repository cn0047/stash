JSON Web Token
-

Compact (small size of token) self-contained (payload contains all the required information,
no need to query the database)

Useful for:

* Authentication.
* Information Exchange
  (JWTs can be signed using public/private key pairs - ensures
  the senders are who they say they are).

Tokens consist of three parts separated by dots (.), which are:

* Header (hashing algorithm being used: HMAC, SHA256 or RSA).
* Payload.
* Signature (used to verify that the sender of the JWT is who it says it is).
