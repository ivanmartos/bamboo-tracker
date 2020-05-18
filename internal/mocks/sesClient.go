package mocks

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

type MockSesClient struct {
	sesiface.SESAPI

	SendEmailInputRequestArgument ses.SendEmailInput

	SendEmailOutputResponse ses.SendEmailOutput
}

func (m *MockSesClient) SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
	m.SendEmailInputRequestArgument = *input

	return &m.SendEmailOutputResponse, nil
}
