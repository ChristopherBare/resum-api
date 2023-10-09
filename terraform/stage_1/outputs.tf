output "bucket_name" {
  value = aws_s3_bucket.lambda_bucket.bucket
}

output "bucket_exists" {
  value = aws_s3_bucket.lambda_bucket ? true : false
}
