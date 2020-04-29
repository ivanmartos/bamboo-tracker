package timesheetUploader

import (
	"time"
)

type BambooApi interface {
	LogIn(username string, password string) BambooSession
	AddTimesheetRecord(session BambooSession, entries []TimesheetEntry)
}

type TimesheetParser interface {
	GetTimesheetEntries(weekday time.Weekday) []TimesheetEntry
}

type S3Repository interface {
	GetS3FileContent(key string, bucket string) string
}
