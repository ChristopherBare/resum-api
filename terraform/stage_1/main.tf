data "github_repository" "repo" {
  full_name = var.github_repo_full.name
}

resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "${var.bucket_short_name}-${var.github_repo}-${var.branch_name}"
}


