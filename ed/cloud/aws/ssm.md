SSM (AWS Systems Manager)
-

Parameter Store - to store secrets & credentials.

````sh
aws ssm put-parameter --overwrite --name='/prj/qa/prm' --type='String' --value='none'

p='/prj/qa/'
aws ssm get-parameters-by-path --path=$p | jq
aws ssm get-parameters-by-path --path=$p | jq jq '.Parameters[]|.Name'
aws ssm get-parameters-by-path --recursive --path=$p | jq

aws ssm get-parameter --name=/prj/qa/prm | jq
````
