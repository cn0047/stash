resource "aws_lambda_function" "st_ddb_lambda" {
  function_name    = "st-ddb-lambda"
  role             = "${aws_iam_role.st_iam_role_for_lambda.arn}"
  filename         = "/tmp/awsLambdaOne.zip"
  source_code_hash = "${filebase64sha256("/tmp/awsLambdaOne.zip")}"
  handler          = "awsLambdaOne"
  runtime          = "go1.x"
  environment {
    variables = {
      APP_ENV = "dev"
    }
  }
}

resource "aws_lambda_event_source_mapping" "st_ddb_lambda_trigger" {
  event_source_arn  = "${aws_dynamodb_table.st_ddb.stream_arn}"
  function_name     = "${aws_lambda_function.st_ddb_lambda.arn}"
  batch_size        = 2
  starting_position = "LATEST"
  enabled           = true
}
