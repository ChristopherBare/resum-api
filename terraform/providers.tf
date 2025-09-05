terraform {
  required_providers {

  }
}

provider "aws" {
  region = "us-east-1"
}

data "aws_region" "current" {}