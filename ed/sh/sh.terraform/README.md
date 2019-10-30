Terraform
-

[docs](https://www.terraform.io/docs/index.html)
[registry](https://registry.terraform.io/)
[aws](https://www.terraform.io/docs/providers/aws/index.html)

Module - container for multiple resources that are used together.
Data sources - allows configuration to use information defined outside of tf.

````sh
# list
aws_instance.example[*].id
# one
aws_instance.example[0].id

[for k, device in aws_instance.example.device : k => device.size]

var.a != "" ? var.a : "default-a"
"Hello, %{ if var.name != "" }${var.name}%{ else }unnamed%{ endif }!"
````

````sh
terraform init

# shows plan which will by applied
terraform plan
terraform plan -target=resource -out=plan
terraform plan -var "my_tag=${MY_TAG_FROM_ENV}"

# apply plan
terraform apply
terraform apply plan
terraform apply -auto-approve

# destroy all from tf file
terraform destroy
terraform destroy -target=resource

terraform graph
````

````sh
# aws lambda
export GOPATH=$PWD/ed/go/examples/aws
go get -u "github.com/aws/aws-sdk-go/aws"
GOOS=linux go build -o /tmp/awsLambdaOne $GOPATH/src/app/lambda
cd /tmp && zip awsLambdaOne.zip awsLambdaOne && cd -

cd ed/sh/sh.terraform/examples/aws.st
terraform init
terraform plan
terraform apply -auto-approve
terraform refresh
````
