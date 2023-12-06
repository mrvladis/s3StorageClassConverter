package main

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"sync"
)


type S3BucketObject struct {
	S3Client   *s3.Client
	BucketName *string
}

type WaitGroupCount struct {
    sync.WaitGroup
    count int64
}