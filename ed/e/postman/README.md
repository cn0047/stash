Postman
-

[docs](https://learning.postman.com/docs)

````sh
# share
cp -r /Users/user/Library/Application\ Support/Postman ~/Downloads
mv /Users/user/Library/Application\ Support/Postman /Users/user/Library/Application\ Support/_Postman
mv ~/Downloads/Postman /Users/user/Library/Application\ Support/
rm -rf /Users/user/Library/Application\ Support/_Postman
````

Post-response script:
````js
const res = JSON.parse(pm.response.text());
pm.environment.set("ID", res._id);
console.log('Updated ENV var ID to:', res._id);
````

#### Authorization:

OAuth1:

````sh
Signature:      HMAC-SHA1

# ENV vars:
ConsumerKey:    {{consumer_key}}
ConsumerSecret: {{consumer_secret}}
Token:          {{token}}
TokenSecret:    {{token_secret}}
````
