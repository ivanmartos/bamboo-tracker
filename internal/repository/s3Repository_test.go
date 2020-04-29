package repository

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"io/ioutil"
	"strings"
	"testing"
)

func TestS3RepositoryImpl_GetS3FileContent(t *testing.T) {
	expectedBody := "This should be returned"

	s3ClientMock := &mocks.MockS3Client{
		GetObjectResponseOutput: s3.GetObjectOutput{
			Body: ioutil.NopCloser(strings.NewReader(expectedBody)),
		},
	}

	s3Repo := S3RepositoryImpl{api: s3ClientMock}
	key := "anyKey"
	bucket := "anyBucket"
	result := s3Repo.GetS3FileContent(key, bucket)

	if result != expectedBody {
		t.Errorf("Expected %v as GetS3FileContent result, received %v", expectedBody, result)
	}

	if *s3ClientMock.GetObjectRequestArg.Bucket != bucket {
		t.Errorf("Expected %v as requested bucket, received %v", bucket, *s3ClientMock.GetObjectRequestArg.Bucket)
	}

	if *s3ClientMock.GetObjectRequestArg.Key != key {
		t.Errorf("Expected %v as requested key, received %v", key, *s3ClientMock.GetObjectRequestArg.Key)
	}

}
