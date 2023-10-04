variable "bucket_short_name" {
  description = "bucket name"
}
variable "github_repo" {
  description = "repo name"
}
variable "github_repo_user" {
  description = "repo user name"
}
variable "github_repo_full" {
  description = "full repo name"
}

variable "branch_name" {
  default = "master"
  description = "default branch name"
}