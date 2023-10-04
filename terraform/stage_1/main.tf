resource "aws_s3_bucket" "example_bucket" {
  bucket = var.bucket_short_name
}

# Step 2: Use a data source to fetch repository and branch information (GitHub in this example)
data "github_repository" "repo" {
  full_name = var.github_repo_full
}

# Step 4: Create an S3 bucket
resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "${var.bucket_short_name}-${var.repository_name}-${var.branch_name}"
}


