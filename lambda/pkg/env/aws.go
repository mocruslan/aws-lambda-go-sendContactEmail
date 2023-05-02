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

	sess = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(GetEnv(AWSRegion, "eu-north-1")),
	}))
	return sess
}
