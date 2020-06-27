package model

type BambooSession struct {
	EmployeeId string
	Id         string
}

type TimesheetEntry struct {
	Date  string
	Start string
	End   string
}

type TimeTracking struct {
	DailyTotal  float64
	WeeklyTotal float64
}

type TimeTrackingEmailPayload struct {
	Sender    string
	Recipient string
	HtmlBody  string
	Subject   string
}
