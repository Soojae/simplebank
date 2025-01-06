terraform {
  required_version = "~> 1.5.0"

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

  ecr_repository_name      = var.ecr_repository_name
  ecr_image_tag_mutability = var.ecr_image_tag_mutability
  ecr_scan_on_push         = var.ecr_scan_on_push
  ecr_encryption_type      = var.ecr_encryption_type
}

module "db" {
  source = "../../modules/db"

  aws_region           = var.aws_region
  db_identifier        = var.db_identifier
  db_engine_version    = var.db_engine_version
  db_instance_class    = var.db_instance_class
  db_allocated_storage = var.db_allocated_storage
  db_storage_type      = var.db_storage_type
  db_username          = var.db_username
  db_subnet_group_name = var.db_subnet_group_name
  publicly_accessible  = var.publicly_accessible
  multi_az             = var.multi_az
  storage_encrypted    = var.storage_encrypted
  existing_sg_name     = var.existing_sg_name
  db_name              = var.db_name
}