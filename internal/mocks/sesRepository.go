package mocks

type MockSesRepository struct {
	SendEmailSenderArg    string
	SendEmailRecipientArg string
	SendEmailHtmlBodyArg  string
	SendEmailSubjectArg   string
}

func (m *MockSesRepository) SendEmail(sender string, recipient string, htmlBody string, subject string) {
	m.SendEmailSenderArg = sender
	m.SendEmailRecipientArg = recipient
	m.SendEmailHtmlBodyArg = htmlBody
	m.SendEmailSubjectArg = subject
}
