data "terraform_remote_state" "stage_1" {
  backend = "s3"

  config = {
    bucket = "terraform-backend-state-bucket-lambda-resum-api"
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}

data "github_repository" "repo" {
  full_name = var.github_repo_full.name
}

locals {
  create_bucket = length(data.aws_s3_bucket.bucket_check) == 0 ? false : true
}

resource "aws_s3_bucket" "lambda_bucket" {
  count = local.create_bucket ? 1 : 0
  bucket = "${var.bucket_short_name}-${var.github_repo}-${var.branch_name}"
}

data "aws_s3_bucket" "bucket_check" {
  bucket = aws_s3_bucket.lambda_bucket.bucket
}

# Set the `create_bucket` variable based on the existence of the S3 bucket
variable "create_bucket" {
  description = "Create the S3 bucket if it doesn't exist"
  type        = bool
  default     = false
}



