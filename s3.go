package main

import (
	// "fmt"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (S3Bucket S3BucketObject) BucketExit(ctx context.Context) (bool, error) {
	exists := false
	S3BucketList, err := S3Bucket.S3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Printf("Couldn't determine existence of S3 Bucket: %v. Here's why: %v\n", *S3Bucket.BucketName, err)
	}
	for _, bucket := range S3BucketList.Buckets {
		// log.Printf("Checking if Bucket [%v] is the same as [%v] \n", *bucket.Name, *S3Bucket.BucketName)
		if *bucket.Name == *S3Bucket.BucketName {
			exists = true
		}
	}
	return exists, err
}

func ChangeStorageClass(ctx context.Context, S3Bucket S3BucketObject, S3ListOfObjects *s3.ListObjectsV2Output, s3SourceClass string, s3DestinationClass string) error {
	for _, Object := range S3ListOfObjects.Contents {
		//	log.Printf("Object [%v] of Storage Class [%v] and we are looking for [%v]\n", *Object.Key, Object.StorageClass,types.ObjectStorageClass(s3SourceClass))
		if Object.StorageClass == types.ObjectStorageClass(s3SourceClass) {
			log.Printf("Found Object [%v] of as Storage Class [%v]\n", *Object.Key, Object.StorageClass)
			CopySource := *S3Bucket.BucketName + "/"  + *Object.Key
			PutObjectUpdate := s3.CopyObjectInput {
				Bucket:      S3ListOfObjects.Name,
				Key:         Object.Key,
				CopySource:  &CopySource,
				StorageClass: types.StorageClass(s3DestinationClass),
			 }
			_, err := S3Bucket.S3Client.CopyObject(ctx, &PutObjectUpdate)
			if err != nil {
				log.Printf("Couldn't update Storage Class of object [%v]. Here's why: %v\n", *Object.Key, err)
			}
		}

	}

	// if err != nil {
	// 	log.Printf("Couldn't change storage class of S3 Bucket: %v. Here's why: %v\n", *S3Bucket.BucketName, err)
	// }
	// return err

	return nil
}

func isStorageClassCorrect(s3StorageClass types.ObjectStorageClass) bool {

	for _, Class := range s3SourceStorageClass.Values() {
		if Class == s3StorageClass {
			log.Printf("Found Storage Class [%v] in the list of the Available Storage Class [%v]\n", s3StorageClass, Class)
			return true
		}
	}
	return false
}
