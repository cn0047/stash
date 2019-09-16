Terraform
-

[docs](https://www.terraform.io/docs/index.html)
[registry](https://registry.terraform.io/)
[aws](https://www.terraform.io/docs/providers/aws/index.html)

````sh
terraform init
terraform plan    # shows plan which will by applied
terraform apply   # apply plan
terraform destroy # destroy all from tf file

terraform plan -var "my_tag=${MY_TAG_FROM_ENV}"
````

Module - container for multiple resources that are used together.
Data sources - allows configuration to use information defined outside of tf.

````
# list
aws_instance.example[*].id
# one
aws_instance.example[0].id

[for k, device in aws_instance.example.device : k => device.size]

var.a != "" ? var.a : "default-a"
"Hello, %{ if var.name != "" }${var.name}%{ else }unnamed%{ endif }!"
````
