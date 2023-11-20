package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


type S3BucketObject struct {
	S3Client   *s3.Client
	BucketName *string
}
