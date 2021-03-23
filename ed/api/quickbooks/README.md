QuickBooks
-

[dashboard](https://developer.intuit.com/app/developer/dashboard)
[app keys](https://developer.intuit.com/app/developer/appdetail/test/keys)
[]()

````sh
cId=''
scrt=''
rId=''



# Get token, step 1
cb='http://localhost:8080'
cb='http%3A%2F%2Flocalhost%3A8080'
s='7d04dcc2-11e8-44fc-b466-8bb0c1d5ff48'
open "https://appcenter.intuit.com/app/connect/oauth2/authorize?client_id=$cId&redirect_uri=$cb&response_type=code&scope=com.intuit.quickbooks.accounting&state=$s&realm_id=$rId"

# Get token, step 2
code=''
t=`echo -n "$cId:$scrt" | base64`
a="Authorization: Basic $t"
jha='Accept: application/json'
curl -s -X POST -H $a -H $jha 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
  -d grant_type=authorization_code \
  -d code=$code \
  -d redirect_uri=$cb \
  | jq
````
