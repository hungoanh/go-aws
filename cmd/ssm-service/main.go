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