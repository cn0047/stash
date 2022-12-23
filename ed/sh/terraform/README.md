Terraform
-

<br>v0.12.12

[docs](https://www.terraform.io/docs/index.html)
[registry](https://registry.terraform.io/)
[aws](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)
[gcp](https://registry.terraform.io/providers/hashicorp/google/latest/docs)
[functions](https://www.terraform.io/docs/configuration/functions/strrev.html)

Terraform - tool for defining infrastructure as code.

<br>HCL - configuration files.
<br>Syntax - blocks and nested blocks.
<br>Module - container for multiple resources that are used together.
<br>Workspace - holds backend that defines how operations executed and where persistent state and else.
<br>Sentinel - embedded policy-as-code framework integrated with the HashiCorp Enterprise products.
<br>Providers - aws, gcp, etc.
<br>Resources - ec2, s3, rds, etc.
<br>Input variables - serve as parameters for a tf module (`variable "x" {type = string}`).
<br>Output values - return values of a Terraform module (`val = aws_instance.server.private_ip`).
<br>Expressions - refer to or compute values within a configuration.
<br>Backends - determines how state is loaded/stored (s3, artifactory, consul, etc.).
<br>Data sources - allows configuration to use information defined outside of tf
(`data "aws_ami" "example" { ...`) (it could be list of aws availability zones, etc.).
<br>Provisioners - used to model specific actions on machine in order
to prepare servers or other infrastructure objects (file, local-exec, remote-exec).

````sh
src
└─ terraform
   ├─ environments
   │  ├ dev
   │  └ qa
   └─ modules

.auto.tfvars
.auto.tfvars.json
terraform.tfvars
terraform.tfvars.json
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

flatten([["a", "b"], [], ["c"]])      # result: ["a", "b", "c"]
merge({a="b", c="d"}, {e="f", c="z"}) # result: {"a" = "b", "c" = "z", "e" = "f"}
````

````sh
brew install tflint

terraform console

terraform init
terraform init -reconfigure
terraform init -backend-config="bucket=$bkt"

terraform validate

# shows plan which will by applied
terraform plan
terraform plan -target=resource -out=plan
terraform plan -var "my_tag=${MY_TAG_FROM_ENV}"
terraform plan -state="./dev.state" -var-file="common.tfvars" -var-file="dev.tfvars"

# apply plan
terraform apply
terraform apply plan
terraform apply -auto-approve

terraform refresh # update the state

terraform state list # list all resources in the state
terraform state show # show details about object
terraform state rm $resource # remove object from state
terraform state pull
terraform state push

terraform show -json

# import already existing resource into tf state
terraform import -var-file=tf.vars $addr $id
terraform import -var-file=tf.vars module.vpc.aws_subnet.public[2] subnet-x9cb23
terraform import -var-file=tf.vars module.my-dev-gcp-sa-module.module.gh_oidc.google_iam_workload_identity_pool_provider.main projects/test-prj/locations/global/workloadIdentityPools/my-ga-pool/providers/my-ga-pool-provider

terraform destroy # destroy all from tf file
terraform destroy -target=resource # only certain resource, like: google_cloud_run_service.default
terraform destroy -auto-approve -input=false -target=$resource

terraform taint # mark resource for recreation

terraform graph
terraform fmt # format tf files
terraform fmt -recursive
terraform output

# workspace
terraform workspace list
terraform workspace select $ws
terraform workspace new dev
terraform workspace show
````

````sh
# simple interpolation
"${var.prefix}-app"

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

# module
source = "git@github.com:org/tf-module.git?ref=0.7.0"
````

#### GCP

````sh
google_compute_network                       # VPC network
google_compute_region_network_endpoint_group # regional NEG (bind network and cloudrun)
google_compute_router                        # router within GCE
google_compute_router_nat                    #
google_compute_subnetwork                    #
google_compute_url_map                       # route requests to BE svc
google_compute_target_https_proxy            # for global forwarding rule
                                             # to route incoming HTTPS requests to URL map
google_compute_global_forwarding_rule        # to forward traffic to correct HTTP LB
````

#### AWS

````sh
resource "aws_instance" "myec2" {
  count         = 1
  ami           = "ami-00aa4671cbf840d82" # default Amazon Linux 2 AMI
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
