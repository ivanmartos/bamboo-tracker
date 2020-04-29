package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() *s3.S3 {
	awsSession, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	return s3.New(awsSession)
}
