package repository

import (
	"github.com/ivanmartos/bamboo-tracker/internal/timesheetUploader"
	"strconv"
)

type timesheetEntryDto struct {
	Id         *int   `json:"id"`
	TrackingId int    `json:"trackingId"`
	EmployeeId int    `json:"employeeId"`
	Date       string `json:"date"`
	Start      string `json:"start"`
	End        string `json:"end"`
}

type timesheetDto struct {
	Entries []timesheetEntryDto `json:"entries"`
}

func mapToDto(entries []timesheetUploader.TimesheetEntry, session timesheetUploader.BambooSession) timesheetDto {
	timesheetDto := timesheetDto{}

	employeeId, _ := strconv.Atoi(session.EmployeeId)

	for _, entry := range entries {

		dto := &timesheetEntryDto{
			Id:         nil,
			TrackingId: 1,
			EmployeeId: employeeId,
			Date:       entry.Date,
			Start:      entry.Start,
			End:        entry.End,
		}
		timesheetDto.Entries = append(timesheetDto.Entries, *dto)
	}

	return timesheetDto
}
