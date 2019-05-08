Google Merchants
-

[merchants](https://merchants.google.com)

````sh
apiKey=${key}
merchantId=${id}

curl "https://www.googleapis.com/content/v2/accounts/authinfo?key=${apiKey}"

curl "https://www.googleapis.com/content/v2/$merchantId/products"
````
