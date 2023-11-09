provider "aws" {
  region = "us-east-1"
}
variable "GITHUB_TOKEN" {
  type = string
}
provider "github" {
  token = var.GITHUB_TOKEN
}
