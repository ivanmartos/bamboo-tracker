package repository

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"testing"
)

func TestSesRepositoryImpl_SendEmail(t *testing.T) {
	sesClientMock := &mocks.MockSesClient{}

	sesRepo := SesRepositoryImpl{api: sesClientMock}

	sender := "sender"
	recipient := "recipient"
	subject := "subject"
	htmlBody := "htmlContent"

	expectedChaset := "UTF-8"

	sesRepo.SendEmail(sender, recipient, htmlBody, subject)

	if len(sesClientMock.SendEmailInputRequestArgument.Destination.ToAddresses) != 1 {
		t.Errorf("Received unexpected size of toAddresses")
	}

	if *sesClientMock.SendEmailInputRequestArgument.Destination.ToAddresses[0] != recipient {
		t.Errorf("Received unexpected toAddresses")
	}

	if len(sesClientMock.SendEmailInputRequestArgument.Destination.CcAddresses) != 0 {
		t.Errorf("Received unexpected size of toAddresses")
	}

	if *sesClientMock.SendEmailInputRequestArgument.Message.Subject.Data != subject {
		t.Errorf("Received unexpected message subject - %v", sesClientMock.SendEmailInputRequestArgument.Message.Subject.String())
	}

	if *sesClientMock.SendEmailInputRequestArgument.Message.Subject.Charset != expectedChaset {
		t.Errorf("Received unexpected subject charset subject - %v", sesClientMock.SendEmailInputRequestArgument.Message.Subject.Charset)
	}

	if *sesClientMock.SendEmailInputRequestArgument.Message.Body.Html.Data != htmlBody {
		t.Errorf("Received unexpected message html body - %v", sesClientMock.SendEmailInputRequestArgument.Message.Body.Html.String())
	}

	if *sesClientMock.SendEmailInputRequestArgument.Message.Body.Html.Charset != expectedChaset {
		t.Errorf("Received unexpected html body charset subject - %v", sesClientMock.SendEmailInputRequestArgument.Message.Body.Html.Charset)
	}
}
