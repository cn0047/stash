resource "aws_sqs_queue" "st_queue" {
  name                      = "simpleQueue"
  delay_seconds             = 1
  max_message_size          = 2048
  message_retention_seconds = 86400
  receive_wait_time_seconds = 1
  tags = {
    APP_ENVIRONMENT = "dev"
  }
}
