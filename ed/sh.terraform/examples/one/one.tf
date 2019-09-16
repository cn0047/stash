/*
This is simple example.
*/

variable "base_cidr_block" {
  description = "A /16 CIDR range definition, such as 10.1.0.0/16, that the VPC will use"
  default     = "10.1.0.0/16"
}

provider "aws" {
  // in place:
  access_key = "key"
  secret_key = "scr"
  region     = "eu-central-1"
  // or
  shared_credentials_file = "/Users/tf_user/.aws/creds"
  // or
  // export AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY=""
  // export AWS_SECRET_ACCESS_KEY="" or export AWS_SECRET_KEY=""
  // export AWS_REGION="us-west-2" or export AWS_DEFAULT_REGION=""
}

# module "servers" {
#   source = "./app-cluster"

#   servers = 5
# }

resource "aws_instance" "ktest" {
  count             = 3
  ami               = "Required"
  availability_zone = "Optional"
  placement_group   = "Optional"
  host_id           = "Required"
  tags = [
    "${var.my_tag}"
  ]
  instance_type      = ""
  vpc_security_group = [aws_security_group.id]
  user_data_2        = file("user_data.sh")
  user_data          = <<EOF
#!/bin/bash
ech 200
EOF
  user_data_3        = <<-EOF
    #!/bin/bash
    ech 200
  EOF
}

resource "aws_security_group" "gtest" {
  name                   = "Optional"
  name_prefix            = "Optional"
  description            = "Optional"
  egress                 = "Optional"
  revoke_rules_on_delete = "Optional"
  vpc_id                 = "Optional"
  tags                   = "Optional"
  ingress {
    from_port    = 80
    to_port      = 80
    protoclol    = "tcp"
    cidr_block   = ["0.0.0.0/0"]
    cidr_block_2 = var.base_cidr_block
  }
  egress {
    from_port  = 0
    to_port    = 0
    protoclol  = "-1"
    cidr_block = ["0.0.0.0/0"]
  }
}
