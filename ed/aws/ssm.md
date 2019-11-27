SSM (AWS Systems Manager)
-

Parameter Store - to store secrets & credentials.

````sh
aws ssm get-parameters-by-path --path=/prj/qa/ | jq
aws ssm get-parameters-by-path --path=/prj/qa/ | jq jq '.Parameters[]|.Name'
aws ssm get-parameter --name=/prj/qa/PRM | jq
````
