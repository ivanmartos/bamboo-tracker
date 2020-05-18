package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
)

func GetS3Client() *s3.S3 {
	awsSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	return s3.New(awsSession)
}

func GetSesClient() *ses.SES {
	awsSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	return ses.New(awsSession)
}
