resource "aws_dynamodb_table" "st_ddb" {
  name             = "hotdata"
  billing_mode     = "PAY_PER_REQUEST"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"
  ttl {
    attribute_name = ""
    enabled        = false
  }
  tags = {
    Name = "st-ddb"
  }
  hash_key = "key"
  attribute {
    name = "key"
    type = "S"
  }
}

resource "aws_dynamodb_table_item" "st_ddb_item1" {
  table_name = "${aws_dynamodb_table.st_ddb.name}"
  hash_key   = "${aws_dynamodb_table.st_ddb.hash_key}"
  item       = <<-ITEM
  {
    "key": {"S": "name"},
    "val": {"S": "st"}
  }
  ITEM
}
