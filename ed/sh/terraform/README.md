Terraform
-

<br>v0.12.12

[docs](https://www.terraform.io/docs/index.html)
[registry](https://registry.terraform.io/)
[aws](https://www.terraform.io/docs/providers/aws/index.html)
[functions](https://www.terraform.io/docs/configuration/functions/strrev.html)

<br>Module - container for multiple resources that are used together.
<br>Sentinel - embedded policy-as-code framework integrated with the HashiCorp Enterprise products.
<br>Providers - aws, gcp, etc.
<br>Data sources - allows configuration to use information defined outside of tf (`data "aws_ami" "example" { ...`).
<br>Input variables - serve as parameters for a tf module (`variable "x" {type = string}`).
<br>Output values - return values of a Terraform module (`val = aws_instance.server.private_ip`).
<br>Expressions - refer to or compute values within a configuration.
<br>Backends - determines how state is loaded (s3, artifactory, consul, etc.).
<br>Provisioners.

````sh
src
└─ terraform
   ├─ environments
   │  ├ dev
   │  └ qa
   └─ modules
````

````
+ created
~ updated
- deleted
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
terraform plan -state="./dev.state" -var-file="common.tfvars" -var-file="dev.tfvars"

# apply plan
terraform apply
terraform apply plan
terraform apply -auto-approve

terraform refresh

terraform show -json

# destroy all from tf file
terraform destroy
terraform destroy -target=resource

terraform graph
terraform fmt # format tf files
terraform output

# workspace
terraform workspace list
terraform workspace select $ws
terraform workspace new dev
terraform workspace show
````

#### Examples:

````sh
# aws.st
cd ed/sh/terraform/examples/aws.st
# aws lambda
export GOPATH=$PWD/ed/l/go/examples/aws
go get -u "github.com/aws/aws-sdk-go/aws"
# build
GOOS=linux go build -o /tmp/awsLambdaOne $GOPATH/src/app/lambda
cd /tmp && zip awsLambdaOne.zip awsLambdaOne && mv awsLambdaOne.zip /Users/kovpakvolodymyr/Downloads && cd -
# test
open https://realtimelog.herokuapp.com:443/64kfym341kp2
go run $GOPATH/src/app/main.go k31 val200 200
for i in $(seq 2000 2999); do go run $GOPATH/src/app/main.go "k$i" val200 $i; done

# aws.ec2
cd ed/sh/terraform/examples/aws.ec2
c=plan
c=apply
c=refresh
c=destroy
terraform $c -var-file=ec2.tfvars -lock=false

# aws.x
# @see: docket image build in k8s
cd ed/sh/terraform/examples/aws.x
export AWS_PROFILE=x
terraform init
terraform plan
# run
h=''
key=/Users/kovpakvolodymyr/web/kovpak/gh/ed/sh/ssh/examples/nopwd/id_rsa
scp -i $key /tmp/xgoapp ec2-user@$h:/tmp
# ssh and run bin file
ssh -i $key ec2-user@$h
````

````sh
resource "random_int" "rand" {
  min = 100
  max = 999
}

variable "images" {
  type = "map"
  default = {
    us-east-1 = "image-1234"
    us-west-2 = "image-4567"
  }
}
ami = lookup(var.amis, "us-east-1", "error")

# place to store state
terraform {
  backend "s3" {
    key = "k"
    region = "us-east-1"
  }
}

resource "aws_instance" "myec2" {
  count         = 1
  ami           = "ami-00aa4671cbf840d82" // default Amazon Linux 2 AMI
  instance_type = "t2.micro"
  tags = {
    Name = "myec2-1"
  }
  provisioner "remote-exec" {
    inline = [
      "curl -i -XPOST 'https://realtimelog.herokuapp.com:443/sdf83k' -H 'Content-Type: application/json' -d '{\"msg\": \"ec2\"}'"
    ]
  }
}
````
