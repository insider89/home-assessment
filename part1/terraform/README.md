## AWS s3 helm chart repository

We use this terraform code to setup s3 bucket for helm chart repository.
For simplicity, we are using local state instead of remote one.

### Requirements

| Name  | Version |
| ------------- | ------------- |
| terraform  | >=1.0  |
| aws  | >=2.2.19  |


We assume, that you have configured aws cli and logged in to correct aws account(using aws cli).

### Provider
| Name  | Version |
| ----- | -----|
| aws  | ~> 3.0  |

### Usage

Run following command to init terraform:
```
terraform init
```

To view what terraform code will create, use following command:
```
terraform plan
```

To create bucket, use following command:
```
terrafarm apply
```

To destroy bucket, use following command:
```
terrafrom destroy
```

### Inputs
| Name  | Description |
| ------------- | ------------- |
| aws_region  | AWS region where s3 bucket will create. String type. |
| bucket_name  | S3 bucket name. String type. |
| acl  | ACL to apply. String type. |
| force_destroy  | All objects should be deleted from the bucket so that the bucket can be destroyed without error. Boolean type. |
| versioning  | Enable versioning. Boolean type. |
| key  | Name of the object once it is in the bucket. String type. |

### Outputs

| Name  | Description |
| ------------- | ------------- |
| s3_bucket_name  | s3 bucket name |
