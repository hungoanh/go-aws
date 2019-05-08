package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"os"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(os.Getenv("AWS_TEST_REGION")),
			Credentials: credentials.NewSharedCredentials("", os.Getenv("AWS_TEST_PROFILE")),
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

	var paraminput = make([]*string, 0)
	n := "north"
	paraminput = append(paraminput, &n)
	s := "south"
	paraminput = append(paraminput, &s)

	params, err := svc.GetParameters(&ssm.GetParametersInput{
		Names: paraminput,
	})

	for _, p := range params.Parameters {
		fmt.Printf("parameter value is: %s\n", *p.Value)
	}
}

func exitErrorf(msg string, args ...interface{}) {

}
