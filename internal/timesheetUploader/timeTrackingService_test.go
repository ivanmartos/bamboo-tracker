package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"os"
	"testing"
)

func TestTimeTrackingService_SendCurrentTimeTrackEmail(t *testing.T) {
	username := "username"
	password := "password"
	senderEmail := "foo@bar.lol"
	recipientEmail := "bar@foo.lol"

	_ = os.Setenv("BAMBOO_USERNAME", username)
	_ = os.Setenv("BAMBOO_PASSWORD", password)
	_ = os.Setenv("DAILY_TIME_TRACKING_SENDER_EMAIL", senderEmail)
	_ = os.Setenv("DAILY_TIME_TRACKING_RECIPIENT_EMAIL", recipientEmail)

	anyTimeTracking := model.TimeTracking{
		DailyTotal:  10,
		WeeklyTotal: 20,
	}
	mockBambooApi := &mocks.MockBambooApi{}
	mockBambooApi.GetHomeContentFunc = func() model.TimeTracking {
		return anyTimeTracking
	}
	mockSesRepository := &mocks.MockSesRepository{}

	timeTrackingService := InitTimeTrackingService(mockBambooApi, mockSesRepository)

	timeTrackingService.SendCurrentTimeTrackEmail()

	if mockBambooApi.LogInCalled != true {
		t.Error("Expected login to be called")
	}

	if *mockBambooApi.LogInUsernameParam != username {
		t.Errorf("Received unexpected login username param. Got %v", *mockBambooApi.LogInUsernameParam)
	}

	if *mockBambooApi.LogInPasswordParam != password {
		t.Errorf("Received unexpected login password param. Got %v", *mockBambooApi.LogInPasswordParam)
	}

	if mockBambooApi.GetHomeContentCalled != true {
		t.Error("Expected GetHomeContent to be called")
	}

	if mockSesRepository.SendEmailSubjectArg == "" {
		t.Error("Received empty send email subject arg")
	}

	if mockSesRepository.SendEmailSenderArg != senderEmail {
		t.Errorf("Received unexpected email sender address - %v", mockSesRepository.SendEmailSenderArg)
	}

	if mockSesRepository.SendEmailRecipientArg != recipientEmail {
		t.Errorf("Received unexpected email recipuent address - %v", mockSesRepository.SendEmailRecipientArg)
	}

	if mockSesRepository.SendEmailHtmlBodyArg == "" {
		t.Error("Received empty send email html body")
	}
}
