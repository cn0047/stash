SSM (AWS Systems Manager)
-

Parameter Store - to store secrets & credentials.

````sh
aws ssm put-parameter --overwrite --name='/prj/qa/prm' --type='String' --value='none'

aws ssm get-parameters-by-path --path=/prj/qa/ | jq
aws ssm get-parameters-by-path --path=/prj/qa/ | jq jq '.Parameters[]|.Name'

aws ssm get-parameter --name=/prj/qa/prm | jq
````
