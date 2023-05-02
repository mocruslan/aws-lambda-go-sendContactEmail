package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"lambda/pkg/env"
	"lambda/pkg/logging"
	"lambda/pkg/mailer"
)

type emailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func handler(request emailRequest) (events.APIGatewayProxyResponse, error) {
	fromEmail := env.GetEnv(env.FromEmail, "")
	if fromEmail == "" {
		logging.Error("From email was not set")
		return events.APIGatewayProxyResponse{}, errors.New("from email not set")
	}

	mail := mailer.NewMailer()
	err := mail.Send(fromEmail, request.To, request.Subject, request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, errors.New("error sending email")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
