locals {
  env_name = lower(var.env_name)
}

data "aws_vpc" "main" {
  default = true
}

data "aws_subnet_ids" "all" {
  vpc_id = data.aws_vpc.main.id
}

# template
data "template_file" "init" {
  template = "${file("${path.module}/init.tpl")}"
  vars = {
    dns = "ec2_x1_unknown"
  }
}

resource "aws_instance" "ec2_x1" {
  count         = 1
  ami           = "ami-00aa4671cbf840d82" // default Amazon Linux 2 AMI
  instance_type = "t2.micro"
  user_data     = "${data.template_file.init.rendered}"
  key_name      = aws_key_pair.mysshkey.key_name

  subnet_id = tolist(data.aws_subnet_ids.all.ids)[0]
  vpc_security_group_ids = [
    aws_security_group.allow_tls.id,
    aws_security_group.allow_http.id,
    aws_security_group.myssh.id,
  ]

  tags = {
    Name = "ec2_x1"
  }
}

output "aws_instance_public_dns" {
  value = aws_instance.ec2_x1[0].public_dns
}
