resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "lambda-bucket-${random_uuid.uuid.result}"
}

resource "random_uuid" "uuid" {}
