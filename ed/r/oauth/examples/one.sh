# OAuth

k=""   # consumerKey
sk=""  # consumerSecretKey
at=""  # token
ast="" # tokenSecret

# ok (on ubuntu)
m='GET'
u='https://dev.org.com/api/v2.0/my-groups'
#
t=$(date +%s)
nonce=$(openssl rand -base64 32 | tr -dc 'a-zA-Z0-9' | cut -c1-16)
#
s=$m$u'?oauth_consumer_key='$k'&oauth_token='$at'&oauth_timestamp='$t'&oauth_nonce='$nonce'&oauth_version=1.0&oauth_signature_method=HMAC-SHA1'
key="$sk&$ast"
signature=$(echo -n "$s" | openssl dgst -sha1 -hmac "$key" | sed 's/^.* //')
a='OAuth oauth_consumer_key="'$k'",oauth_token="'$at'",oauth_timestamp="'$t'",oauth_nonce="'$nonce'",oauth_version="1.0",oauth_signature_method="HMAC-SHA1",oauth_signature="'$signature'"'
#
curl -H "Authorization: $a" $u
