output "lambda_bucket_name" {
  value = aws_s3_bucket.lambda_bucket[0].bucket
}
