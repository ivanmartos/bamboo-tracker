package timesheetUploader

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"time"
)

type TimesheetParserImpl struct{}

func InitTimesheetParser() TimesheetParser {
	return TimesheetParserImpl{}
}

const (
	timesheetEntryDateLayout = "2006-01-02"
	timesheetEntryTimeLayout = "15:04"
)

type configEntry struct {
	Start string
	End   string
}

func getValidatedConfigTime(configTime string) string {
	date, err := time.Parse(timesheetEntryTimeLayout, configTime)
	if err != nil {
		log.Println("Unable to parse config time entry", configTime)
		panic(err)
	}
	return date.Format(timesheetEntryTimeLayout)
}

func (TimesheetParserImpl) GetTimesheetEntries(weekday time.Weekday) []TimesheetEntry {
	var config = make(map[string][]configEntry)

	timesheets := os.Getenv("TIMESHEETS")

	log.Println(timesheets)
	err := yaml.Unmarshal([]byte(timesheets), &config)
	if err != nil {
		panic(err)
	}

	currentConfigEntries := config[strings.ToLower(weekday.String())]

	var entries []TimesheetEntry

	dateStr := time.Now().Format(timesheetEntryDateLayout)

	for _, entryConfig := range currentConfigEntries {
		entry := &TimesheetEntry{
			Date:  dateStr,
			Start: getValidatedConfigTime(entryConfig.Start),
			End:   getValidatedConfigTime(entryConfig.End),
		}

		entries = append(entries, *entry)
	}

	log.Println("Parsed timesheetUploader entries for current day", entries)

	return entries
}
