# s3listobjects

`s3listobjects` is simple http server which return s3 bucket objects.

## Prerequisites

- go 1.15+
- curl 7.64.1

## How to use

Make sure that you have configured you aws [credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html). AWS region should set via env variable AWS_REGION.

Run utility with:
```bash
$ go run main.go
```

Use `curl` or any other http client to make get request to server:

```bash
$ curl localhost:8080/?bucket=bucket_name
```

## How to build

Use following command to build binary executable:
```bash
$ go build -o s3listobjects
```