package repository

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type BambooApiImpl struct {
	Client    HTTPClient
	csrfToken string
}

func InitBambooApi() timesheetUploader.BambooApi {
	var cookieJar, _ = cookiejar.New(nil)

	return &BambooApiImpl{
		Client: &http.Client{
			Jar: cookieJar,
		},
		csrfToken: "",
	}
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
const acceptEncoding = "gzip, deflate, br"

func parseCsrfToken(body string) string {
	r, _ := regexp.Compile(`var CSRF_TOKEN = "([^"]+)";`)
	return r.FindStringSubmatch(body)[1]
}

func parseSession(body string) model.BambooSession {
	r, _ := regexp.Compile(`var SESSION_USER=([^;]+);`)

	var session model.BambooSession
	sessionUserJsonStr := r.FindStringSubmatch(body)[1]

	_ = json.Unmarshal([]byte(sessionUserJsonStr), &session)

	return session
}

func getResponseBody(resp http.Response) string {
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	bodyBytes, bytesErr := ioutil.ReadAll(reader)
	if bytesErr != nil {
		panic(bytesErr)
	}

	return string(bodyBytes)
}

func setHeaders(r http.Request) {
	r.Header.Add("accept-encoding", acceptEncoding)
	r.Header.Add("upgrade-insecure-requests", "1")
	r.Header.Add("user-agent", userAgent)
}

func (api *BambooApiImpl) updateCsrfToken(body string) {
	api.csrfToken = parseCsrfToken(body)
}

func (api *BambooApiImpl) setInitialCsrfToken() {
	u, _ := url.Parse(os.Getenv("BAMBOO_HOST") + "/login.php")

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	setHeaders(*req)

	resp, err := api.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Initial response status", resp.Status)

	if resp.StatusCode != http.StatusOK {
		panic("Initial did not return 200. Returned - " + resp.Status)
	}

	body := getResponseBody(*resp)

	api.updateCsrfToken(body)
}

func (api *BambooApiImpl) LogIn(username string, password string) model.BambooSession {
	api.setInitialCsrfToken()

	u, _ := url.Parse(os.Getenv("BAMBOO_HOST") + "/login.php")
	data := url.Values{}
	data.Set("tz", "Europe/Berlin")
	data.Set("login", "Log in")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("CSRFToken", api.csrfToken)

	req, _ := http.NewRequest(http.MethodPost, u.String(), strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	setHeaders(*req)

	resp, err := api.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Log in response status", resp.Status)
	if resp.StatusCode != http.StatusOK {
		panic("LogIn api method did not return 200. Returned - " + resp.Status)
	}

	body := getResponseBody(*resp)

	api.updateCsrfToken(body)

	return parseSession(body)
}

func (api *BambooApiImpl) AddTimesheetRecord(session timesheetUploader.BambooSession, entries []timesheetUploader.TimesheetEntry) {
	dto := mapToDto(entries, session)
	dtoJson, _ := json.Marshal(dto)

	log.Println("Timesheet request json", string(dtoJson))

	u, _ := url.Parse(os.Getenv("BAMBOO_HOST") + "/timesheet/clock/entries")
	req, _ := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(dtoJson))

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("x-csrf-token", api.csrfToken)
	setHeaders(*req)

	resp, err := api.Client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("Create timesheetUploader response status", resp.Status)

	if resp.StatusCode != http.StatusOK {
		panic("Adding timesheet record did not return 200. Returned - " + resp.Status)
	}
}
