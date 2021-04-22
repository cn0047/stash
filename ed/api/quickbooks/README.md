QuickBooks
-

[dashboard](https://developer.intuit.com/app/developer/dashboard)
[app keys](https://developer.intuit.com/app/developer/appdetail/test/keys)

````sh
access_token  expires_in                 3600    = 60min
refresh_token x_refresh_token_expires_in 8726109 = few years
````

````sh
# consumer_key/client_id; consumer_secret; realm_id
cId=''
scrt=''
rId=''

jha='Accept: application/json'



# get token, step 1
go run ed/l/go/examples/http/http.server.debug.go
cb='http://localhost:8080'
cb='http%3A%2F%2Flocalhost%3A8080'
s='7d04dcc2-11e8-44fc-b466-8bb0c1d5ff48'
open "https://appcenter.intuit.com/app/connect/oauth2/authorize?client_id=$cId&redirect_uri=$cb&response_type=code&scope=com.intuit.quickbooks.accounting&state=$s&realm_id=$rId"

# get token, step 2
code=''
rId=''
t=`echo -n "$cId:$scrt" | base64`
a="Authorization: Basic $t"
curl -s -X POST -H $a -H $jha 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
  -d grant_type=authorization_code \
  -d code=$code \
  -d redirect_uri=$cb \
  | jq

# refresh token
rt=''
curl -s -X POST -H $a -H $jha 'https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer' \
  -d grant_type=refresh_token \
  -d refresh_token=$rt \
  | jq
````

````sh
# h=https://sandbox-accounts.platform.intuit.com
h=https://quickbooks.api.intuit.com
h=https://sandbox-quickbooks.api.intuit.com

jh='Content-Type: application/json'

t=''
a="Authorization: Bearer $t"

# query
qr() {
  curl -X GET -H $a -H $jha -H $jh "$h/v3/company/$rId/query?query=$q"
}



# get company info
curl -X GET -H $a -H $jha -H $jh "$h/v3/company/$rId/companyinfo/$rId?minorversion=8" | jq



# get product
curl -X GET -H $a -H $jha -H $jh "$h/v3/company/$rId/item/1" | jq

# products query
q='select count(*) from Item'
q='select+count%28%2A%29+from+Item'
q='select * from Item'
q='select+%2A+from+Item'
q='select * from Item order by Id startposition 0 maxresults 10'
q='select%20%2A%20from%20Item%20order%20by%20Id%20startposition%200%20maxresults%2010'
q='select * from Item where Type='Inventory' order by Id startposition 0 maxresults 10'
q='select%20%2A%20from%20Item%20where%20Type%3D%27Inventory%27%20order%20by%20Id%20startposition%200%20maxresults%2010'
qr | jq



# get customer
curl -X GET -H $a -H $jha -H $jh "$h/v3/company/$rId/customer/1" | jq

# customers query
q='select count(*) from Customer'
q='select+count%28%2A%29+from+Customer'
q='select * from Customer'
q='select+%2A+from+Customer'
qr | jq
````
