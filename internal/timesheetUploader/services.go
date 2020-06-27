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

type SesRepository interface {
	SendEmail(sender string, recipient string, htmlBody string, subject string)
}

type EmailComposer interface {
	ComposeTimeTrackingEmailPayload(timeTracking model.TimeTracking) model.TimeTrackingEmailPayload
}

type EmailSender interface {
	SendEmail(payload model.TimeTrackingEmailPayload)
}

type TimeTrackingService interface {
	GetCurrentTimeTracking() model.TimeTracking
	UploadTimesheetEntries(timesheetEntries []model.TimesheetEntry)
}
