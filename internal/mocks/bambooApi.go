package mocks

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
)

type MockBambooApi struct {
	BambooSession                  model.BambooSession
	LogInUsernameParam             *string
	LogInPasswordParam             *string
	LogInCalled                    bool
	AddTimesheetRecordCalled       bool
	AddTimesheetRecordSessionParam *model.BambooSession
	AddTimesheetRecordEntriesParam []model.TimesheetEntry
}

func (m *MockBambooApi) LogIn(username string, password string) model.BambooSession {
	m.LogInCalled = true
	m.LogInUsernameParam = &username
	m.LogInPasswordParam = &password
	return m.BambooSession
}

func (m *MockBambooApi) AddTimesheetRecord(session model.BambooSession, entries []model.TimesheetEntry) {
	m.AddTimesheetRecordCalled = true
	m.AddTimesheetRecordSessionParam = &session
	m.AddTimesheetRecordEntriesParam = entries
}

func (m *MockBambooApi) GetHomeContent() model.TimeTracking {
	//TODO
	return model.TimeTracking{}
}

func (m *MockBambooApi) Reset() {
	m.AddTimesheetRecordCalled = false
	m.AddTimesheetRecordEntriesParam = nil
	m.AddTimesheetRecordSessionParam = nil
	m.LogInCalled = false
	m.LogInPasswordParam = nil
	m.LogInUsernameParam = nil
}
