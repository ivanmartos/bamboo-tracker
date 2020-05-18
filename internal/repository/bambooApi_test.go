package repository

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/ivanmartos/bamboo-tracker/internal/mocks"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

const expectedCsfrToken = "2a1f616088d916693f6096382e9deb003e72abf62d5f560c430fdca1e106402d32db57f96fa61e4fa8f6c6c77e4979202bed983eb8c1ba0f91094989c0714b7f"

const csfrSnippet = `
<script type="text/javascript">
	var DEFAULT_ERROR_MESSAGE = "Uh oh, something went wrong. Please try again or contact support@bamboohr.com";
	var DEFAULT_PERMISSIONS_ERROR_MESSAGE = "You do not have sufficient permissions for this";
	var GLOBAL_DATEPICKER_MASK = "dd\x2Fmm\x2Fyy";
	var GLOBAL_NUMBER_REGEX = "\x5E\x28\x3F\x21\x2D0\x24\x29\x2D\x3F\x28\x3F\x3D\x5B0\x2D9\x5C.\x5D\x29\x28\x3F\x210\x5B0\x2D9\x5D\x29\x28\x3F\x3A\x5Cd\x2B\x7C\x5Cd\x7B1,3\x7D\x28\x3F\x3A\x5C\x27\x5Cd\x7B3\x7D\x29\x2B\x29\x3F\x28\x3F\x3A\x5C.\x5Cd\x2B\x29\x3F\x24";
	var CSRF_TOKEN = "` + expectedCsfrToken + `";
	var CAN_EMAIL_FILES_AND_REPORTS = 1;
	var EVENT_TRACKING_DEFINITIONS = ["login","addedAdminUser","addedUser","addedEmployee","addedLogo","addedTimeOffPolicy","addedTimeTrackingEmployeeSettings","addedTimeTrackingEmployeeProfile","addedTimeTrackingEmployeeApi","addedTimeTrackingEmployeeImporter","hiredApplicant","eventApprovedTimesheetBulk","eventApprovedTimesheetIndividual","assignedPolicy","assignedPolicyBulk","completedGoal","changedAccountOwner","completedManagerAssessment","scheduleManagerAssessmentFuture","completedSelfAssessment","createdBenefitGroup","createdBenefitPlan","createdAlert","addedHoliday","createdOnboardingTask","createdOffboardingTask","createdSignatureTemplate","createdTraining","importedOnboarding","importedOffboarding","savedPaySchedule","exportedPayrollHours","powerEdited","requestedFeedback","requestedSignature","requestedTimeOffMobile","requestedTimeOffWeb","sentEmailAlert","sentOfferLetter","sentNewHirePacket","submittedFeedback ","viewedFutureBalanceWeb","viewedFutureBalanceMobile","viewedReport","addedJobOpening","updatedBenefitStatus"];
</script>
`

const bambooHostEnvVar = "www.bamboo.com"

func TestBambooApiImpl_AddTimesheetRecord(t *testing.T) {
	_ = os.Setenv("BAMBOO_HOST", bambooHostEnvVar)

	session := model.BambooSession{
		EmployeeId: "123",
		Id:         "321",
	}

	timesheetEntry := model.TimesheetEntry{
		Date:  "2020-04-15",
		Start: "09:00",
		End:   "10:00",
	}
	timesheets := []model.TimesheetEntry{timesheetEntry}

	var request *http.Request
	mockClient := &mocks.MockClient{}
	mockClient.DoFunc = func(req *http.Request) (response *http.Response, err error) {
		request = req
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil

	}
	bambooApi := BambooApiImpl{Client: mockClient}

	bambooApi.AddTimesheetRecord(session, timesheets)

	if request.Method != http.MethodPost {
		t.Errorf("Request has wrong method. Received %v", request.Method)
	}

	if request.URL.String() != bambooHostEnvVar+"/timesheet/clock/entries" {
		t.Errorf("Request has wrong URL. Received %v", request.URL.String())
	}

	var requestDto timesheetDto
	_ = json.NewDecoder(request.Body).Decode(&requestDto)
	fmt.Println(requestDto)

	if len(requestDto.Entries) != 1 {
		t.Errorf("There are too many entries in request dto. Expected 1, got %v", len(requestDto.Entries))
	}

	expectedEntryDto := timesheetEntryDto{
		Id:         nil,
		TrackingId: 1,
		EmployeeId: 123,
		Date:       timesheetEntry.Date,
		Start:      timesheetEntry.Start,
		End:        timesheetEntry.End,
	}

	if requestDto.Entries[0] != expectedEntryDto {
		t.Errorf("Request DTO has different content then expected.")
	}

}

func TestBambooApiImpl_LogIn(t *testing.T) {
	_ = os.Setenv("BAMBOO_HOST", bambooHostEnvVar)

	sessionId := "123"
	employeeSessionId := "987"
	expectedCsfrToken2 := "123456789asdfghjkl"
	logInResponseHtmlSnippet := `
<script type="text/javascript">
    var SESSION_USER={"id":"` + sessionId + `","username":"user.name@mail.com","employeeId":"` + employeeSessionId + `","companyId":"1111"};
	var CSRF_TOKEN = "` + expectedCsfrToken2 + `";
</script>
`

	var sessionInitRequest http.Request
	var logInRequest http.Request

	mockClient := &mocks.MockClient{}
	mockClient.DoFunc = func(req *http.Request) (response *http.Response, err error) {
		switch req.Method {
		case http.MethodGet:
			sessionInitRequest = *req
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(csfrSnippet)),
			}, nil
		case http.MethodPost:
			logInRequest = *req
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(logInResponseHtmlSnippet)),
			}, nil
		default:
			t.Errorf("Unexpected http method called")
			return nil, nil
		}
	}

	bambooApi := BambooApiImpl{Client: mockClient}

	const (
		username = "anyUsername"
		password = "anyPassword"
	)

	session := bambooApi.LogIn(username, password)

	if session.EmployeeId != employeeSessionId {
		t.Errorf("Received wrong employee session id as expected. Received %v", session.EmployeeId)
	}

	if session.Id != sessionId {
		t.Errorf("Received wrong session id as expected. Received %v", session.Id)
	}

	if sessionInitRequest.URL.String() != bambooHostEnvVar+"/login.php" {
		t.Errorf("Init session has wrong URL. Received %v", sessionInitRequest.URL.String())
	}

	if sessionInitRequest.Method != http.MethodGet {
		t.Errorf("Init session has wrong method. Received %v", sessionInitRequest.Method)
	}

	if logInRequest.URL.String() != bambooHostEnvVar+"/login.php" {
		t.Errorf("Log in has wrong URL. Received %v", sessionInitRequest.URL.String())
	}

	if logInRequest.Method != http.MethodPost {
		t.Errorf("Log in has wrong host. Received %v", sessionInitRequest.Method)
	}

	if logInRequest.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Log in request does not have correct Content-Type header")
	}

	query, _ := ioutil.ReadAll(logInRequest.Body)
	logInRequestQuery, _ := url.ParseQuery(string(query))

	if logInRequestQuery.Get("username") != username {
		t.Errorf("LogIn request contains invalid username. Received %v", logInRequestQuery.Get("username"))
	}

	if logInRequestQuery.Get("password") != password {
		t.Errorf("LogIn request contains invalid password. Received %v", logInRequestQuery.Get("password"))
	}

	if logInRequestQuery.Get("tz") != "Europe/Berlin" {
		t.Errorf("LogIn request contains invalid tz. Received %v", logInRequestQuery.Get("tz"))
	}

	if logInRequestQuery.Get("login") != "Log in" {
		t.Errorf("LogIn request contains invalid login. Received %v", logInRequestQuery.Get("login"))
	}

	if logInRequestQuery.Get("CSRFToken") != expectedCsfrToken {
		t.Errorf("LogIn request contains invalid CSRFToken. Received %v", logInRequestQuery.Get("CSRFToken"))
	}

}

func Test_getResponseBody(t *testing.T) {
	expectedText := "foo bar"

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, _ = gz.Write([]byte(expectedText))
	_ = gz.Close()

	tests := []struct {
		name     string
		response http.Response
	}{
		{
			name: "Plain text response",
			response: http.Response{
				Body: ioutil.NopCloser(strings.NewReader(expectedText)),
			},
		},
		{
			name: "Encoded response",
			response: http.Response{
				Header: http.Header{"Content-Encoding": {"gzip"}},
				Body:   ioutil.NopCloser(bytes.NewReader(b.Bytes())),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedResponse := getResponseBody(tt.response)

			if expectedText != parsedResponse {
				t.Errorf("Did not receive expected test result")
			}
		})
	}

}

func Test_parseCsrfToken(t *testing.T) {
	csfrToken := parseCsrfToken(csfrSnippet)

	if csfrToken != expectedCsfrToken {
		t.Errorf("Expected token not recieved")
	}
}

func Test_setHeaders(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "www.foo.bar", nil)

	setHeaders(*req)

	tests := []struct {
		input string
		want  string
	}{
		{
			input: "accept-encoding",
			want:  acceptEncoding,
		},
		{
			input: "upgrade-insecure-requests",
			want:  "1",
		},
		{
			input: "user-agent",
			want:  userAgent,
		},
	}
	for _, tt := range tests {
		t.Run("For header "+tt.input, func(t *testing.T) {
			headerValue := req.Header.Get(tt.input)
			if headerValue != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, headerValue)
			}
		})
	}
}
