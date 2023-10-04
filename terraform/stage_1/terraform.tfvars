# terraform.tfvars
bucket_short_name = "lambda_bucket"
github_repo       = "resum-api"
github_repo_user  = "ChristopherBare"
repository_name   = {
  default = data.github_repository.repo.name
}
branch_name = {
  default = "master"
}
github_repo_full = {
  name = "${var.github_repo_user}/${var.github_repo}"
}