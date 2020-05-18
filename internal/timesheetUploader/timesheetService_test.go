package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"os"
	"reflect"
	"testing"
)

var anyTimesheetEntries = []model.TimesheetEntry{{Date: "2020-02-02", Start: "11:00", End: "12:00"}}

var anyBambooSession = model.BambooSession{
	EmployeeId: "123",
	Id:         "321",
}

func TestTimesheetService_UploadCurrentTimesheet(t *testing.T) {
	var timesheetService TimesheetService

	var mockBambooApi = &mocks.MockBambooApi{
		BambooSession: anyBambooSession,
	}

	var mockEmptyTimesheetParser = mocks.MockTimesheetParser{}

	// for empty timesheetUploader should not call api
	timesheetService = InitTimesheetService(
		mockBambooApi,
		mockEmptyTimesheetParser,
	)

	timesheetService.UploadCurrentTimesheet()

	if mockBambooApi.LogInCalled || mockBambooApi.AddTimesheetRecordCalled {
		t.Errorf("LogIn or AddTimesheet of bambooApi was called even though there are no records to be added")
	}

	username := "username"
	password := "password"

	_ = os.Setenv("BAMBOO_USERNAME", username)
	_ = os.Setenv("BAMBOO_PASSWORD", password)
	mockBambooApi.Reset()

	// for entries in timesheetUploader should log in and call api
	timesheetService = InitTimesheetService(
		mockBambooApi,
		mocks.MockTimesheetParser{TimesheetEntries: anyTimesheetEntries},
	)

	timesheetService.UploadCurrentTimesheet()

	if !mockBambooApi.LogInCalled || !mockBambooApi.AddTimesheetRecordCalled {
		t.Errorf("LogIn or AddTimesheet of bambooApi was not called when there are records to be added")
	}

	if !reflect.DeepEqual(anyTimesheetEntries, mockBambooApi.AddTimesheetRecordEntriesParam) {
		t.Errorf("Bamboo API received wrong timesheetUploader params for upload")
	}

	if anyBambooSession != *mockBambooApi.AddTimesheetRecordSessionParam {
		t.Errorf("BambooAPI received wrong session param for upload")
	}

	if *mockBambooApi.LogInUsernameParam != username {
		t.Errorf("BambooAPI login received wrong username. Expected %v received %v", username, *mockBambooApi.LogInUsernameParam)
	}

	if *mockBambooApi.LogInPasswordParam != password {
		t.Errorf("BambooAPI login received wrong password. Expected %v received %v", password, *mockBambooApi.LogInPasswordParam)
	}
}
