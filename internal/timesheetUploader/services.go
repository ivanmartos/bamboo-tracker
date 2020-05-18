package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"time"
)

type BambooApi interface {
	LogIn(username string, password string) model.BambooSession
	AddTimesheetRecord(session model.BambooSession, entries []model.TimesheetEntry)
	GetHomeContent() model.TimeTracking
}

type TimesheetParser interface {
	GetTimesheetEntries(weekday time.Weekday) []model.TimesheetEntry
}

type S3Repository interface {
	GetS3FileContent(key string, bucket string) string
}
