Terraform
-

[docs](https://www.terraform.io/docs/index.html)
[registry](https://registry.terraform.io/)
[aws](https://www.terraform.io/docs/providers/aws/index.html)

Module - container for multiple resources that are used together.
Data sources - allows configuration to use information defined outside of tf.

````sh
src
└─ terraform
   ├─ environments
   └─ modules
````

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

#### Examples:

````sh
# aws.st
# aws lambda
export GOPATH=$PWD/ed/l/go/examples/aws
go get -u "github.com/aws/aws-sdk-go/aws"
# build
GOOS=linux go build -o /tmp/awsLambdaOne $GOPATH/src/app/lambda
cd /tmp && zip awsLambdaOne.zip awsLambdaOne && cd -
# test
open https://realtimelog.herokuapp.com:443/64kfym341kp2
go run $GOPATH/src/app/main.go k31 val200 200
for i in $(seq 2000 2999); do go run $GOPATH/src/app/main.go "k$i" val200 $i; done

cd ed/sh/terraform/examples/aws.st
terraform init
terraform plan
terraform apply -auto-approve
terraform refresh
terraform show
````
