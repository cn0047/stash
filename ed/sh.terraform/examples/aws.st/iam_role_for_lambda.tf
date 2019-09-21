resource "aws_iam_role" "st_iam_role_for_lambda" {
  name               = "st-iam-role-for-lambda"
  assume_role_policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
  }
  EOF
}

resource "aws_iam_policy" "st_iam_policy_for_lambda" {
  name   = "st-iam-policy-for-lambda"
  policy = <<-EOF
  {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Sid": "VisualEditor0",
              "Effect": "Allow",
              "Action": [
                  "dynamodb:BatchGetItem",
                  "dynamodb:BatchWriteItem",
                  "dynamodb:UpdateTimeToLive",
                  "dynamodb:PutItem",
                  "dynamodb:DeleteItem",
                  "dynamodb:Scan",
                  "dynamodb:ListTagsOfResource",
                  "dynamodb:Query",
                  "dynamodb:DescribeStream",
                  "dynamodb:UpdateGlobalTable",
                  "dynamodb:UpdateItem",
                  "dynamodb:DescribeTimeToLive",
                  "dynamodb:DeleteTable",
                  "dynamodb:CreateTable",
                  "dynamodb:UpdateGlobalTableSettings",
                  "dynamodb:DescribeGlobalTableSettings",
                  "dynamodb:DescribeTable",
                  "dynamodb:DescribeGlobalTable",
                  "dynamodb:GetItem",
                  "dynamodb:CreateGlobalTable",
                  "dynamodb:UpdateTable",
                  "dynamodb:GetRecords",
                  "dynamodb:GetShardIterator"
              ],
              "Resource": "*"
          },
          {
              "Sid": "VisualEditor1",
              "Effect": "Allow",
              "Action": [
                  "dynamodb:ListGlobalTables",
                  "dynamodb:ListTables",
                  "dynamodb:ListStreams"
              ],
              "Resource": "*"
          }
      ]
  }
  EOF
}

resource "aws_iam_policy_attachment" "st_iam_policy_attachment" {
  name       = "st-iam-policy-attachment"
  roles      = ["${aws_iam_role.st_iam_role_for_lambda.name}"]
  policy_arn = "${aws_iam_policy.st_iam_policy_for_lambda.arn}"
}
