package timesheetUploader

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func verifyTimesheetEntry(t *testing.T, parser TimesheetParser, weekday time.Weekday, expectedEntries []TimesheetEntry) {
	result := parser.GetTimesheetEntries(weekday)

	if len(expectedEntries) == 0 {
		if len(result) != 0 {
			t.Errorf("Expected len of result array for %v to be 0, got %v", weekday, len(result))
		}
	} else {
		if !reflect.DeepEqual(result, expectedEntries) {
			t.Errorf("Expected result of %v entries was %v, got %v", weekday, expectedEntries, result)

		}
	}
}

func TestGetTimesheetEntries(t *testing.T) {
	parser := InitTimesheetParser()
	yaml := `
monday:
 - start: 11:00
   end: 12:00
saturday:
 - start: 11:10
   end: 12:10
 - start: 13:10
   end: 14:10
Sunday:
 - start: 11:10
   end: 12:10
`
	_ = os.Setenv("TIMESHEETS", yaml)

	currentDate := time.Now().Format(timesheetEntryDateLayout)

	verifyTimesheetEntry(t, parser, time.Monday, []TimesheetEntry{{currentDate, "11:00", "12:00"}})
	verifyTimesheetEntry(t, parser, time.Saturday, []TimesheetEntry{{currentDate, "11:10", "12:10"}, {currentDate, "13:10", "14:10"}})
	// Sunday is not written in correct format
	verifyTimesheetEntry(t, parser, time.Sunday, []TimesheetEntry{})
	verifyTimesheetEntry(t, parser, time.Tuesday, []TimesheetEntry{})
}
