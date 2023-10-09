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

resource "aws_s3_bucket" "lambda_bucket" {
  count = data.terraform_remote_state.stage_1.outputs.bucket_exists
  bucket = "${var.bucket_short_name}-${var.github_repo}-${var.branch_name}"
}



