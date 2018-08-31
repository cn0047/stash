Apple News
-

[newspublisher](https://www.icloud.com/#newspublisher)
[doc](https://developer.apple.com/documentation/apple_news/apple_news_api)

````sh
channelId=""
key=""
secret=""

# ⚠️ Not ready yet(

url="https://news-api.apple.com/channels/"$channelId
date=`date +%Y-%m-%dT%T%z`
canonicalURL="GET"$url$date
secretDecoded=`echo -n $secret | base64 -D`
signature=`echo -n $canonicalURL | openssl sha256 -hmac '"'$secretDecoded'"'`
signatureInBase64=`echo $signature | base64`

curl -v -X GET $url \
  -H 'Accept: application/json' \
  -H  'Authorization: HHMAC; key="'$key'"; signature="'$signatureInBase64'"; date="'$date'"'

````

#### Apple News Format

[components](https://developer.apple.com/documentation/apple_news/apple_news_format/components)
