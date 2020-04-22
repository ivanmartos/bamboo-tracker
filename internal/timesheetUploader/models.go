package timesheetUploader

type BambooSession struct {
	EmployeeId string
	Id         string
}

type TimesheetEntry struct {
	Date  string
	Start string
	End   string
}
