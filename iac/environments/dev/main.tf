terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {}
}

provider "aws" {
  region = "ap-northeast-2"
}

module "container_registry" {
  source = "../../modules/container-registry"

  ecr_repository_name = var.ecr_repository_name
  ecr_image_tag_mutability = var.ecr_image_tag_mutability
  ecr_scan_on_push = var.ecr_scan_on_push
  ecr_encryption_type = var.ecr_encryption_type
}