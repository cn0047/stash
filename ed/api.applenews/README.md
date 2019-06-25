Apple News
-

[newspublisher](https://www.icloud.com/#newspublisher)
[doc](https://developer.apple.com/documentation/apple_news/apple_news_api)
[components](https://developer.apple.com/documentation/apple_news/apple_news_format/components)
[linking-to-channel](https://www.apple.com/itunes/marketing-on-news/identity-guidelines.html#linking-to-your-channel)

````sh
key=""
secret=""
channelId=""

# 1️⃣ common
contentType='application/json'
date=`date +%Y-%m-%dT%T%z` # MUST be value like: 2019-03-26T13:41:58+02:00

# get channel
url="https://news-api.apple.com/channels/$channelId"
canonicalURL="GET$url$date"
# run curl

# get sections
url="https://news-api.apple.com/channels/$channelId/sections"
canonicalURL="GET$url$date"
# run curl

# get articles
url="https://news-api.apple.com/channels/$channelId/articles"
canonicalURL="GET$url$date"
# run curl

# get article
appleArticleID=""
url="https://news-api.apple.com/articles/$appleArticleID"
canonicalURL="GET$url$date"
# run curl

# promote articles in section
sectionId=""
appleArticleID=""
url="https://news-api.apple.com/sections/$sectionId/promotedArticles"
body='{"data":{"promotedArticles":["https://news-api.apple.com/articles/'$appleArticleID'"]}}'
canonicalURL="POST$url$date$contentType$body"
# run POST curl

# 2️⃣ before run curl
secretRaw=`echo -n $secret | base64 -D`
signatureBin=`echo -n "$canonicalURL" | openssl dgst -sha256 -hmac "${secretRaw}" -binary`
signature=`echo -n $signatureBin | base64`

# run GET curl
curl -X GET $url \
  -H 'Accept: '$contentType \
  -H  'Authorization: HHMAC; key="'$key'"; signature="'$signature'"; date="'$date'"' \
  | jq

# run POST curl
curl -X POST $url \
  -H 'Accept: '$contentType \
  -H 'Content-Type: '$contentType \
  -H  'Authorization: HHMAC; key="'$key'"; signature="'$signature'"; date="'$date'"' \
  -d $body \
  | jq

````

#### Apple News Format

[components](https://developer.apple.com/documentation/apple_news/apple_news_format/components)

https://apple.news/AqoCs5rTJQq-EI8o1mMxiwg

````
"minimumHeight": "50vh" # Viewport Height
````
