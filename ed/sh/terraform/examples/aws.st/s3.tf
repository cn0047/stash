resource "aws_s3_bucket" "st_s3" {
  bucket = "basicbkt"
  acl    = "private"
  force_destroy = true
  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
