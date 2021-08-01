package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func s3func(s3bucket string) (*s3.ListObjectsV2Output, error) {
	bucket := s3bucket

	// please provide region via AWS_REGION env variable
	// credentials from the shared credentials file ~/.aws/credentials, env variabled or iam roles
	sess, err := session.NewSession()

	// Create S3 service client
	svc := s3.New(sess)

	// Get the list of items
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to list objects")
	}

	return resp, nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	buckets, ok := r.URL.Query()["bucket"]

	if !ok || len(buckets[0]) < 1 {
		fmt.Fprintln(w, "Url Param 'bucket' is missing. Use syntaxes: localhost:8080/?bucket=$BUCKET_NAME")
		return
	}

	bucket := buckets[0]

	resp, err := s3func(bucket)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, resp)
	}
}
