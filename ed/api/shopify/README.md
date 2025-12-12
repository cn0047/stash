Shopify
-

````sh
shop=mystore
t=token

a="X-Shopify-Access-Token: $t"
jh='Content-Type: application/json'

# shop info
curl -X GET -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-10/shop.json" | jq

# shop name
curl -X POST -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-10/graphql.json" \
  -d '{"query": "query { shop { name } }"}' | jq

````

````sh
shop=mystore
client=c
token_secret=t

curl -X POST "https://$shop.myshopify.com/admin/oauth/access_token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -d "client_id=$client" \
  -d "client_secret=$token_secret" | jq

````
