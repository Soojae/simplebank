provider "aws" {
  region = var.aws_region
}

# 기존 Subnet Group 가져오기
data "aws_db_subnet_group" "existing_subnet_group" {
  name = var.db_subnet_group_name
}

# 기존 Security Group 가져오기
data "aws_security_group" "existing_sg" {
  name = var.existing_sg_name
}

# 랜덤 비밀번호 생성
resource "random_password" "db_password" {
  length           = 16
  special          = true
  override_special = "!@#$%^&*()-_+=<>?"
}

# RDS 인스턴스 생성
resource "aws_db_instance" "simple_bank" {
  identifier           = var.db_identifier
  engine               = "postgres"
  engine_version       = var.db_engine_version
  instance_class       = var.db_instance_class
  allocated_storage    = var.db_allocated_storage
  storage_type         = var.db_storage_type
  username             = var.db_username
  password             = random_password.db_password.result
  db_subnet_group_name = data.aws_db_subnet_group.existing_subnet_group.name
  publicly_accessible  = var.publicly_accessible
  multi_az             = var.multi_az
  skip_final_snapshot  = true
  storage_encrypted    = var.storage_encrypted
  db_name = var.db_name

  vpc_security_group_ids = [data.aws_security_group.existing_sg.id]

  tags = {
    Name = var.db_identifier
  }
}

