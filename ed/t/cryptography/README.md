Cryptography
-

Symmetric encryption (symmetric key cryptography) - sender and receiver have same secret key.
Same key for encrypting and decrypting.

Public key cryptography - sender have public key, receiver have private key.
It is expensive type of cryptography.

Hashing algorithms:

* ~~3DES~~ (168 bits)
* AES (128 or 256 bits) - Advanced Encryption Standard
* Chacha20
* DH (Diffie-Hellman)
* ECDH (Elliptical Curve Diffie-Hellman)
* ECDHE
* ECDSA (asymmetric)
* HMAC (symmetric)
* ~~MD5~~
* RSA (asymmetric)
* ECC
* ~~RC4~~
* ~~SHA1~~
* **SHA2**
* SHA256
* **SHA3**
* SHA384
* Blowfish
* Twofish

* Argon2 (winner of the password hashing competition)
* bcrypt
* scrypt
* PBKDF2

Don't use (it's vulnerable) : MD5, SHA1, 3DES, RC4.
<br>Use SHA2, SHA3 (most secure).
Not secure: CRC32.
