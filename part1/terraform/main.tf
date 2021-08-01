resource "aws_s3_bucket" "helm_repository" {
  bucket        = var.bucket_name
  acl           = var.acl
  force_destroy = var.force_destroy

  versioning {
    enabled = var.versioning
  }

  tags = {
    Name        = var.bucket_name
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_object" "home_assessment" {
  bucket = aws_s3_bucket.helm_repository.id
  key    = var.key
}