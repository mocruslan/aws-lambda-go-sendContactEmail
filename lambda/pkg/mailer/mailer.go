package mailer

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"lambda/pkg/env"
	"lambda/pkg/logging"
)

type Mailer interface {
	Send(from string, to string, subject string, body string) error
}

type mailer struct {
	client *sesv2.SESV2
}

// NewMailer creates a new mailer
func NewMailer() Mailer {
	return &mailer{}
}

// Send sends an email
func (m *mailer) Send(from string, to string, subject string, body string) error {
	logging.Debugf("Sending email from [%s] to [%s]", from, to)
	sendEmailInput := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(from),
		Destination: &sesv2.Destination{
			ToAddresses: aws.StringSlice([]string{to}),
		},
		Content: &sesv2.EmailContent{
			Simple: &sesv2.Message{
				Subject: &sesv2.Content{
					Data: aws.String(subject),
				},
				Body: &sesv2.Body{
					Text: &sesv2.Content{
						Data: aws.String(body),
					},
				},
			},
		},
	}

	_, err := m.getClient().SendEmail(sendEmailInput)
	if err != nil {
		logging.Errorf("Error sending email: %s", err.Error())
		return err
	}

	logging.Infof("Successfully sent email to [%s]", to)
	return nil

}

// getClient returns the AWS SES client
func (m *mailer) getClient() *sesv2.SESV2 {
	if m.client != nil {
		logging.Debug("Using cached SESV2 client")
		return m.client
	}

	logging.Debug("Creating new SESV2 client")
	m.client = sesv2.New(env.GetAWSSession())
	return m.client
}
