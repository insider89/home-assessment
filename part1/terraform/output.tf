output "s3_bucket_name" {
  description = "s3 bucket name"
  value       = aws_s3_bucket.helm_repository.bucket_domain_name
}