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

resource "null_resource" "set_create_bucket" {
  triggers = {
    create_bucket = var.create_bucket
  }

  provisioner "local-exec" {
    command = "echo 'create_bucket = ${var.create_bucket}' > create_bucket.auto.tfvars"
  }
}
resource "aws_s3_bucket" "lambda_bucket" {
  count = data.terraform_remote_state.stage_1.outputs.bucket_exists
  bucket = "${var.bucket_short_name}-${var.github_repo}-${var.branch_name}"
}

# Set the `create_bucket` variable based on the existence of the S3 bucket
variable "create_bucket" {
  description = "Create the S3 bucket if it doesn't exist"
  type        = bool
  default     = data.terraform_remote_state.stage_1.outputs.bucket_exists ? false : true
}



