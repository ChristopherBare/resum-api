provider "aws" {
  region = "us-east-1"
}
provider "github" {
  token = env("GITHUB_TOKEN")
}
