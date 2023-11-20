package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var s3SourceStorageClass types.ObjectStorageClass

func main() {
	fmt.Println("Welcome to S3 Storage Class Converter")
	s3BucketName := flag.String("s3BucketName", "", "S3 Bucket Name")
	s3SourceClass := flag.String("s3SourceClass", "STANDARD", "Current S3 Storage Class; Accepted values are [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
	s3DestinationClass := flag.String("s3DestinationClass", "GLACIER_IR", "To be S3 Storage Class; Accepted values are [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
	s3Region := flag.String("s3Region", "eu-west-2", "AWS Region")
	flag.Parse()
	fmt.Println("s3BucketName:", *s3BucketName)
	// fmt.Println("s3SourceClass:", *s3SourceClass)
	//fmt.Println("s3DestinationClass:", *s3DestinationClass)
	fmt.Println("s3Region:", *s3Region)

	if isStorageClassCorrect(types.ObjectStorageClass(*s3SourceClass)) {
		fmt.Printf("s3SourceClass [%v] is a correct S3 storage class \n", *s3SourceClass)
	} else {
		fmt.Printf("s3SourceClass [%v] is NOT a correct S3 storage class \n", *s3SourceClass)
		fmt.Println("Please use one of the following storage classes: [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
		//panic("Please use one of the following storage classes: [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
		return
	}

	if isStorageClassCorrect(types.ObjectStorageClass(*s3DestinationClass)) {
		fmt.Printf("s3SourceClass [%v] is a correct S3 storage class \n", *s3DestinationClass)
	} else {
		fmt.Printf("s3SourceClass [%v] is NOT a correct S3 storage class \n", *s3DestinationClass)
		fmt.Println("Please use one of the following storage classes: [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
		//panic("Please use one of the following storage classes: [STANDARD REDUCED_REDUNDANCY GLACIER STANDARD_IA ONEZONE_IA INTELLIGENT_TIERING DEEP_ARCHIVE OUTPOSTS GLACIER_IR SNOW]")
		return
	}

	// connecting to AWS
	ctx := context.TODO()
	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("eu-west-2"),
	)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	S3Client := s3.NewFromConfig(awsCfg)

	S3Bucket := S3BucketObject{
		S3Client:   S3Client,
		BucketName: s3BucketName,
	}
	BucketExists := false

	BucketExists, err = S3Bucket.BucketExit(ctx)
	if err != nil {
		log.Printf("Couldn't determine existence of S3 Bucket: %v. Here's why: %v\n", S3Bucket.BucketName, err)
	}
	if BucketExists {
		fmt.Printf("Bucket [%v] exists \n", *S3Bucket.BucketName)
	} else {
		fmt.Printf("Bucket [%v] not found \n", *S3Bucket.BucketName)
	}

	ListObjectsInput := s3.ListObjectsV2Input{
		Bucket: S3Bucket.BucketName,
	}

	S3ListOfObjects, s3err := S3Bucket.S3Client.ListObjectsV2(ctx, &ListObjectsInput)

	if s3err != nil {
		log.Printf("Couldn't get list of object(s) from S3: %v. Here's why: %v\n", S3Bucket.BucketName, err)
	}

	err = ChangeStorageClass(ctx,S3Bucket, S3ListOfObjects, *s3SourceClass, *s3DestinationClass)

	if err != nil {
		log.Printf("Couldn't convert object(s) in S3: %v. Here's why: %v\n", S3Bucket.BucketName, err)
	} else {
		log.Printf("Conversion of object(s) in S3 has been finished")
	}

	// fmt.Println(ListObjectsInput)
	// fmt.Println(s3SourceStorageClass)
	fmt.Println(s3SourceStorageClass.Values())

}
