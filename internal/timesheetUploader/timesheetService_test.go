package timesheetUploader

import (
	"os"
	"reflect"
	"testing"
	"time"
)

type MockTimesheetParser struct{}

var anyTimesheetEntries = []TimesheetEntry{{Date: "2020-02-02", Start: "11:00", End: "12:00"}}

func (m MockTimesheetParser) GetTimesheetEntries(_ time.Weekday) []TimesheetEntry {
	return anyTimesheetEntries
}

type MockEmptyTimesheetParser struct{}

func (m MockEmptyTimesheetParser) GetTimesheetEntries(_ time.Weekday) []TimesheetEntry {
	return []TimesheetEntry{}
}

type MockBambooApi struct{}

var logInUsernameParam *string
var logInPasswordParam *string
var logInCalled = false

var anyBambooSession = BambooSession{
	EmployeeId: "123",
	Id:         "321",
}

func (m MockBambooApi) LogIn(username string, password string) BambooSession {
	logInCalled = true
	logInUsernameParam = &username
	logInPasswordParam = &password
	return anyBambooSession
}

var addTimesheetRecordCalled = false
var addTimesheetRecordSessionParam *BambooSession
var addTimesheetRecordEntriesParam []TimesheetEntry

func (m MockBambooApi) AddTimesheetRecord(session BambooSession, entries []TimesheetEntry) {
	addTimesheetRecordCalled = true
	addTimesheetRecordSessionParam = &session
	addTimesheetRecordEntriesParam = entries
}

func reset() {
	addTimesheetRecordCalled = false
	addTimesheetRecordEntriesParam = nil
	addTimesheetRecordSessionParam = nil
	logInCalled = false
	logInPasswordParam = nil
	logInUsernameParam = nil
}

func TestTimesheetService_UploadCurrentTimesheet(t *testing.T) {
	var timesheetService TimesheetService

	// for empty timesheetUploader should not call api
	timesheetService = InitTimesheetService(
		MockBambooApi{},
		MockEmptyTimesheetParser{},
	)

	timesheetService.UploadCurrentTimesheet()

	if logInCalled || addTimesheetRecordCalled {
		t.Errorf("LogIn or AddTimesheet of bambooApi was called even though there are no records to be added")
	}

	username := "username"
	password := "password"

	_ = os.Setenv("BAMBOO_USERNAME", username)
	_ = os.Setenv("BAMBOO_PASSWORD", password)
	reset()

	// for entries in timesheetUploader should log in and call api
	timesheetService = InitTimesheetService(
		MockBambooApi{},
		MockTimesheetParser{},
	)

	timesheetService.UploadCurrentTimesheet()

	if !logInCalled || !addTimesheetRecordCalled {
		t.Errorf("LogIn or AddTimesheet of bambooApi was not called when there are records to be added")
	}

	if !reflect.DeepEqual(anyTimesheetEntries, addTimesheetRecordEntriesParam) {
		t.Errorf("Bamboo API received wrong timesheetUploader params for upload")
	}

	if anyBambooSession != *addTimesheetRecordSessionParam {
		t.Errorf("BambooAPI received wrong session param for upload")
	}

	if *logInUsernameParam != username {
		t.Errorf("BambooAPI login received wrong username. Expected %v received %v", username, *logInUsernameParam)
	}

	if *logInPasswordParam != password {
		t.Errorf("BambooAPI login received wrong password. Expected %v received %v", password, *logInPasswordParam)
	}
}
