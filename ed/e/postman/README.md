Postman
-

[docs](https://learning.postman.com/docs)

````sh
ls /Users/$USER/Library/Application\ Support/Postman
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

#### Scripts

Pre-response script:
````js
const cryptoJS = require('crypto-js');

function calculateChecksum(data) {
    const secret = pm.environment.get("emailSecret");
    return cryptoJS.HmacSHA512(JSON.stringify(data), secret).toString(cryptoJS.enc.Hex);
}

const body = JSON.parse(pm.request.body.raw);
const checksum = calculateChecksum(body);
pm.collectionVariables.set("checksum", checksum);
````

Post-response script:
````js
const res = JSON.parse(pm.response.text());
pm.environment.set("ID", res._id);
console.log('Updated ENV var ID to:', res._id);
````
