OAuth
-
<br>OAuth1
<br>OAuth2

[rfc](https://tools.ietf.org/html/rfc6749)

OpenID needed to understand who user is.

OIDC - OpenID Connect.

OAuth2 it's about delegated authorization, for example:
facebook delegates permissions via token to perform certain actions on user's behalf.
It's simplified version of OAuth.

OAuth2 - is not protocol, because big companies (google, facebook, etc) made own changes
that's why it's framework/set of rules.

Benefit: user won't share credentials, but token.
Benefit: revoke token any time.

````sh
# oauth1
Authorization: OAuth oauth_consumer_key='',oauth_timestamp='1488918667',oauth_nonce='Mevx8gQQ3pzmkdXq',oauth_version='1.0',oauth_signature_method='HMAC-SHA1',oauth_token='',oauth_signature='95556c9e7...',oauth_callback='nop',application_name='app'

# oauth2
# works only with SSL.
Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9
````
````
ConsumerKey.
ConsumerSecret.
Realm - realm of authorization.
Signer - OAuth1 Signer (HMAC-SHA1, HMAC-SHA256).
Noncer - creates request nonces.
````

Example (Facebook):
1st time:
  * Login page "Login with Facebook" and href leads to `fb-callback.php`.
  * `SDK.getAccessTokenFromFacebook()` opens page "Facebook agreement".
  * Facebook sends callback to `fb-callback.php` with token.
  * `$tokenMetadata->validateAppId(APP_ID);`.
  * `$tokenMetadata->validateExpiration();`.
  * Deal with Facebook user.
  * `$_SESSION['fb_access_token'] = (string)$accessToken;`.
  * ...
2nd time:
  * On any page we have `$_SESSION['fb_access_token']`.
  * `$tokenMetadata->validateAppId(APP_ID);`.
  * `$tokenMetadata->validateExpiration();`.
  * ...
If token expired - go to step "1st time" (show login page).

Bash example:
````sh
h=https://oauth.prj.com
cId=stark # client_id
scp=read # scope

# get token
d='{
  "username": "'$usr'",
  "password": "'$pwd'",
  "client_id": "'$cId'",
  "grant_type": "password",
  "scope":     "client"
}'
t=`curl -X POST $h/oauth/token -H 'content-type: application/json' -d $d | jq -r '.access_token'`
echo $t

# login flow (step 1), get code
cb='http://localhost:14000/oauth/code'
open "$h/oauth/authorize?response_type=code&client_id=$cId&redirect_uri=$cb&scope=$scp&secret=1"
# login flow (step 2), get token
c=''
curl -v -X POST $h/oauth/token -H 'content-type: application/x-www-form-urlencoded' \
  -d "client_id=$cId&grant_type=authorization_code&code=$c&redirect_uri=$cb&scope=$scp" | jq
# login flow (step 2), get token
c=''
d='{
  "client_id": "'$cId'",
  "code": "'$c'",
  "grant_type": "authorization_code",
  "redirect_uri": "'$cb'",
  "scope": "'$scp'"
}'
curl -X POST $h/oauth/token -H 'content-type: application/json' -d $d | jq

# refresh token
rt=''
d='{
  "client_id": "'$cId'",
  "client_secret": "client-secret",
  "grant_type": "refresh_token",
  "refresh_token": "'$rt'"
}'
t=`curl -X POST $h/oauth/token -H 'content-type: application/json' -d $d | jq -r '.access_token'`
echo $t
````
