JSON Web Token
-

[online](https://jwt.io/)

JWK (JSON Web Key) - key containing the public key that should be used to verify any JWT.

Compact (small size of token) self-contained (payload contains all the required information,
no need to query the database).

Useful for:
* Authentication.
* Information Exchange
  (JWTs can be signed using public/private key pairs - ensures
  the senders are who they say they are).

Tokens consist of three parts separated by dots (.), which are:
* Header (hashing algorithm being used: HMAC, SHA256 or RSA).
* Payload (claims).
* Signature (used to verify that the sender of the JWT is who it says it is).

````
exp - expiration, unix time stamp
iat - issued at
````

HMAC (symmetric) algorithm - probably the most common algorithm for signed JWTs.
HS256 - HMAC with SHA-256.
RSA and ECDSA algorithms (asymmetric).

Never include secret info into JWT.

Token could be stolen, so it must have expiration.

AccessToken - with short TTL.
RefreshToken - with long TTL, used only to get new AccessToken, will be expired after 1 usage.
