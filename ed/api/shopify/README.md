Shopify
-

[dev dashboard](https://dev.shopify.com/dashboard)
[API access scopes](https://shopify.dev/docs/api/usage/access-scopes)
[API GraphQL](https://shopify.dev/docs/api/admin-graphql/2025-10)

````sh
shop=mystore
t=token

a="X-Shopify-Access-Token: $t"
jh='Content-Type: application/json'

# shop info
curl -X GET -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-10/shop.json" | jq

curl -X GET -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-01/products.json" | jq

# shop name
curl -X POST -H $a -H $jh "https://$shop.myshopify.com/admin/api/2026-01/graphql.json" \
  -d '{"query": "query { shop { name } }"}' | jq

curl -X POST -H $a -H $jh "https://$shop.myshopify.com/admin/api/2026-01/graphql.json" -d @- <<EOF
{
  "query": "{
    products(first: 5) {
      edges {
        node {
          id
          handle
        }
      }
      pageInfo {
        hasNextPage
      }
    }
  }"
}
EOF

# step 1
curl -X POST -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-10/graphql.json" \
  -d '{ "query": "mutation { bulkOperationRunQuery( query:\"\"\" { customers(query: \"NuORDER Sync*\") { edges{ node{ id createdAt tags email firstName lastName addresses { address_id: id company address1 address2 city province zip country } defaultAddress { id company address1 address2 city province zip country phone } } } } } \"\"\" ) { bulkOperation { id status } userErrors { field message } }}" }' | jq
# step 2
curl -X POST -H $a -H $jh "https://$shop.myshopify.com/admin/api/2025-10/graphql.json" \
  -d '{ "query": "{ node(id: \"gid://shopify/BulkOperation/5986806202466\") { ... on BulkOperation { id status errorCode createdAt completedAt objectCount fileSize url partialDataUrl } } }" }' | jq

````

````sh
shop=mystore
client_id=c
token_secret=t

# when shop and app in same org
curl -X POST "https://$shop.myshopify.com/admin/oauth/access_token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  -d "client_id=$client_id" \
  -d "client_secret=$token_secret"
# response examples
{"access_token":"shpua_tkn","scope":"","expires_in":86399}
{"errors":"[API] Invalid API key or access token (unrecognized login or wrong password)"}

# hen shop and app in various org
# 1
go run ed/l/go/examples/http/http.server.debug.go
cb=http://localhost:8080
nonce=ninfjwo04ojnh2l3
scopes=read_companies,write_companies,read_customers,write_customers,read_price_rules,read_discounts,write_draft_orders,read_draft_orders,read_inventory,read_locations,read_metaobject_definitions,read_metaobjects,write_order_edits,read_order_edits,read_orders,write_orders,read_payment_terms,read_product_listings,read_products
echo open "https://$shop.myshopify.com/admin/oauth/authorize?client_id=$client_id&scope=$scopes&state=$nonce&redirect_uri=$cb"
# 2
code=2e133cb6d3ba722d24f52c149391016f
curl -X POST "https://$shop.myshopify.com/admin/oauth/access_token" \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "client_id=$client_id" \
  -d "client_secret=$token_secret" \
  -d "code=$code"

````
