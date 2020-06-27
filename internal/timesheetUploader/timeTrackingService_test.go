package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"os"
	"reflect"
	"testing"
)

func TestTimeTrackingServiceImpl_GetCurrentTimeTracking(t *testing.T) {
	username := "username"
	password := "password"

	_ = os.Setenv("BAMBOO_USERNAME", username)
	_ = os.Setenv("BAMBOO_PASSWORD", password)

	anyTimeTracking := model.TimeTracking{
		DailyTotal:  10,
		WeeklyTotal: 20,
	}

	mockBambooApi := &mocks.MockBambooApi{}
	mockBambooApi.GetHomeContentFunc = func() model.TimeTracking {
		return anyTimeTracking
	}

	tests := []struct {
		name string
		want model.TimeTracking
	}{
		{
			name: "Return output of API",
			want: anyTimeTracking,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := InitTimeTrackingService(mockBambooApi)
			if got := s.GetCurrentTimeTracking(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentTimeTracking() = %v, want %v", got, tt.want)
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
		})
	}
}

func TestTimeTrackingServiceImpl_UploadTimesheetEntries(t *testing.T) {
	username := "username"
	password := "password"

	_ = os.Setenv("BAMBOO_USERNAME", username)
	_ = os.Setenv("BAMBOO_PASSWORD", password)

	var anyBambooSession = model.BambooSession{
		EmployeeId: "123",
		Id:         "321",
	}

	var mockBambooApi = &mocks.MockBambooApi{
		BambooSession: anyBambooSession,
	}

	var anyTimesheetEntries = []model.TimesheetEntry{{Date: "2020-02-02", Start: "11:00", End: "12:00"}}

	s := InitTimeTrackingService(mockBambooApi)

	s.UploadTimesheetEntries(anyTimesheetEntries)

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
