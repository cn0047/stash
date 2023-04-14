ECR (Elastic Container Registry)
-

````sh
aws ecr get-login --region us-east-1 --no-include-email
aws ecr create-repository --region us-east-1 --repository-name rName

aws ecr list-images --repository-name=x-eval

aws ecr describe-images --region us-east-1 --repository-name lf-dev
````
