package timesheetUploader

import (
	"log"
	"time"
)

type TimesheetService struct {
	api    BambooApi
	config TimesheetParser
}

func InitTimesheetService(api BambooApi, parser TimesheetParser) TimesheetService {
	return TimesheetService{
		api:    api,
		config: parser,
	}
}

func (s TimesheetService) UploadCurrentTimesheet() {
	timesheetEntries := s.config.GetTimesheetEntries(time.Now().Weekday())

	if len(timesheetEntries) != 0 {
		session := s.api.LogIn(getEnvVariable("BAMBOO_USERNAME"), getEnvVariable("BAMBOO_PASSWORD"))
		s.api.AddTimesheetRecord(session, timesheetEntries)
		log.Println("Successfully uploaded timesheetUploader entries for today")
	} else {
		log.Println("There are bamboo timesheetUploader entries for today")
	}
}
