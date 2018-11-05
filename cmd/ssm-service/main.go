package main

import (
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"fmt"
)

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
	svc := ssm.New(sess)

	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name: aws.String("test_param"),
	})

	if param != nil {
		fmt.Printf("parameter value is: %s\n", *param.Parameter.Value)
	}
}

func exitErrorf(msg string, args ...interface{}) {

}