variable "bucket_short_name" {}
variable "github_repo" {}
variable "github_repo_user" {}
variable "repository_name" {
  default = data.github_repository.repo.name
}

variable "branch_name" {
  default = "master"
}