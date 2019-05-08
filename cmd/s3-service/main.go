package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func exitErrorf(msg string, args ...interface{}) {

}

func main() {
	awsProfile := os.Getenv("AWS_TEST_PROFILE")
	testBucket := os.Getenv("AWS_TEST_BUCKET")
	testRegion := os.Getenv("AWS_TEST_REGION")
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(testRegion),
			Credentials: credentials.NewSharedCredentials("", awsProfile),
		},
	})

	if err != nil {
		return
	}

	svc := s3.New(sess)

	ListButkets(svc)

	CreateBucket(svc, testBucket)

	ListButkets(svc)

	ListObjects(svc, testBucket)
}

func CreateBucket(svc *s3.S3, bucket string) {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	result, err := svc.CreateBucket(input)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, awsErr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, awsErr.Error())
			default:
				fmt.Println(awsErr.Error())
			}
		} else {
			fmt.Println(awsErr.Error())
		}
		return
	}
	fmt.Println(result)
}

func ListButkets(svc *s3.S3) {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func ListObjects(svc *s3.S3, bucket string) {
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(bucket),
		MaxKeys: aws.Int64(2),
	}
	result, err := svc.ListObjects(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, awsErr.Error())
			default:
				fmt.Println(awsErr.Error())
			}
		} else {
			fmt.Println(awsErr.Error())
		}
		return
	}

	fmt.Println(result)
}

func UploadObject() {

}
