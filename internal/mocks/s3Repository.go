package mocks

type MockS3Repository struct {
	GetS3FileContentFunc      func(key string, bucket string) string
	GetS3FileContentKeyArg    string
	GetS3FileContentBucketArg string
}

func (m *MockS3Repository) GetS3FileContent(key string, bucket string) string {
	m.GetS3FileContentKeyArg = key
	m.GetS3FileContentBucketArg = bucket
	return m.GetS3FileContentFunc(key, bucket)
}
