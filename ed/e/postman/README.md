Postman
-

[docs](https://learning.postman.com/docs)

````sh
ls /Users/$USER/Library/Application\ Support/Postman
````

Post-response script:
````js
const res = JSON.parse(pm.response.text());
pm.environment.set("ID", res._id);
console.log('Updated ENV var ID to:', res._id);
````

#### Authorization

OAuth1:

````sh
Signature:      HMAC-SHA1

# ENV vars:
ConsumerKey:    {{consumer_key}}
ConsumerSecret: {{consumer_secret}}
Token:          {{token}}
TokenSecret:    {{token_secret}}
````
