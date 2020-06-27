package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
)

type TimeTrackingServiceImpl struct {
	api BambooApi
}

func InitTimeTrackingService(api BambooApi) TimeTrackingService {
	return TimeTrackingServiceImpl{
		api: api,
	}
}

func (s TimeTrackingServiceImpl) GetCurrentTimeTracking() model.TimeTracking {
	s.api.LogIn(getEnvVariable("BAMBOO_USERNAME"), getEnvVariable("BAMBOO_PASSWORD"))
	timeTracking := s.api.GetHomeContent()
	return timeTracking
}

func (s TimeTrackingServiceImpl) UploadTimesheetEntries(timesheetEntries []model.TimesheetEntry) {
	session := s.api.LogIn(getEnvVariable("BAMBOO_USERNAME"), getEnvVariable("BAMBOO_PASSWORD"))
	s.api.AddTimesheetRecord(session, timesheetEntries)
}
