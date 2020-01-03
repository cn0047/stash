provider "aws" {
  shared_credentials_file = "/Users/k/.aws/credentials"
  region                  = "eu-central-1"
}

resource "aws_instance" "myec2" {
  count         = 0
  ami           = "ami-00aa4671cbf840d82" // default Amazon Linux 2 AMI
  instance_type = "t2.micro"
  tags = {
    Name = "myec2-1"
  }
}
