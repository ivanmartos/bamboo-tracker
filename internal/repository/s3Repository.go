package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
	"io/ioutil"
)

type S3RepositoryImpl struct {
	api s3iface.S3API
}

func InitS3RepositoryImpl(api s3iface.S3API) timesheetUploader.S3Repository {
	return S3RepositoryImpl{api: api}
}

func (repo S3RepositoryImpl) GetS3FileContent(key string, bucket string) string {
	resp, err := repo.api.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		panic(err)
	}

	s3objectBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(s3objectBytes)
}
