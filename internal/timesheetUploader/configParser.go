package timesheetUploader

import (
	"github.com/ivanmartos/bamboo-tracker/internal/model"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
	"time"
)

type TimesheetParserImpl struct {
	s3Repository S3Repository
}

func InitTimesheetParser(s3Repository S3Repository) TimesheetParser {
	return TimesheetParserImpl{s3Repository: s3Repository}
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

func (p TimesheetParserImpl) getTimesheetContent() string {
	log.Println("Reading timesheets from s3")
	return p.s3Repository.GetS3FileContent(getEnvVariable("TIMESHEET_S3_KEY"), getEnvVariable("TIMESHEET_S3_BUCKET"))
}

func (p TimesheetParserImpl) GetTimesheetEntries(weekday time.Weekday) []model.TimesheetEntry {
	var config = make(map[string][]configEntry)

	timesheets := p.getTimesheetContent()

	log.Println(timesheets)
	err := yaml.Unmarshal([]byte(timesheets), &config)
	if err != nil {
		panic(err)
	}

	currentConfigEntries := config[strings.ToLower(weekday.String())]

	var entries []model.TimesheetEntry

	dateStr := time.Now().Format(timesheetEntryDateLayout)

	for _, entryConfig := range currentConfigEntries {
		entry := &model.TimesheetEntry{
			Date:  dateStr,
			Start: getValidatedConfigTime(entryConfig.Start),
			End:   getValidatedConfigTime(entryConfig.End),
		}

		entries = append(entries, *entry)
	}

	log.Println("Parsed timesheetUploader entries for current day", entries)

	return entries
}
