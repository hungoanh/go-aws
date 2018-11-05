package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"fmt"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {

}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
	    Profile: os.Getenv("AWS_PERSONAL_SECONDARY_PROFILE"),
	    Config: aws.Config{
	    	Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		return
	}

	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)

	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

