variable "aws_region" {
  type        = string
  description = "AWS region. String type."
  default     = "us-east-2"
}

variable "bucket_name" {
  type        = string
  description = "S3 bucket name. String type."
  default     = "home-helm-repository"
}

variable "acl" {
  type        = string
  description = "ACL to apply. String type."
  default     = "private"
}

variable "force_destroy" {
  type        = bool
  description = "All objects should be deleted from the bucket so that the bucket can be destroyed without error. Boolean type."
  default     = true
}

variable "versioning" {
  type        = bool
  description = "Enable versioning. Boolean type."
  default     = true
}

variable "key" {
  type        = string
  description = "Name of the object once it is in the bucket. String type."
  default     = "stable/home-assessment"
}
