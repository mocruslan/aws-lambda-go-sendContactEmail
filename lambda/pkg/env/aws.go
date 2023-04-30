package env

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session = nil

// GetAWSSession returns the AWS session
func GetAWSSession() *session.Session {
	if sess != nil {
		return sess
	}

	var err error
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		panic(err)
	}
	return sess
}
