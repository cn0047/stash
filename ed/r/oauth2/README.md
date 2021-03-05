OAuth
-
<br>OAuth1
<br>OAuth2

OpenID needed to understand who user is.

OAuth2 it's about delegated authorization, for example:
facebook delegates permissions via token to perform certain actions on user's behalf.
It's simplified version of OAuth.

OAuth2 - is not protocol, because big companies (google, facebook, etc) made own changes that's why it's framework.

`Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9` works only with SSL.

````
ConsumerKey.
ConsumerSecret.
Realm - realm of authorization.
Signer - OAuth1 Signer (HMAC-SHA1, HMAC-SHA256).
Noncer - creates request nonces.
````

Example (Facebook):

1st time:
  * Login page "Log in with Facebook" and href leads to fb-callback.php.
  * SDK.getAccessTokenFromFacebook() opens page "Facebook agreement".
  * Facebook sends callback to fb-callback.php with token.
  * $tokenMetadata->validateAppId(APP_ID);
  * $tokenMetadata->validateExpiration();
  * Deal with Facebook user.
  * $_SESSION['fb_access_token'] = (string) $accessToken.
  * ...

2nd time:
  * On any page we have $_SESSION['fb_access_token']
  * $tokenMetadata->validateAppId(APP_ID);
  * $tokenMetadata->validateExpiration();
  * ...

If token expired - go to spet "1st time" (show login page).
