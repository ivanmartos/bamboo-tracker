package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
)

type SesRepositoryImpl struct {
	api sesiface.SESAPI
}

func InitSesRepositoryImpl(api sesiface.SESAPI) timesheetUploader.SesRepository {
	return SesRepositoryImpl{api: api}
}

const charset = "UTF-8"

func (s SesRepositoryImpl) SendEmail(sender string, recipient string, htmlBody string, subject string) {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(htmlBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err := s.api.SendEmail(input)

	if err != nil {
		panic(err)
	}
}
