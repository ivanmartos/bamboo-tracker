package mocks

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type MockS3Client struct {
	s3iface.S3API

	GetObjectRequestArg s3.GetObjectInput

	GetObjectResponseOutput s3.GetObjectOutput
}

func (m *MockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	m.GetObjectRequestArg = *input

	return &m.GetObjectResponseOutput, nil
}
