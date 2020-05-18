package mocks

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"time"
)

type MockTimesheetParser struct {
	TimesheetEntries []model.TimesheetEntry
}

func (m MockTimesheetParser) GetTimesheetEntries(_ time.Weekday) []model.TimesheetEntry {
	return m.TimesheetEntries
}
