provider "aws" {
  shared_credentials_file = "/Users/k/.aws/credentials"
  region                  = "eu-central-1"
  profile                 = "x"
}

locals {
  env_name = lower(terraform.workspace)
}

variable "instance_size" {
  type = map(string)
}

resource "aws_instance" "myec1st2" {
  count         = 1
  ami           = "ami-00aa4671cbf840d82" // default Amazon Linux 2 AMI
  instance_type = var.instance_size[terraform.workspace]
  tags = {
    Name = "myec1st2"
  }
  # user_data = file("user_data.sh")
  user_data = <<EOF
    #!/bin/bash
    curl -XPOST 'https://realtimelog.herokuapp.com:443/dyhnt08p53n' \
    -H 'Content-Type: application/json' -d '{"msg": "tf"}'
  EOF
}
