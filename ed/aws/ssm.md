SSM (AWS Systems Manager)
-

Parameter Store - to store secrets & credentials.

````sh
aws ssm get-parameters-by-path --path=/prj/qa/ | jq
````
